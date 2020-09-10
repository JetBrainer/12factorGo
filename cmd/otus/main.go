package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main(){
	logrus.Info("Hello World")

	port := os.Getenv("PORT")
	if port == ""{
		logrus.Fatal("Port is not set")
	}


}