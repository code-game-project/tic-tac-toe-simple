package tictactoesimple

import (
	"fmt"

	"github.com/code-game-project/go-server/cg"
)

type Game struct {
	cg     *cg.Game
	config GameConfig

	crossPlayer  *cg.Player
	circlePlayer *cg.Player
	currentTurn  Sign

	board [][]Field
}

func NewGame(cgGame *cg.Game, config GameConfig) *Game {
	game := &Game{
		cg:          cgGame,
		config:      config,
		currentTurn: SignO,
	}

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
		cmd, ok := g.cg.WaitForNextCommand()
		if !ok {
			break
		}
		g.handleCommand(cmd.Origin, cmd.Cmd)
	}
}

func (g *Game) onPlayerLeft(player *cg.Player) {
	g.cg.Close()
}

func (g *Game) onPlayerSocketConnected(player *cg.Player, socket *cg.GameSocket) {
	if g.crossPlayer == nil {
		g.crossPlayer = player
		return
	} else if g.circlePlayer == nil {
		if g.crossPlayer != player {
			g.circlePlayer = player
			g.start()
		}
		return
	}

	socket.Send(StartEvent, StartEventData{
		Signs: map[string]Sign{
			g.crossPlayer.Id:  SignX,
			g.circlePlayer.Id: SignO,
		},
	})

	socket.Send(BoardEvent, BoardEventData{
		Board: g.board,
	})

	socket.Send(TurnEvent, TurnEventData{
		Sign: g.currentTurn,
	})
}

func (g *Game) handleCommand(origin *cg.Player, cmd cg.Command) {
	switch cmd.Name {
	case MarkCmd:
		g.mark(origin, cmd)
	default:
		origin.Log.ErrorData(cmd, fmt.Sprintf("unexpected command: %s", cmd.Name))
	}
}

func (g *Game) mark(player *cg.Player, cmd cg.Command) {
	if (g.currentTurn == SignX && player != g.crossPlayer) || (g.currentTurn == SignO && player != g.circlePlayer) {
		player.Send(InvalidActionEvent, InvalidActionEventData{
			Message: "It is not your turn.",
		})
		return
	}

	var data MarkCmdData
	err := cmd.UnmarshalData(&data)
	if err != nil {
		player.Log.ErrorData(cmd, "invalid event data")
		return
	}
	if data.Row < 0 || data.Row > 2 || data.Column < 0 || data.Column > 2 {
		player.Send(InvalidActionEvent, InvalidActionEventData{
			Message: "Invalid coordinates.",
		})
		return
	}

	if g.board[data.Row][data.Column].Sign != SignNone {
		player.Send(InvalidActionEvent, InvalidActionEventData{
			Message: "The field is alread occupied.",
		})
		return
	}

	sign := SignX
	if player == g.circlePlayer {
		sign = SignO
	}

	g.board[data.Row][data.Column].Sign = sign

	g.sendBoard()
	if !g.checkDone() {
		g.turn()
	}
}

func (g *Game) start() {
	g.cg.Send(StartEvent, StartEventData{
		Signs: map[string]Sign{
			g.crossPlayer.Id:  SignX,
			g.circlePlayer.Id: SignO,
		},
	})
	g.sendBoard()
	g.turn()
}

func (g *Game) turn() {
	if g.currentTurn == SignX {
		g.currentTurn = SignO
	} else {
		g.currentTurn = SignX
	}
	g.cg.Send(TurnEvent, TurnEventData{
		Sign: g.currentTurn,
	})
}

func (g *Game) sendBoard() {
	g.cg.Send(BoardEvent, BoardEventData{
		Board: g.board,
	})
}

func (g *Game) checkDone() bool {
	for i := 0; i < 3; i++ {
		if g.board[i][0].Sign != SignNone && g.board[i][0].Sign == g.board[i][1].Sign && g.board[i][1].Sign == g.board[i][2].Sign {
			g.gameOver(false, []Field{g.board[i][0], g.board[i][1], g.board[i][2]})
			return true
		}

		if g.board[0][i].Sign != SignNone && g.board[0][i].Sign == g.board[1][i].Sign && g.board[1][i].Sign == g.board[2][i].Sign {
			g.gameOver(false, []Field{g.board[0][i], g.board[1][i], g.board[2][i]})
			return true
		}
	}

	if g.board[0][0].Sign != SignNone && g.board[0][0].Sign == g.board[1][1].Sign && g.board[1][1].Sign == g.board[2][2].Sign {
		g.gameOver(false, []Field{g.board[0][0], g.board[1][1], g.board[2][2]})
		return true
	}

	if g.board[0][2].Sign != SignNone && g.board[0][2].Sign == g.board[1][1].Sign && g.board[1][1].Sign == g.board[2][0].Sign {
		g.gameOver(false, []Field{g.board[0][0], g.board[1][1], g.board[2][2]})
		return true
	}

	for row := range g.board {
		for column := range g.board[row] {
			if g.board[row][column].Sign == SignNone {
				return false
			}
		}
	}

	g.gameOver(true, nil)
	return true
}

func (g *Game) gameOver(tie bool, fields []Field) {
	if tie {
		g.cg.Send(GameOverEvent, GameOverEventData{
			Tie: true,
		})
	} else {
		g.cg.Send(GameOverEvent, GameOverEventData{
			WinnerSign: fields[0].Sign,
			WinningRow: fields,
		})
	}

	g.cg.Close()
}
