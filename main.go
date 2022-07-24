package main

import "os"

func main() {
	message, err := os.Open("myFile.txt")
	if err != nil{
		panic(err)
	}
	defer message.Close()

}