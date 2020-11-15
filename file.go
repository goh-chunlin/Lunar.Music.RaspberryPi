// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func downloadDriveItem(downloadUrl string, fileName string) error {
	cmd := exec.Command("wget", "-O", fileName, downloadUrl)

	cmd.Dir = "/home/pi/audio"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func isDriveItemDownloaded(driveItemId string) bool {
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

	return isMusicFileDownloaded
}

func updateDownloadedDriveItemsList(driveItemId string, driveItemFileName string) {
	dataFile, err := os.OpenFile("playlist.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		failOnError(err, "Failed to open the file playlist.dat")
	}

	defer dataFile.Close()

	dataFile.Write([]byte(fmt.Sprintf("%v##########%v\n", driveItemId, driveItemFileName)))
}
