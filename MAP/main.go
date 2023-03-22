package main

import "fmt"

func main() {
	kunjal := make(map[string]string)
	kunjal["name"] = "Nadhira"
	kunjal["class"] = "P1E"

	kani := make(map[string]string)
	kani["name"] = "kanimozhi"
	kani["class"] = "Msc"
	sampath := map[string]string{
		"name":  "Sampath",
		"class": "Sap",
	}
	//fmt.Println(kunjal)
	delete(sampath, "class")
	printMap(kani)
}

func printMap(name map[string]string) {
	for key, value := range name {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
