package main

import (
  "fmt"
)

func main () {
  ch := make(chan string)

  defer close(ch)

  go task(ch)
  go task(ch)
  go task(ch)
  go task(ch)

  s := <- ch
  fmt.Println(s)
}

func task(ch chan<- string) {
  ch <- "I'm a goroutine"
}

