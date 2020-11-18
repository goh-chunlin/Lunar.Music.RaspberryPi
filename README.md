# Lunar.Music.RaspberryPi

<div align="center">
    <img src="https://gclstorage.blob.core.windows.net/images/Lunar.Music.RaspberryPi-banner.png" />
</div>

![Go Build](https://github.com/goh-chunlin/Lunar.Music.RaspberryPi/workflows/Go%20Build/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/goh-chunlin/Lunar.Music.RaspberryPi)](https://goreportcard.com/report/github.com/goh-chunlin/Lunar.Music.RaspberryPi)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Donate](https://img.shields.io/badge/$-donate-ff69b4.svg)](https://www.buymeacoffee.com/chunlin)

This is a backend of Lunar Music project on Raspberry Pi to play music according to command sent from the Lunar.Music.Web.

This Golang application is designed as a music player hosted on a Raspberry Pi. It has a listener which will receive message from RabbitMQ server and perform the following actions.
- Play a music
  - The music needs to be in MP3 format;
  - The music file needs to be stored online
    Since this project is designed to work together with [Lunar.Music.Web](https://github.com/goh-chunlin/Lunar.Music.Web) which is using Microsoft OneDrive as the storage, the music files are originally all stored on the Microsoft OneDrive Music folder. Before playing a music, this program will download the MP3 file from Microsoft OneDrive and then keep it locally so that subsequent plays of the same music file can be loaded locally at `/home/pi/audio` on the Raspberry Pi without downloading it again.
- Play all music
  - This will play all the MP3 files which are stored in the `/home/pi/audio`.
  
## How to use? ##

1. Install Golang on the Raspberry Pi;
1. Clone this project to Raspberry Pi;
1. Install necessary Golang packages;
1. Create a `.env` file in the `app` directory with the following content;
   ```
   RABBITMQ_SERVER_CONNECTION_STRING=
   RABBITMQ_CHANNEL_NAME=
   ```
   
   If you are using this project together with [Lunar.Music.Web](https://github.com/goh-chunlin/Lunar.Music.Web), then the RabbitMQ part above must be same as the ones defined in the `.env` file in Lunar.Music.Web so that the communication between the Raspberry Pi and the web app can work;
1. Build the go web project in the root directory of this project;
1. Run the output program.
   Optionally, you can set this program to be run automatically when Raspberry Pi is connected so that you don't have to manually run the program to make the music player work.

## Contributing ##
First and foremost, thank you! I appreciate that you want to contribute to this project which is my personal project. Your time is valuable, and your contributions mean a lot to me. You are welcomed to contribute to this project development and make it more awesome every day.

Don't hasitate to contact me, open issue, or even submit a PR if you are intrested to contribute to the project.

Together, we learn better.

## License ##

This library is distributed under the GPL-3.0 License found in the [LICENSE](./LICENSE) file.
