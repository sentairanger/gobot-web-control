# gobot-web-control
This is a Golang web application that runs a robot using the CamJam Edukit 3 motor board.


## Introduction

This project uses Golang web development along with HTML, CSS and jQuery to control a Raspberry Pi robot using the CamJam Edukit 3 motor board. However, this can be altered to use other motor boards like the L298N or even the L293D motor module.

## Getting started

First, install Go using the instructions found [here](). Make sure to use the correct download depending on your Pi. Next, clone this repository and then make sure to check that Go was installed by running `go version`. Next, you can either build the image using `go build` and then running it with `./gobot-web-control` or you can directly run the application with `go run .`. Once you run the application wait for the "Robot ready to go" message to pop up and then you can go to the IP of your Pi and then to port 6111 using the link `<ip-address-of-pi>:6111` and then you should see the application appear. Then you can move the robot with the buttons and when you let go of the buttons the robot stops. The application uses jQuery to handle movement.

## Picture of Application

![GoBot](https://github.com/sentairanger/gobot-web-control/blob/main/gobot.png)
