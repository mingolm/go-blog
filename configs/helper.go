package configs

func IsProd() bool {
	return DefaultConfigs.Mode == "prod"
}
