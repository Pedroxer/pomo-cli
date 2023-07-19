package main

import (
	"fmt"
	"log"
	"os"
	pomodoro "pomo-cli/Pomodoro"
	"time"

	"github.com/urfave/cli/v2"
)

func Info() {
	for true {
		time.Sleep(2 * time.Minute)
		h, m, _ := time.Now().Clock()
		 
		fmt.Printf("hour: %d minute: %d\n", h, m)
	}
}

func main() {
	var name string
	var dur string
	
	app := &cli.App{
		Name:        "Pomo-timer",
		Description: "make your work less stressful by adding pomodoro technique",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Value:       "none",
				Usage:       "Name of the subject",
				Destination: &name,
			},
			&cli.StringFlag{
				Name:        "dur",
				Value:       "25m",
				Usage:       "how long tomato is gonna be",
				Destination: &dur,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
				dur = cCtx.Args().Get(1)
			}
			tomat := pomodoro.Constructor(name, dur)

			fmt.Printf("Start the task. Point to %s\n", tomat.SubjectName)
			t := time.NewTimer(tomat.Timer)
			go Info()
			<-t.C

			fmt.Println("Your tomato is done")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
