// Map Variable is just a concept of Key-Value pair in Go. 
package main

import "fmt"

func main() {
	// Create Map by var declaration. -> make(map[keytype]valuetype) 
	var stat = make(map[string]int64) 

	// Add Key-Value Pair. -> mapname[key] = value 
	stat["atk"] = 100 
	stat["def"] = 200

	// Update Key-Value Pair. -> mapname[existedkey] = newvalue 
	stat["atk"] = 250

	// Delete Key-Value Pair. -> delete(mapname, key) 
	delete(stat, "def") 

	// Print Map. 
	fmt.Println(stat) 

	// Get Value from Key -> ValueVariable = mapname[key] 
	atkvalue := stat["atk"]
	fmt.Println(atkvalue)

	// ================================================================= // 

	// Create Map by Short Variable Declaration. -> MapVariable := map[keytype]valuetype{key1:value1, key2:value2}  
	charclass := map[string]string{"Daiso": "Warrior"}

	// Print Map. 
	fmt.Println(charclass) 
}