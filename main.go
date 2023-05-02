package main

import (
	"fmt"
	"log"
	"github.com/rubensseva/go-dark-forest/system"
)

func main() {
	log.Printf(fmt.Sprintf("Hello, World!"))
    systems := []system.System{}
    for i := 0; i < 100; i++ {
        systems = append(systems, system.GenSystem(systems))
    }
    for _, s := range systems {
        fmt.Println(s)
    }
}
