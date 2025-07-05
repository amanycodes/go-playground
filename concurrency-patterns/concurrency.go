package main

import "fmt"

func main() {

	// select loop pattern

	chnl1 := make(chan string)
	chnl2 := make(chan string)
	go func() {
		chnl1 <- "data"
	}()

	go func() {
		chnl2 <- "anotherData"
	}()

	select {
	case msg1 := <-chnl1:
		fmt.Println(msg1)
	case msg2 := <-chnl2:
		fmt.Println(msg2)
	}

	// for select loop

	charchnl := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		// select {
		// case charchnl <- s:
		// }
		charchnl <- s
	}

	close(charchnl)

	for result := range charchnl {
		fmt.Println(result)
	}
}
