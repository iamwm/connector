package conf

type MongoConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MqConfig struct {
	Host        string   `yaml:"host"`
	Port        string   `yaml:"port"`
	User        string   `yaml:"user"`
	Password    string   `yaml:"password"`
	Exchange    string   `yaml:"exchange"`
	Queue       string   `yaml:"queue"`
	RoutingKeys []string `yaml:"routing_keys"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   int    `yaml:"db"`
}
