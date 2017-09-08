package container

import (
	"app/providers"
	"github.com/limingxinleo/di"
	"fmt"
)

var DI di.Context

func GetInstance() di.Context {
	// Create a Builder with the default scopes.
	builder, _ := di.NewBuilder()

	providers.BuildConfigProvider(builder)
	providers.BuildLoggerProvider(builder)

	DI = builder.Build()
	fmt.Println("DI Build")
	return DI
}


