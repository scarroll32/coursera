package main  

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Person struct {
  firstName string
  lastName string
}

func readfilename() string {
  input := bufio.NewScanner(os.Stdin)
  fmt.Printf("Enter filename: ")
  for input.Scan(){
    _, err := os.Stat(input.Text())
    if err != nil {
      fmt.Printf("Error: %s\n",err)
      fmt.Printf("Enter filename: ")
    } else {
      break
    }
  }
  return input.Text()
}

func processFile(filename string) {
  var people []Person
  var firstName, lastName string
  const FIELDSIZE = 20

  fmt.Printf("processing file: %s\n", filename)
  fileHandle, _ := os.Open(filename)
	inputStream := bufio.NewReader(fileHandle)

	for true {
		line, err := inputStream.ReadString('\n')
		lineSlice := strings.Split(strings.Replace(line, "\n", "", -1), " ")

    switch true {
      case len(lineSlice) == 1:
        firstName = lineSlice[0]
  			lastName = ""
      case len(lineSlice) > 1:
        firstName = lineSlice[0]
        lastName = lineSlice[1]
      default:
        break
    }

		if len(firstName) > FIELDSIZE {
			firstName = firstName[0:FIELDSIZE]
		}

		if len(lastName) > FIELDSIZE {
			lastName = lastName[0:FIELDSIZE]
		}

		people = append(people, Person{
			firstName: firstName,
			lastName: lastName,
		})

		if err != nil {
			break
		}
	}

  fileHandle.Close()
  for i, v := range people {
    fmt.Printf("%d) First Name: '%s', Last Name: '%s'\n", i+1, v.firstName, v.lastName)
	}
}

func main() {
  filename := readfilename()
  processFile(filename)
}