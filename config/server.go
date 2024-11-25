package config

type Server struct {
	System System `json:"system"`
}

type System struct {
	Addr        string   `json:"addr" yaml:"addr"`
	Grpc        string   `json:"grpc" yaml:"grpc"`
	Logfile     string   `json:"logfile" yaml:"logfile"`
	Serct       string   `json:"serct"yaml:"serct"`
	Client_port string   `json:"client_port" yaml:"client_port"`
	Database    Database `json:"database" yaml:"database"`
}
