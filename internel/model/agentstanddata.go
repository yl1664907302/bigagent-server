package model

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

// AgentStandData 暴露原生utils数据
type AgentStandData struct {
	Secret  string       `json:"secret"`
	Uuid    string       `json:"uuid"`
	Cpu     *CpuInfo     `json:"cpu_info"`
	Mem     *MemInfo     `json:"mem_info"`
	Disk    *DiskInfo    `json:"disk_info"`
	Process *ProcessInfo `json:"progress"`
}

func NewAgentStandData() *AgentStandData {
	return &AgentStandData{}
}

// CpuInfo 计算
type CpuInfo struct {
	CpuInfo *cpu.InfoStat
	TimeCpu *cpu.TimesStat
}

// DiskInfo 磁盘
type DiskInfo struct {
	Usage      *disk.UsageStat
	Partition  *disk.PartitionStat
	IOCounters *disk.IOCountersStat
}

// MemInfo 内存
type MemInfo struct {
	VirMem     *mem.VirtualMemoryStat
	SwapMem    *mem.SwapMemoryStat
	SwapDevice *mem.SwapDevice
}

// ProcessInfo 进程
type ProcessInfo struct {
	Process []*process.Process
}
