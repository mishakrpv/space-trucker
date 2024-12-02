package config

type Cfg struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`

	Log      *Log      `yaml:"log"`
	Postgres *Postgres `yaml:"postgres"`
	Redis    *Redis    `yaml:"redis"`
}
