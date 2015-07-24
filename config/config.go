package config

type Config interface {
	Save()
	Get() ConfigData
	Put(ConfigData)
}
