package commands

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/smook1980/medialocker/app/server"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "HTTP server port",
		Value:  80,
		EnvVar: "PHOTOPRISM_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "http-host, i",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "PHOTOPRISM_HTTP_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		Value:  "",
		EnvVar: "PHOTOPRISM_HTTP_MODE",
	},
}

func startAction(ctx *cli.Context) error {
	// conf := context.NewConfig(ctx)

	// if conf.HttpServerPort() < 1 {
	//	log.Fatal("Server port must be a positive integer")
	// }

	// if err := conf.CreateDirectories(); err != nil {
	//	log.Fatal(err)
	// }

	// conf.MigrateDb()

	// fmt.Printf("Starting web server at %s:%d...\n", ctx.String("http-host"), ctx.Int("http-port"))

	// server.Start(conf)

	server.Start()
	fmt.Println("Done.")

	return nil
}
