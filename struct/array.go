package main

type Array struct {
  Data []int
}

func (a *Array) Add(value int) {
  a.Data = append(a.Data, value)
}

func (a *Array) Remove(index int) {
  a.Data = append(a.Data[:index], a.Data[index+1:]...)
}

func (a *Array) Find(value int) int {
  for _, v := range a.Data {
    if v == value {
      return v
    }
  }
  return -1
}

func (a *Array) FindIndex(value int) int {
  for n, v := range a.Data {
    if v == value {
      return n
    }
  }
  return -1
}
