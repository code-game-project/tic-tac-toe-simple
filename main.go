package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/spf13/pflag"

	"github.com/code-game-project/tic-tac-toe-simple/tictactoesimple"
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
		DisplayName:             "Tic-Tac-Toe Simple",
		Version:                 "0.3",
		MaxPlayersPerGame:       2,
		Description:             "A simple implementation of tic-tac-toe for CodeGame. This game is ideal for familiarizing yourself with CodeGame.",
		RepositoryURL:           "https://github.com/code-game-project/tic-tac-toe-simple",
		Port:                    port,
		CGEFilepath:             "events.cge",
		DeleteInactiveGameDelay: 5 * time.Minute,
		KickInactivePlayerDelay: 15 * time.Minute,
	})

	server.Run(func(cgGame *cg.Game, config json.RawMessage) {
		var gameConfig tictactoesimple.GameConfig
		err := json.Unmarshal(config, &gameConfig)
		if err == nil {
			cgGame.SetConfig(gameConfig)
		} else {
			cgGame.Log.Error("Failed to unmarshal game config: %s", err)
		}

		tictactoesimple.NewGame(cgGame, gameConfig).Run()
	})
}
