package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Ben"
	expected := "hello Ben"
	actual := getHelloString(name)

	assert.Equal(t, expected, actual)
}
