package main

import (
  "fmt"
  "time"
)

func main() {
  for {
    fmt.Println("Hello World 2")
    time.Sleep(time.Second * 3)
  }
}
