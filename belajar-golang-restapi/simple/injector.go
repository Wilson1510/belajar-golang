//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgres,
		NewDatabaseRepository,
	)
	return nil
}

var FooSet = wire.NewSet(
	NewFooRepository,
	NewFooService,
)

var BarSet = wire.NewSet(
	NewBarRepository,
	NewBarService,
)

func InitializeFooBarService() *FooBarService {
	wire.Build(
		FooSet,
		BarSet,
		NewFooBarService,
	)
	return nil
}

var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

func InitializeAB() string {
	wire.Build(A)
	return ""
}

func InitializeFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}