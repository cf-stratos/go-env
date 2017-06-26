package main

import (
	"fmt"
	"time"
	"github.com/tjarratt/babble"
)

func main() {

	for {
		str, _ := getCharacters()
		fmt.Println(str)
	}
}

func getCharacters() (string, error) {

	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 3
	str := babbler.Babble()
	time.Sleep(1 * time.Second)
	return str, nil

}