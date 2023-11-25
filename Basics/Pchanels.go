package Basics

import (
	"fmt"
	"time"
)


func Pinger(c chan string) {
	for i := 0; ; i++ {
	  c <- "ping mama"
	  msg := <- c
	  fmt.Println(msg)
	}
  }
  
  func Printer(c chan string) {
	for {
	  msg := <- c
	  fmt.Println(msg)
	  time.Sleep(time.Second * 1)
	  c <- "ping bagina"
	}
  }