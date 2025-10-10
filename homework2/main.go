package main

import (
	"homeworkMtsGo/homework2/server"
	"homeworkMtsGo/homework2/client"
	"time"
)

func main() {
	go server.Run()
	time.Sleep(10 * time.Second)
	client.Run()
}