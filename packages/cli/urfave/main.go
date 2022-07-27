package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")

			loglevel := c.String("loglevel")

			fmt.Println(111111, loglevel)

			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "loglevel",
				Value: 4,
				Usage: "log level to emit to the screen",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}


}
