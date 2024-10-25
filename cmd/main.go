package main

import (
	"dnc/cmd/set"
	"fmt"
)

func main() {
	dnc := set.NewSet()
	dnc.Add("1231234444")
  fmt.Println(dnc.Size())
}
