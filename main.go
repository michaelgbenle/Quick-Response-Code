package main

import (
	"io/ioutil"
	"os"
)

func main() {
	msg, err := os.Open("myFile.txt")
	if err != nil{
		panic(err)
	}
	defer msg.Close()
	message, err := ioutil.ReadAll()

}