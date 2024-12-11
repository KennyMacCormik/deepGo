package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	typeConstructors map[string]func() interface{}
}

func NewContainer() *Container {
	return &Container{typeConstructors: make(map[string]func() interface{})}
}

func (c *Container) RegisterType(name string, constructor any) {
	c.typeConstructors[name] = constructor.(func() interface{})
}

func (c *Container) Resolve(name string) (any, error) {
	fn, ok := c.typeConstructors[name]
	if !ok {
		return nil, fmt.Errorf("type %sis not registered", name)
	}
	if fn == nil {
		return nil, fmt.Errorf("no constructor registered for %s", name)
	}
	return fn(), nil
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() any {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() any {
		return &MessageService{}
	})

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)
}
