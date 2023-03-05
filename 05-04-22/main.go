package main

import "log"

func main() {

	log.Println(add(2, 3))

}

func add(a int, b int) int {
	return a + b
}
