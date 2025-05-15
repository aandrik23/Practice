package main

import (
	"fmt"
	// "piscine"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
	//all of the are like fmt.Print(piscine.LastWord(" "))
}
