package config

type Server struct {
	System System `json:"system"`
}

type System struct {
	Addr string `json:"addr" yaml:"addr"`
	Grpc string `json:"grpc" yaml:"grpc"`
}
