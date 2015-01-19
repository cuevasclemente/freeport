package freeport

import "testing"
import "fmt"
import "log"

func TestFind(t *testing.T){
  l, err := Find()
  if err != nil {
    log.Panic(err)
  }
  fmt.Println(l)
}
