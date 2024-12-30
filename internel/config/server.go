package config

type Server struct {
	System System `json:"system"`
}

type System struct {
	Addr          string   `json:"addr" yaml:"addr"`
	Grpc          string   `json:"grpc" yaml:"grpc"`
	Logfile       string   `json:"logfile" yaml:"logfile"`
	Serct         string   `json:"serct"yaml:"serct"`
	Database      Database `json:"database" yaml:"database"`
	Redisaddr     string   `json:"redisaddr" yaml:"redisaddr"`
	Agent_port    string   `json:"agent_port"yaml:"agent_port"`
	Times         string   `json:"times" yaml:"times"`
	Agent_outtime int      `json:"agent_outtime" yaml:"agent_outtime"`
}
