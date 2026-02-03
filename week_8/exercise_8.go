package main
import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Keith")
}

func changeLastElement(slice []string, newName string){
    slice[len(slice)-1] = newName
    fmt.Println(slice)
}
