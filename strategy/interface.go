package strategy

// ApiStrategy 定义api接口策略
type ApiStrategy interface {
	Api(key string) (interface{}, error)
}

// PushStrategy 定义推送接口策略
type PushStrategy interface {
	Push() error
}

type CmdbServe struct {
	apiStrategy  ApiStrategy
	pushStrategy PushStrategy
}

func NewCmdbServe() CmdbServe {
	return CmdbServe{}
}

func (a *CmdbServe) SetApiStrategy(strategy ApiStrategy) {
	a.apiStrategy = strategy
}

func (a *CmdbServe) SetPushStrategy(strategy PushStrategy) {
	a.pushStrategy = strategy
}

func (a *CmdbServe) ExecuteApi(key string) (interface{}, error) {
	return a.apiStrategy.Api(key)
}

func (a *CmdbServe) ExecutePush() error {
	return a.pushStrategy.Push()
}

var CmdbServes []CmdbServe
