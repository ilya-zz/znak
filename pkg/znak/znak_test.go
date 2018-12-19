package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAdd(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(100, add(50, 50), "50 + 50 == 100")
}
func testMul(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(99, add(33, 3), "33 * 3 == 99")
}
