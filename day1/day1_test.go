package day1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := "(())"
	assert.Equal(t, input, 0, "they should be equal")
}

func Test2(t *testing.T) {
	input := "()()"
	assert.Equal(t, input, 0, "they should be equal")
}

func Test3(t *testing.T) {
	input := "((("
	assert.Equal(t, input, 3, "they should be equal")
}

func Test4(t *testing.T) {
	input := "(()(()("
	assert.Equal(t, input, 3, "they should be equals")
}
