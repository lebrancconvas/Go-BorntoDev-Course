// Map Variable is just a concept of Key-Value pair in Go. 
package main

import "fmt"

func main() {
	// Create Map by var declaration.
	var stat = make(map[string]int64) // make(map[keytype]valuetype) 

	// Add Key-Value Pair. -> mapname[key] = value 
	stat["atk"] = 100 
	stat["def"] = 200

	// Update Key-Value Pair. -> mapname[existedkey] = newvalue 
	stat["atk"] = 250

	// Print Map. 
	fmt.Println(stat) 
}