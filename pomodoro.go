package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

/*
USAGE: go run pomodoro.go <number of minutes>

SAMPLE:
[kishansagathiya@localhost utility]$ go run pomodoro.go 60
Running pomodoro for 60 minutes
59
58
*/
func main() {

	argsWithProg := os.Args
	i, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Running pomodoro for %d minutes\n", i)

	go func() {
		for {
			time.Sleep(40 * time.Minute)
			err = exec.Command("vlc", "waterMsg.mp3").Run()
		}
	}()

	for i > 0 {
		time.Sleep(1 * time.Minute)
		i--
		fmt.Println(i)
	}

	err = exec.Command("vlc", "alarm.mp3").Run()
	if err != nil {
		log.Fatal(err)
	}

}
