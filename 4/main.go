package main

import (
	"fmt"
	"time"
)

type Token struct{
	data string
	recipient int
	ttl int
}

func SendToken(tc chan Token, recipient int) {
	tokenstate := <-tc
	fmt.Println(tokenstate.ttl)
	if tokenstate.ttl != 0 && tokenstate.recipient != recipient {
		tokenstate.ttl -= 1
		tc <- tokenstate
		time.Sleep(time.Second * 1)
	} else if tokenstate.recipient == recipient {
		fmt.Println(tokenstate.data)
	} else {
		fmt.Println("Error")
	}
}

func main() {
	var n int
	var token Token
	fmt.Scanf("%s\n", &token.data)
	fmt.Scanf("%d\n", &token.recipient)
	fmt.Scanf("%d\n", &token.ttl)
	fmt.Scanf("%d\n", &n)

	tokenchan := make(chan Token)
	for i := 0; i < n; i++ {
		go SendToken(tokenchan, i)
	}

	tokenchan <- token

	var input string
	fmt.Scanln(&input)
}