package main

import (
  "fmt"
  "encoding/json"
  "bufio"
  "os"
)

func readData(fieldType string) string {
  input := bufio.NewScanner(os.Stdin)
  fmt.Printf("Enter your %s: ", fieldType)
  for input.Scan(){
    if input.Text() != "" {
      break
    } else {
      fmt.Printf("Enter your %s: ", fieldType)
    }
  }
  return input.Text()
}


func main(){
  userDetails := make(map[string]string)
  userDetails["name"] = readData("name")
  userDetails["address"] = readData("address")
  fmt.Printf("Name: %s\n", userDetails["name"])
  fmt.Printf("Address: %s\n", userDetails["address"])

  json, _ := json.Marshal(userDetails)
  fmt.Printf("JSON: %s",json)
}