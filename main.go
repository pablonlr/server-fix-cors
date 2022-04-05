package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	var path string
	var port int
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Value:       "./static",
				Usage:       "Path to WebsiteFolder",
				Destination: &path,
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       8080,
				Usage:       "Port in localhost to run the server",
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			fileServer := http.FileServer(http.Dir(path))
			http.Handle("/", fileServer)
			sport := ":" + strconv.Itoa(port)
			log.Printf("Starting server at port %s", sport)
			if err := http.ListenAndServe(sport, nil); err != nil {
				return err
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
