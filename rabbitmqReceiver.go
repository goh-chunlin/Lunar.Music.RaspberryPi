// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"

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

					isMusicFileDownloaded := isDriveItemDownloaded(driveItemId)

					if !isMusicFileDownloaded {
						err = downloadDriveItem(driveItemDownloadUrl, driveItemFileName)

						if err == nil {
							updateDownloadedDriveItemsList(driveItemId, driveItemFileName)
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
