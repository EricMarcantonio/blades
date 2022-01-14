package database

import (
	"fmt"
	"io/ioutil"
	"os"
)

func SeedFromFile() {
	c, ioErr := ioutil.ReadFile("./database/blades_skates.sql")
	if ioErr != nil {
		fmt.Println(ioErr)
		os.Exit(1)
	}
	_, err := DB.Query(string(c))
	if err != nil {
		fmt.Println(err)
	}
}
