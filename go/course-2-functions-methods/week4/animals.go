package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  )

type Animal struct {
  food, locomotion, sound string
}

type AnimalInterface interface {
  Eat()
  Move()
  Speak()
}

func (animal Animal) Eat(){
  fmt.Println(animal.food)
}
func (animal Animal) Move(){
  fmt.Println(animal.locomotion)
}
func (animal Animal) Speak(){
  fmt.Println(animal.sound)
}

func ArrayContains(a []string, x string) bool {
  for _, n := range a {
    if x == n {
      return true
    }
  }
  return false
}


func printUsage(enteredString string) {
  fmt.Println("Sorry, I cannot understand that command")
  fmt.Printf("You entered: %s\n\n", enteredString)
  fmt.Println("Valid commands are:")
  fmt.Println("   newanimal <animal name> <snake | cow | bird>")
  fmt.Println("   query <animal name> <eat | move | speak>")
  fmt.Println("")
}

func validateInput(userInput string) bool {
  var command []string

  command = strings.Split(userInput, " ")
  if len(command) != 3 {
    printUsage(userInput)
  }

  if ArrayContains(ValidCommands, command[0]){
    if command[0] == "query" {
      if ArrayContains(ValidActions, command[2]){
        return true
      } else {
        printUsage(userInput)
      }
    }
    if command[0] == "newanimal" {
      if ArrayContains(ValidAnimals, command[2]){
        return true
      } else {
        printUsage(userInput)
      }
    }
  } else {
    printUsage(userInput)
  }

  return false
}

func readPrompt() []string {
  input := bufio.NewScanner(os.Stdin)
  fmt.Printf(">")
  for input.Scan(){
    if validateInput(input.Text()) {
      break
    } else {
      fmt.Printf(">")
    }
  }
  var command = strings.Split(input.Text(), " ")
  return command
}


var ValidAnimals = []string{"cow", "bird", "snake"}
var ValidActions = []string{"eat", "move", "speak"}
var ValidCommands = []string{"newanimal", "query"}

func main(){

  animals := map[string]Animal{
    "cow": Animal{"grass", "walk", "moo"},
    "bird": Animal{"worms", "fly", "peep"},
    "snake": Animal{"mice", "slither", "hsss"},
  }

  var ani AnimalInterface

  for {
    var prompt = readPrompt()
    var command, animal, requestType = prompt[0], prompt[1], prompt[2]
		if command == "query" {
			ani = animals[animal]
			switch requestType {
			case "eat":
				ani.Eat()
			case "move":
				ani.Move()
			case "speak":
				ani.Speak()
			}
		}
		if command == "newanimal" {
			animals[animal] = animals[requestType]
			fmt.Println("Created it!")
		}
  }

}
