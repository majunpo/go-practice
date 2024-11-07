package main

import (
    "fmt"
    "github.com/sirupsen/logrus"
)

func main(){
    fmt.Println("Hello, Go Module!")
    logrus.Info("This is a long message from logrus")
}
