package configs

var DefaultConfigs = struct {
	HttpListen    string `env:"HTTP_LISTEN" flag:"http-listen" flagUsage:"http server listen port"`
	DatabaseDsn   string `env:"DATABASE_DSN" flag:"database-dsn" flagUsage:"database dsn" apollo:"database.dsn"`
	RedisAddr     string `env:"REDIS_ADDR" flag:"redis-addr" flagUsage:"cache redis addr"`
	RedisPassword string `env:"REDIS_PASSWORD" flag:"redis-password" flagUsage:"cache redis password"`

	// 四方配置
	PAYKey              string `env:"PAY_KEY" flag:"pay-key" flagUsage:"key for pay"`
	PAYPageUrl          string `env:"PAY_PAGE_URL" flag:"pay-page-url" flagUsage:"page url for pay"`
	PAYBGUrl            string `env:"PAY_BG_URL" flag:"pay-bg-url" flagUsage:"bg url for pay"`
	PAYH5RemoteAddr     string `env:"PAY_H5_REMOTE_ADDR" flag:"pay-h5-remote-addr-url" flagUsage:"h5 remote addr url for pay"`
	PAYQRCodeRemoteAddr string `env:"PAY_QRCODE_REMOTE_ADDR" flag:"pay-qrcode-remote-addr-url" flagUsage:"qrcode remote addr url for pay"`

	Run    string `flag:"run" flagUsage:"run what"`
	Mode   string `flag:"mode" flagUsage:"run mode"`
	TmpJob string `flag:"tmpjob" flagUsage:"tmpjob command"`
}{
	HttpListen:    "0.0.0.0:12380",
	DatabaseDsn:   "root:123456@(localhost:3306)/recharge?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai",
	RedisAddr:     "127.0.0.1:6379",
	RedisPassword: "",
	Run:           "http",
	Mode:          "prod",
}
