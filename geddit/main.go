package main

import (
	"fmt"	
	"log"
	"github.com/markmilec/reddit"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
