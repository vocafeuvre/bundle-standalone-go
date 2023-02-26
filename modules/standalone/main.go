package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := Asset("../gui/bin/gui")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Write byte slice to file
	err = ioutil.WriteFile("output", data, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data written to file")
}
