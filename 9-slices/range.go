package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, v := range msg {
		for _, badWord := range badWords {
			if v == badWord {
				return i
			}
		}
	}
	return -1
}

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test(message, badWords)
}