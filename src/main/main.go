package main

import "Blackjack/src/main/Controller"

func main() {
	err := Controller.Run()
	if err != nil {
		panic(err)
	}
}
