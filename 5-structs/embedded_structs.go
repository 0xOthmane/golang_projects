package main

import "fmt"




type sender struct {
	// "user" is embedded, so the definition of a
  	// "sender" now also additionally contains all
  	// of the fields of the user struct
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}


func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}