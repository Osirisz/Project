package fizzBuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test(t *testing.T,num int,m string) {
	assert.Equal(t,FizzBuzz(3), "fizz")
	assert.Equal(t,FizzBuzz(5), "buzz")
	assert.Equal(t,FizzBuzz(15), "fizzBuzz")
}
