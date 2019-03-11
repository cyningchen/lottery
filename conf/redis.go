package conf

type RedisConfig struct {
	Host      string
	Port      int
	IsRunning bool
}

var RedisCacheList = []RedisConfig{
	{
		Host:      "139.196.90.101",
		Port:      6379,
		IsRunning: true,
	},
}

var RedisCache = RedisCacheList[0]
