package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
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
	aim, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
	}
	i := aim
	var exercised bool
	fmt.Println("Did you exercise for one hour?")
	fmt.Scan(&exercised)
	if exercised {
		fmt.Println("Yay, you are going great!")
	} else {
		fmt.Println("Tomorrow is another day, do better.")
	}

	tNow := time.Now()
	ontime := tNow.Hour() < 9
	if ontime {
		fmt.Println("It's great to be on time, you are going to have a great day!")
	} else {
		fmt.Printf("Not on time, You should be starting at 9AM, but it's %s\n\n", tNow.Format(time.Kitchen))
	}

	fmt.Printf("These are your goals:\n- Become a really really good programmer\n- Be fit\n- Be loved, wanted, charismatic, convincing, powerful and influencial\n\n")
	fmt.Printf("Running pomodoro for %d minutes\n", i)

	signalChan := make(chan os.Signal)
	go func() {
		for {
			time.Sleep(40 * time.Minute)
			err = exec.Command("vlc", "waterMsg.mp3").Run()
		}
	}()

	go func() {
		signal.Notify(
			signalChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)
		_, ok := <-signalChan // channel closed.
		if !ok {
			return
		}
		fmt.Printf("\nYou worked for %d minutes, but aim was %d minutes\n", aim-i, aim)
		logDaily(exercised, aim-i)
		os.Exit(1)
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

	logDaily(exercised, aim-i)
}

func logDaily(exercised bool, worked int) {
	f, err := os.OpenFile("daily.md", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	yyyy, mm, dd := time.Now().Date()
	output := fmt.Sprintf("## %d/%s/%d\n- Exercise: %t\n- Work: %d Minutes\n\n", dd, mm.String(), yyyy, exercised, worked)
	if _, err = f.WriteString(output); err != nil {
		panic(err)
	}
}
