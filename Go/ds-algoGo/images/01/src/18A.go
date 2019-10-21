package main

import (
  "fmt"
)

type IProcess interface {
  process()
}
//adapter struct
type Adapter struct {
  adaptee Adaptee
}
