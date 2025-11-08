package main

import "fmt"

func main() {
	// There isn't much about practicing maps
	// but I'll make sure to get a lot out
	// of it in a future learning project

	hashtable := make(map[string]int)

	// name: age
	hashtable["Yuri"] = 18
	hashtable["Caio"] = 17
	hashtable["Pedro"] = 17

	hashtable["Pedro"] = 92

	delete(hashtable, "Pedro")

	fmt.Println(hashtable)
}
