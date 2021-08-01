package main

import (
	"fmt"
	"log"
	"github.com/google/go-cmp/cmp"

	"github.com/DevenKShah/GoCrash/src/morestrings"
	"github.com/DevenKShah/GoCrash/src/greetings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello world!"))
	fmt.Println((cmp.Diff("Hello world", "Hello Go")))

	log.SetPrefix("startup: ")

	message, err := greetings.Get("dev")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
