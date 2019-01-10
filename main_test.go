package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
}

func TestMain_GetRegards(t *testing.T) {
	regards := getRegards()
	assert.Equal(t, "Hello little man", regards)
}
