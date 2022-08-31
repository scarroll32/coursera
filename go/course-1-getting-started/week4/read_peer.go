package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	names := make([]Name, 0, 20)

	fmt.Println("Welcome! Please enter the name of the text file:")
	fileName, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fileName = strings.Trim(fileName, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		var aName Name
		aName.fname, aName.lname = s[0], s[1]
		names = append(names, aName)

	}

	file.Close()

	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}
}

