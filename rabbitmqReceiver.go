// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type Command struct {
	Tasks []*Task `json:"tasks"`
}

type Task struct {
	Name    string   `json:"name"`
	Content []string `json:"content"`
}

func receiveAmqp(rabbitMQServerConnectionString string, rabbitMQChannelName string) {
	conn, err := amqp.Dial(rabbitMQServerConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitMQChannelName, // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			var command *Command
			json.Unmarshal(d.Body, &command)

			for _, task := range command.Tasks {
				if task.Name == "play-all" {
					playAllMusicFiles()
				} else if task.Name == "play-single" {
					driveItemId := task.Content[0]
					driveItemDownloadUrl := task.Content[1]
					driveItemFileName := task.Content[2]

					dataFile, err := os.OpenFile("playlist.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						failOnError(err, "Failed to open the file playlist.dat")
					}

					defer dataFile.Close()

					isMusicFileDownloaded := false
					scanner := bufio.NewScanner(dataFile)
					for scanner.Scan() {
						line := scanner.Text()

						if !isMusicFileDownloaded && 0 == strings.Index(line, driveItemId) {
							lineComponents := strings.Split(line, "##########")

							playSingleMusicFile(lineComponents[1])

							isMusicFileDownloaded = true
						}
					}

					if !isMusicFileDownloaded {
						err = downloadDriveItem(driveItemDownloadUrl, driveItemFileName)

						if err == nil {
							dataFile.Write([]byte(fmt.Sprintf("%v##########%v\n", driveItemId, driveItemFileName)))
							playSingleMusicFile(driveItemFileName)
						}
					}
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
