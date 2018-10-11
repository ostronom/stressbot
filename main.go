package main

import (
	"github.com/dialogs/stressbot/stress"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var serverUrl, certsPath, usersFile, sendFilePath string
	var serverPort, botsNum, cPar, grQty, grMin, grMax, sFreq, rFreq, filesFreq int64

	//sendFilePath := flag.String("filepath", "", "file path for sending")
	app := cli.NewApp()
	app.Name = "streesbot"
	app.Commands = []cli.Command{cli.Command{
		Name: "start",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "server-url", Destination: &serverUrl},
			cli.Int64Flag{Name: "server-port", Destination: &serverPort},
			cli.Int64Flag{Name: "bots-num", Value: 100, Destination: &botsNum},
			cli.Int64Flag{Name: "creation-parallelism", Value: 5, Destination: &cPar},
			cli.Int64Flag{Name: "groups-amount", Value: 10, Destination: &grQty},
			cli.Int64Flag{Name: "groups-members-min", Value: 1, Destination: &grMin},
			cli.Int64Flag{Name: "groups-members-max", Value: 5, Destination: &grMax},
			cli.Int64Flag{Name: "send-frequency", Usage: "Mean interval in seconds between sendMessage requests", Value: 5, Destination: &sFreq},
			cli.Int64Flag{Name: "reads-frequency", Usage: "Mean interval in seconds between readMessage requests", Value: 5, Destination: &rFreq},
			cli.StringFlag{Name: "certs-path", Destination: &certsPath},
			cli.StringFlag{Name: "users-file", Destination: &usersFile},

			cli.StringFlag{Name: "files-path", Destination: &sendFilePath},
			cli.Int64Flag{Name: "file-send-frequency", Usage: "Mean interval in seconds between sending files requests", Value: 5, Destination: &filesFreq},
		},
		Action: func(c *cli.Context) error {
			stress.Stress(serverUrl, certsPath, usersFile, sendFilePath, serverPort, cPar, botsNum, grQty, grMin, grMax, sFreq, rFreq, filesFreq)
			return nil
		},
	}}
	app.Run(os.Args)
}
