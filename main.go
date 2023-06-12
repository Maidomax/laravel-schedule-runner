package main

import (
	"fmt"
	"os/exec"

	"github.com/robfig/cron"
)

func main() {
	fmt.Println("Started Laravel Schedule Runner")

	myCron := cron.New()
	myCron.AddFunc("0 * * * * *", execSchedule)
	myCron.Start()

	select {}
}

func execSchedule() {
	fmt.Println("Running schedule")

	cmd := exec.Command("php", "artisan", "schedule:run")

	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Could not run command: ", err)
	} else {
		fmt.Println(string(out))
	}
}
