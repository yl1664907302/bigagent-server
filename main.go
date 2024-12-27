package main

import (
	_ "bigagent_server/docs"
	"bigagent_server/inits"
	"context"
	"os/signal"
	"syscall"
)

func init() {
	inits.Viper()
	inits.Logger()
	inits.MysqlDB()
	inits.RedisDB()
}

// @title BigAgent API
// @version 1.0
// @description This is a BigAgent API server.
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	inits.CronTask()
	inits.RunG()
	inits.Run(ctx, inits.Router)
}
