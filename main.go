package main

import (
	"io/ioutil"
	"os"

	"github.com/skip2/go-qrcode"
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

	err = qrcode.WriteFile(message, qrcode.Medium,512,myQr)
	if err != nil {
		panic(err)
	}

}