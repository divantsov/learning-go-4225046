package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	//var colors = [3]string{"Red", "Green", "Blue"}
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	colors = append(colors, "Purple", "Fuschia")
	
	colors = remove(colors, 1)

	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
