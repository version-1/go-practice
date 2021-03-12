package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	Hello()
	assert.Equal(t, true, true)
}

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(1, 1), 2)
	assert.Equal(t, Sum(3, 1), 4)
	assert.Equal(t, Sum(1, 3), 4)
	assert.Equal(t, Sum(1, -1), 0)
}

func TestIsUnderage(t *testing.T) {
	assert.Equal(t, IsUnderage(40), false)
	assert.Equal(t, IsUnderage(20), false)
	assert.Equal(t, IsUnderage(19), true)
	assert.Equal(t, IsUnderage(-19), true)
}

func TestPrintTen(t *testing.T) {
	assert.Equal(t, PrintTen(), true)
}

func TestFib(t *testing.T) {
	assert.Equal(t, Fib(0), 0)
	assert.Equal(t, Fib(1), 1)
	assert.Equal(t, Fib(2), 1)
	assert.Equal(t, Fib(3), 2)
	assert.Equal(t, Fib(4), 3)
	assert.Equal(t, Fib(5), 5)
	assert.Equal(t, Fib(10), 55)
}

func TestGetSeqNumberArray (t *testing.T) {
	assert.ElementsMatch(t, GetSeqNumArray(3), [...]int{0, 1, 2})
	assert.ElementsMatch(t, GetSeqNumArray(0), [...]int{})
}

func TestIsDuplicate (t *testing.T) {
	assert.Equal(t, IsDuplicate([]int{1, 2, 3, 4}), false)
	assert.Equal(t, IsDuplicate([]int{}), false)
	assert.Equal(t, IsDuplicate([]int{1, 2, 2}), true)
	assert.Equal(t, IsDuplicate([]int{2, 2, 2}), true)
}
