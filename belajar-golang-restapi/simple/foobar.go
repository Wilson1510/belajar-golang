package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(foo *FooService, bar *BarService) *FooBarService {
	return &FooBarService{FooService: foo, BarService: bar}
}

func A() string {
	return "A"
}

func B() int {
	return 2
}
