package main

import "fmt"

// Country struct goes here
type Country struct{
  name string
  capital string
  latitude, longitude float32
}

func main() {
	var france Country
  fmt.Println(france)
}
