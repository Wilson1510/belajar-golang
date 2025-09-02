package simple

import (
	"github.com/google/wire"
)

type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{
			Name: "Local",
		},
	}
}

func InitializeConfiguration() *Configuration {	
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}
