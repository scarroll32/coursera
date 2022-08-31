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

func (animal Animal) Eat(){
  fmt.Println(animal.food)
}
func (animal Animal) Move(){
  fmt.Println(animal.locomotion)
}
func (animal Animal) Speak(){
  fmt.Println(animal.sound)
}

var ValidAnimals = []string{"cow", "bird", "snake"}
var ValidActions = []string{"eat", "move", "speak"}


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


func printUsage(enteredString string) {
  fmt.Println("Sorry, I cannot understand that command")
  fmt.Printf("You entered: %s\n\n", enteredString)
  fmt.Println("[snake | cow | bird ] [ eat | move | speak]")
}


func validateInput(userInput string) bool {
  var command = strings.Split(userInput, " ")
  var valid = true

  if len(command) != 2 {
    valid = false
  }

  if valid == true && !(ArrayContains(ValidAnimals, command[0])){
    valid = false
  }

  if valid == true && !(ArrayContains(ValidActions, command[1])){
    valid = false
  }

  if valid == false {
    printUsage(userInput)
  }
  return valid
}


func ArrayContains(a []string, x string) bool {
  for _, n := range a {
    if x == n {
      return true
    }
  }
  return false
}


func main(){
  var animal string
  var action string

  animals := map[string]Animal {
    "cow" : Animal{"grass","walk","moo"},
    "bird" : Animal{"worms","fly","peep"},
    "snake" : Animal{"mice","slither","hsss"},
  }

  for {
    command := readPrompt()
    animal = command[0]
    action = command[1]

    switch action {
      case "eat":
        animals[animal].Eat()
      case "move":
        animals[animal].Move()
      case "speak":
        animals[animal].Speak()
    }
  }
}
