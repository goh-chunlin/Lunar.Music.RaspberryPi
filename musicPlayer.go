// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var isMusicPlaying bool = false

func playAllMusicFiles() {
	if !isMusicPlaying {
		isMusicPlaying = true

		var songFiles []string

		songFileRoot := "/home/pi/audio"

		filepath.Walk(songFileRoot, func(path string, info os.FileInfo, _ error) error {
			if filepath.Ext(path) == ".mp3" {
				songFiles = append(songFiles, path)
			}
			return nil
		})

		for _, song := range songFiles {

			cmd := exec.Command("nvlc", song, "--play-and-exit")

			cmd.Dir = songFileRoot
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", err)
			}
		}

		isMusicPlaying = false
	}
}

func playSingleMusicFile(fileName string) {
	if !isMusicPlaying {
		isMusicPlaying = true

		songFileRoot := "/home/pi/audio"

		cmd := exec.Command("nvlc", fileName, "--play-and-exit")

		cmd.Dir = songFileRoot
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}

		isMusicPlaying = false
	}
}
