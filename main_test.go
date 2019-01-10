package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	regards := getRegards()
	assert.Equal(t, "Hello little man", regards)
}
