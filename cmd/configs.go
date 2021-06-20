package main

var defaultConfigs = struct {
	httpListen    string `env:"HTTP_LISTEN" flag:"http-listen" flagUsage:"http server listen port"`
	redisAddr     string `env:"REDIS_ADDR" flag:"redis-addr" flagUsage:"cache redis addr"`
	redisPassword string `env:"REDIS_PASSWORD" flag:"redis-password" flagUsage:"cache redis password"`
	run           string `flag:"run" flagUsage:"run what"`
	mode          string `flag:"mode" flagUsage:"run mode"`
	tmpJob        string `flag:"tmpjob" flagUsage:"tmpjob command"`
}{
	httpListen:    "0.0.0.0:12380",
	redisAddr:     "127.0.0.1:6379",
	redisPassword: "",
	run:           "http",
	mode:          "prod",
}
