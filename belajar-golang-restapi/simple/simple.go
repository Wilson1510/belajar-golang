package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{Error: true}
}

type SimpleService struct {
	SimpleRepository *SimpleRepository
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("Error")
	} else {
		return &SimpleService{SimpleRepository: simpleRepository}, nil
	}
}
