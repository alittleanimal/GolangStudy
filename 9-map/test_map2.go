package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Print("key = ", key)
		fmt.Println(" value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	delete(cityMap, "China")
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)
	fmt.Println("----------")

	printMap(cityMap)
}
