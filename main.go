// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if _, err := net.Listen("tcp", ":8080"); err != nil {
		fmt.Println("An instance was already running")
		return
	}
	dtStart := time.Now()
	fmt.Println("Program starts at: ", dtStart.String())

	receiveAmqp(os.Getenv("RABBITMQ_SERVER_CONNECTION_STRING"), os.Getenv("RABBITMQ_CHANNEL_NAME"))

	dtEnd := time.Now()
	fmt.Println("Program ends at: ", dtEnd.String())
}
