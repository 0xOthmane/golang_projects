package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func main() {
	fmt.Println(concat("Lane,", "happy birthday!"))
	fmt.Println(concat("Elon,", "hope that Tesla thing works out"))
	test("Go", "is fantastic")
	welcome()
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
// When multiple arguments are of the same type, the type only needs to be declared after the last one, assuming they are in order.
func add(x, y int) int {
	return x + y
  }

// Ignore a return value

func getNames() (string, string) {
	return "First", "Last"
}

func welcome() {
	first, _ := getNames()
	if first != "" {
		fmt.Println("Welcome " + first)
	}
}