package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

type Italian struct {}
type Portuguese struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
