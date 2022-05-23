package main

import (
	"os"
	"strconv"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/ogier/pflag"
)

func main() {
	var port int
	pflag.IntVarP(&port, "port", "p", 0, "The network port of the game server.")
	pflag.Parse()

	if port == 0 {
		portStr, ok := os.LookupEnv("CG_PORT")
		if ok {
			port, _ = strconv.Atoi(portStr)
		}
	}

	if port == 0 {
		port = 80
	}

	server := cg.NewServer("tic-tac-toe-simple", cg.ServerConfig{
		Port:                    port,
		CGEFilepath:             "events.cge",
		MaxPlayersPerGame:       2,
		DeleteInactiveGameDelay: 15 * time.Minute,
		KickInactivePlayerDelay: 10 * time.Minute,
		DisplayName:             "Tic-Tac-Toe Simple",
		Version:                 "0.1",
		Description:             "A simple implementation of tic-tac-toe for CodeGame. This game is ideal for familiarizing yourself with CodeGame.",
		RepositoryURL:           "https://github.com/code-game-project/tic-tac-toe-simple",
	})

	server.Run(func(game *cg.Game) {
	})
}
