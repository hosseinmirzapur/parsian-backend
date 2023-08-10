package config

type appConfig struct {
	dev bool
}

type AppModes string

const (
	dev  AppModes = "development"
	prod AppModes = "production"
)

func NewConfig(status AppModes) *appConfig {

	if status == prod {
		return &appConfig{
			dev: false,
		}
	}

	return &appConfig{
		dev: true,
	}

}

func (cfg *appConfig) IsDevelopment() bool {
	return cfg.dev
}
