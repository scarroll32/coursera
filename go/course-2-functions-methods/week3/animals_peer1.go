package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	foodEaten string
	movement string
	sound string
}

func (animal Animal) Speak() string {
	return animal.sound
}

func (animal Animal) Eat() string {
	return animal.foodEaten
}

func (animal Animal) Move() string {
	return animal.movement
}

func NewAnimal(sound, food, movement string) *Animal {
	animal := new(Animal)
	animal.foodEaten = food
	animal.movement = movement
	animal.sound = sound
	return animal
}

var animalMap = setupAnimalsAvailable()

func main() {
	for {
		fmt.Printf("Please enter a request (cow, bird or snake followed by eat, move or speak  >")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var command= scanner.Text()
		var animalAndBehavior= strings.Split(command, " ")
		if len(animalAndBehavior) == 2 {
			fmt.Println(formulateResponse(strings.ToLower(animalAndBehavior[0]), strings.ToLower(animalAndBehavior[1])))
		} else {
			printUsage()
		}
	}
}

func setupAnimalsAvailable() map[string]Animal {
	var cow = NewAnimal("moo", "grass", "walk")
	var bird = NewAnimal("peep", "worms", "fly")
	var snake = NewAnimal("hsss", "mice", "slither")
	return map[string]Animal{"cow":*cow, "bird":*bird, "snake":*snake}

}

func formulateResponse(animalType string, behavior string) string {
	var animalRequested = animalMap[animalType]

	var response string
	switch behavior {
	case "eat":
			response = animalRequested.Eat()
	case "move":
		response = animalRequested.Move()
	case "speak":
		response = animalRequested.sound // just testing.
	default:
		response = "Possible behaviors are eat, move and speak.  " + behavior + " is not one of these types."
	}
	if response == "" {
		response = "Valid animals are cow, bird and snake.  " + animalType + " is not one of these animals."
	}
	return response
}

func printUsage() {
	fmt.Println("Your request was not properly formed. " +
		"Please enter an animal type followed by a space and then a behavior.  " +
		"E.g., `cow eat`  Possible behaviors are eat, move or speak.  Animal types are cow, bird or snake.")
}
