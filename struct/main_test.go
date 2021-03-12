package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArray (t *testing.T) {
  a := Array{[]int{1, 2, 3}}
  assert.Equal(t, a.Data, []int{1, 2, 3})
}

func TestAdd (t *testing.T) {
  a := Array{[]int{}}
  a.Add(1)
  a.Add(2)
  a.Add(3)
  assert.Equal(t, a.Data, []int{1, 2, 3})
}

func TestRemove (t *testing.T) {
  a := Array{[]int{}}
  a.Add(1)
  a.Add(2)
  a.Add(3)
  a.Remove(1)
  assert.Equal(t, a.Data, []int{1, 3})
}

func TestFind (t *testing.T) {
  a := Array{[]int{}}
  a.Add(1)
  a.Add(2)
  a.Add(3)
  assert.Equal(t, a.Find(3), 3)
  assert.Equal(t, a.Find(5), -1)
}

func TestFindIndex (t *testing.T) {
  a := Array{[]int{}}
  a.Add(1)
  a.Add(2)
  a.Add(3)
  assert.Equal(t, a.FindIndex(1), 0)
  assert.Equal(t, a.FindIndex(3), 2)
  assert.Equal(t, a.Find(5), -1)
}

