package main

import "fmt"

// logData logs any type of data using the empty interface
func logData(data any) {
    fmt.Println("Data logged: ", data)
}

func main() {
    logData("Hello, world!")  
    logData(12345)          
}
