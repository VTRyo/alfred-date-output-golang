package main

import (
	"fmt"
	"os"
	"time"
)

func timeFormat(time time.Time) {
	fmt.Println(time.Format("2006/01/02(Mon)"))
}

func today() {
	now := time.Now()
	timeFormat(now)
}
func yesterday() {
	yesterday := time.Now().AddDate(0, 0, -1)
	timeFormat(yesterday)
}
func tomorrow() {
	tomorrow := time.Now().AddDate(0, 0, 1)
	timeFormat(tomorrow)
}

func main() {
	args := os.Args[1]

	switch args {
	case "today":
		today()
	case "tomorrow":
		tomorrow()
	case "yesterday":
		yesterday()
	}
}
