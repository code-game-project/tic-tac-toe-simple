package tictactoe

import (
	"fmt"

	"github.com/code-game-project/go-server/cg"
)

type Game struct {
	cg *cg.Game

	crossPlayer  *cg.Player
	circlePlayer *cg.Player
	currentTurn  Sign

	board [][]Field
}

func NewGame(cgGame *cg.Game) *Game {
	game := &Game{
		cg:          cgGame,
		currentTurn: SignO,
	}

	cgGame.OnPlayerJoined = game.onPlayerJoined
	cgGame.OnPlayerLeft = game.onPlayerLeft
	cgGame.OnPlayerSocketConnected = game.onPlayerSocketConnected

	game.board = make([][]Field, 3)
	for i := range game.board {
		game.board[i] = make([]Field, 3)
		for j := range game.board[i] {
			game.board[i][j] = Field{
				Row:    i,
				Column: j,
				Sign:   SignNone,
			}
		}
	}

	return game
}

func (g *Game) Run() {
	for g.cg.Running() {
		event, ok := g.cg.WaitForNextEvent()
		if !ok {
			break
		}
		g.handleEvent(event.Player, event.Event)
	}
}

func (g *Game) onPlayerJoined(player *cg.Player) {
	if g.crossPlayer == nil {
		g.crossPlayer = player
	} else {
		g.circlePlayer = player
		g.start()
	}
}

func (g *Game) onPlayerLeft(player *cg.Player) {
	g.cg.Close()
}

func (g *Game) onPlayerSocketConnected(player *cg.Player, socket *cg.Socket) {
	socket.Send("server", EventStart, EventStartData{
		Signs: map[string]Sign{
			g.crossPlayer.Id:  SignX,
			g.circlePlayer.Id: SignO,
		},
	})

	socket.Send("server", EventBoard, EventBoardData{
		Board: g.board,
	})

	socket.Send("server", EventTurn, EventTurnData{
		Sign: g.currentTurn,
	})
}

func (g *Game) handleEvent(player *cg.Player, event cg.Event) {
	switch event.Name {
	case EventMark:
		g.mark(player, event)
	default:
		player.Send(player.Id, cg.EventError, cg.EventErrorData{
			Message: fmt.Sprintf("unexpected event: %s", event.Name),
		})
	}
}

func (g *Game) mark(player *cg.Player, event cg.Event) {
	if (g.currentTurn == SignX && player != g.crossPlayer) || (g.currentTurn == SignO && player != g.circlePlayer) {
		player.Send("server", EventInvalidAction, EventInvalidActionData{
			Message: "It is not your turn.",
		})
		return
	}

	var data EventMarkData
	err := event.UnmarshalData(&data)
	if err != nil {
		player.Send("server", cg.EventError, cg.EventErrorData{
			Message: "invalid event data",
		})
		return
	}
	if data.Row < 0 || data.Row > 2 || data.Column < 0 || data.Column > 2 {
		player.Send("server", EventInvalidAction, EventInvalidActionData{
			Message: "Invalid coordinates.",
		})
		return
	}

	if g.board[data.Row][data.Column].Sign != SignNone {
		player.Send("server", EventInvalidAction, EventInvalidActionData{
			Message: "The field is alread occupied.",
		})
		return
	}

	sign := SignX
	if player == g.circlePlayer {
		sign = SignO
	}

	g.board[data.Row][data.Column].Sign = sign
	g.turn()
}

func (g *Game) start() {
	g.cg.Send("server", EventStart, EventStartData{
		Signs: map[string]Sign{
			g.crossPlayer.Id:  SignX,
			g.circlePlayer.Id: SignO,
		},
	})
	g.turn()
}

func (g *Game) turn() {
	g.sendBoard()
	if g.currentTurn == SignX {
		g.currentTurn = SignO
	} else {
		g.currentTurn = SignX
	}
	g.cg.Send("server", EventTurn, EventTurnData{
		Sign: g.currentTurn,
	})
}

func (g *Game) sendBoard() {
	g.cg.Send("server", EventBoard, EventBoardData{
		Board: g.board,
	})
}
