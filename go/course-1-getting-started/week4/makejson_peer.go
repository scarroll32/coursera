package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var userName, userAddr string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a name:")
	scanner.Scan()
	userName = scanner.Text()
	fmt.Println("Enter an address:")
	scanner.Scan()
	userAddr = scanner.Text()

	m := make(map[string]string)
	m["Name"] = userName
	m["Address"] = userAddr

	barr, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(barr))
	}
}

