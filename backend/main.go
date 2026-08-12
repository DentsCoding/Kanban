package main

import "fmt"

type Greeter interface {
	Greet() string
	ReturnOrigin() string
}

type SpanishGreeter struct {
	From string
}

type EnglishGreeter struct {
	From string
}

func (g SpanishGreeter) Greet() string {
	return "Hola Mundo!"
}

func (g SpanishGreeter) ReturnOrigin() string {
	return g.From
}

func (g EnglishGreeter) ReturnOrigin() string {
	return g.From
}

func (g EnglishGreeter) Greet() string {
	return "Hello world!"
}

func BehaveGreetingly(g Greeter) string {
	if g.ReturnOrigin() == "Spain" {
		return "Hola mundo!"
	} else {
		return "Hello world!"
	}
}

func main() {
	spanishGreeter := SpanishGreeter{"Spain"}
	englishGreeter := EnglishGreeter{"US"}

	fmt.Println("English greeter says: ", BehaveGreetingly(englishGreeter), "Spanish greeter says: ", BehaveGreetingly(spanishGreeter))
}
