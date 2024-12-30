package grpcs

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/web/grpcs/client"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// 定义错误常量
var (
	ErrPoolClosed = errors.New("连接池已关闭")
	ErrPoolEmpty  = errors.New("连接池已空")
)

// GrpcConnPool 连接池结构体
type GrpcConnPool struct {
	mu       sync.Mutex
	conns    []*grpc.ClientConn
	host     string
	size     int
	current  int
	isClosed bool
}

// GrpcPoolManager 连接池管理器
type GrpcPoolManager struct {
	mu    sync.RWMutex
	pools map[string]*GrpcConnPool
}

// PoolConfig 连接池配置结构体
type PoolConfig struct {
	size        int
	maxIdleTime time.Duration
}

// PoolOption 定义连接池选项函数类型
type PoolOption func(*PoolConfig)

// defaultPoolConfig 返回默认连接池配置
func defaultPoolConfig() *PoolConfig {
	return &PoolConfig{
		size:        10,               // 默认连接池大小
		maxIdleTime: 30 * time.Minute, // 默认最大空闲时间
	}
}

// WithPoolSize 设置连接池大小的选项
func WithPoolSize(size int) PoolOption {
	return func(config *PoolConfig) {
		if size > 0 {
			config.size = size
		}
	}
}

// WithMaxIdleTime 设置最大空闲时间的选项
func WithMaxIdleTime(duration time.Duration) PoolOption {
	return func(config *PoolConfig) {
		if duration > 0 {
			config.maxIdleTime = duration
		}
	}
}

// NewGrpcConnPool 创建新的连接池
func NewGrpcConnPool(host string, size int) (*GrpcConnPool, error) {
	if size <= 0 {
		return nil, errors.New("连接池大小必须大于0")
	}

	pool := &GrpcConnPool{
		host:     host,
		size:     size,
		conns:    make([]*grpc.ClientConn, 0, size),
		isClosed: false,
	}

	// 初始化连接
	for i := 0; i < size; i++ {
		conn, err := grpc_client.grpc_client.InitClient(host, config.CONF.System.Serct)
		if err != nil {
			pool.Close()
			return nil, fmt.Errorf("初始化连接失败: %v", err)
		}
		pool.conns = append(pool.conns, conn)
	}

	return pool, nil
}

// Get 获取连接
func (p *GrpcConnPool) Get() (*grpc.ClientConn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isClosed {
		return nil, ErrPoolClosed
	}

	if len(p.conns) == 0 {
		return nil, ErrPoolEmpty
	}

	// 检查连接是否有效
	conn := p.conns[p.current]
	if conn == nil {
		// 如果连接无效，尝试重新创建
		newConn, err := grpc_client.InitClient(p.host, config.CONF.System.Serct)
		if err != nil {
			return nil, fmt.Errorf("重新创建连接失败: %v", err)
		}
		p.conns[p.current] = newConn
		conn = newConn
	}

	// 轮询策略
	p.current = (p.current + 1) % len(p.conns)

	return conn, nil
}

// Put 将连接放回池中（如果需要实现自定义放回逻辑）
func (p *GrpcConnPool) Put(conn *grpc.ClientConn) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isClosed {
		return ErrPoolClosed
	}

	// 这里可以添加连接健康检查逻辑
	return nil
}

// Close 关闭连接池
func (p *GrpcConnPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isClosed {
		return
	}

	p.isClosed = true
	for i, conn := range p.conns {
		if conn != nil {
			conn.Close()
			p.conns[i] = nil
		}
	}
	p.conns = nil
}

// Status 获取连接池状态
func (p *GrpcConnPool) Status() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	return fmt.Sprintf("Pool{host: %s, size: %d, current: %d, closed: %v}",
		p.host, len(p.conns), p.current, p.isClosed)
}

// IsHealthy 检查连接池是否健康
func (p *GrpcConnPool) IsHealthy() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isClosed {
		return false
	}

	return len(p.conns) > 0
}

// NewGrpcPoolManager 创建新的连接池管理器
func NewGrpcPoolManager() *GrpcPoolManager {
	return &GrpcPoolManager{
		pools: make(map[string]*GrpcConnPool),
	}
}

// GetPool 获取或创建指定host的连接池
func (pm *GrpcPoolManager) GetPool(host string, opts ...PoolOption) (*GrpcConnPool, error) {
	// 先���试读锁获取
	pm.mu.RLock()
	if pool, exists := pm.pools[host]; exists {
		pm.mu.RUnlock()
		return pool, nil
	}
	pm.mu.RUnlock()

	// 如果不存在，使用写锁创建
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// 双重检查
	if pool, exists := pm.pools[host]; exists {
		return pool, nil
	}

	// 使用可配置的参数创建连接池
	poolConfig := defaultPoolConfig()
	for _, opt := range opts {
		opt(poolConfig)
	}

	pool, err := NewGrpcConnPool(host, poolConfig.size)
	if err != nil {
		return nil, fmt.Errorf("创建连接池失败 [%s]: %v", host, err)
	}

	pm.pools[host] = pool
	return pool, nil
}

// ClosePool 关闭指定host的连接池
func (pm *GrpcPoolManager) ClosePool(host string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pool, exists := pm.pools[host]; exists {
		pool.Close()
		delete(pm.pools, host)
	}
}

// CloseAllPools 关闭所有连接池
func (pm *GrpcPoolManager) CloseAllPools() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for host, pool := range pm.pools {
		pool.Close()
		delete(pm.pools, host)
	}
}

// ListPools 列出所有连接池状态
func (pm *GrpcPoolManager) ListPools() map[string]string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	status := make(map[string]string)
	for host, pool := range pm.pools {
		status[host] = pool.Status()
	}
	return status
}
