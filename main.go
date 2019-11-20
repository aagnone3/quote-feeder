package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func entry(c *cli.Context) error {
	fmt.Println("Grabbing the quote of the day")
	err := parseQuoteOfDay()
	err = queryQuoteOfDay()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "quote-feeder"
	app.Usage = "Feed me quotes on a recurring basis"
	app.Action = entry

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
