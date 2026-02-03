package main
import "fmt"

func main() {
  donuts := map[string]int{
    "frosted":   10,
    "chocolate": 15,
    "jelly":     8,
  }

  // Print out all the donuts
  fmt.Println(donuts)

  // Add your code here
  firstChoice := donuts["frosted"]
  fmt.Println(firstChoice)

  secondChoice,status := donuts["bavarian cream"]

  fmt.Println(status)
   
  if status {
    fmt.Println(secondChoice)
  } else {
    fmt.Println("We do not sell that donut!")
  }
}
