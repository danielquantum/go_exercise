package main

import "fmt"

func main() {
    myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
    fmt.Println(myTutors)
    myTutorSlice := myTutors[:]
    mySubjects := []string{"Go", "Java", "Python"}
    fmt.Println(myTutorSlice)
    fmt.Println(mySubjects)
}
