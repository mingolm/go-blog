package configs

var DefaultConfigs = struct {
	HttpListen    string `env:"HTTP_LISTEN" flag:"http-listen" flagUsage:"http server listen port"`
	DatabaseDsn   string `env:"DATABASE_DSN" flag:"database-dsn" flagUsage:"database dsn" apollo:"database.dsn"`
	RedisAddr     string `env:"REDIS_ADDR" flag:"redis-addr" flagUsage:"cache redis addr"`
	RedisPassword string `env:"REDIS_PASSWORD" flag:"redis-password" flagUsage:"cache redis password"`
	Run           string `flag:"run" flagUsage:"run what"`
	Mode          string `flag:"mode" flagUsage:"run mode"`
	TmpJob        string `flag:"tmpjob" flagUsage:"tmpjob command"`
}{
	HttpListen:    "0.0.0.0:12380",
	DatabaseDsn:   "root:123456@(localhost:3306)/recharge?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai",
	RedisAddr:     "127.0.0.1:6379",
	RedisPassword: "",
	Run:           "http",
	Mode:          "prod",
}
