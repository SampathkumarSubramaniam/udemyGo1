package main

import "fmt"

type englishBot struct {
	name string
}

type spanishBot struct {
	name string
}

type bot interface {
	getGreetings() string
}

func (englishBot) getGreetings() string {
	return "Hello My Friend!"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}

func printGreetings(b bot) {
	fmt.Printf(b.getGreetings()+"- from %v\n", b)
}
func main() {
	tom := englishBot{name: "Tom - English Bot"}
	jack := spanishBot{name: "jack - Spanish Bot"}
	printGreetings(tom)
	printGreetings(jack)

}
