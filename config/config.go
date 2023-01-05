package config

type Configuration interface {
	ParseConfig() error
}
