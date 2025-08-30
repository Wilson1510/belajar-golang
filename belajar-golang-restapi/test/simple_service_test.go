package test

import (
	"testing"

	"belajar-golang-restapi/simple"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceNoError(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}