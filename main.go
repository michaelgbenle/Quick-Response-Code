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
	byteMessage, err := ioutil.ReadAll(msg)
	if err != nil {
		panic(err)
	}
	message := string(byteMessage)

	myQr := "myQr.png"

	err = qrcode.writeFile(message, qrcode.Medium)

}