package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"White": "#FFFFFF",
	}

	var colorsWithVar map[string]string

	colorsMake := make(map[string]string)

	colorsMake["White"]= "#FFFFFF"
	colorsMake["Blue"]= "#0000FF"

	delete(colorsMake,"White")

	fmt.Println(colors)
	fmt.Println(colorsWithVar)
	fmt.Println(colorsMake)
	fmt.Println("--------Iterating Maps------")
	printMap(colors)
}

func printMap (c map[string]string){
	for color,hexCode:=range c{
		fmt.Println("Hex code for",color, "is",hexCode)
	}
}

