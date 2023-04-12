package singleton_test

import (
	"sdp/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionWithoutSingleton(t *testing.T) {
	var instance []any

	for i := 0; i < 1000; i++ {
		instance = append(instance, singleton.GetConnectionInstanceWithoutSingleton())
	}

	for i := 0; i < 1000; i++ {
		assert.Samef(t, instance[0], instance[i], "Duplicate instance")
	}
}

func TestConnectionWithSingleton(t *testing.T) {
	var instance []any

	for i := 0; i < 1000; i++ {
		instance = append(instance, singleton.GetConnectionInstance())
	}

	for i := 0; i < 1000; i++ {
		assert.Samef(t, instance[0], instance[i], "Duplicate instance")
	}
}
