package main 

import (
	"fmt"
	"log"
	"github.com/annuraagggIIIT/GO/greetings"
)
func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("a")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}