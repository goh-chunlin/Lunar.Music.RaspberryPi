// Copyright 2020 The Lunar.Music.RaspberryPi AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
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
