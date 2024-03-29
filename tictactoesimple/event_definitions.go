package tictactoesimple

import "github.com/code-game-project/go-server/cg"

// The 'mark' command can be sent to the server to mark a field with the player's sign.
const MarkCmd cg.CommandName = "mark"

type MarkCmdData struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

// The 'start' event is sent to all players when the game starts.
const StartEvent cg.EventName = "start"

type StartEventData struct {
	// A map of player IDs mapped to their respective signs.
	Signs map[string]Sign `json:"signs"`
}

// The 'board' event is sent to all players when the board changes.
const BoardEvent cg.EventName = "board"

type BoardEventData struct {
	// The board as rows of columns.
	Board [][]Field `json:"board"`
}

// The 'invalid_action' event notifies the player that their action was not allowed.
const InvalidActionEvent cg.EventName = "invalid_action"

type InvalidActionEventData struct {
	// A message containing details on what the player did wrong.
	Message string `json:"message"`
}

// The 'turn' event is sent to all players when it is the next player's turn.
const TurnEvent cg.EventName = "turn"

type TurnEventData struct {
	// The sign of the player whose turn it is now.
	Sign Sign `json:"sign"`
}

// The 'game_over' event is sent to all players when it's a tie or a player has won.
const GameOverEvent cg.EventName = "game_over"

type GameOverEventData struct {
	// Whether it's a tie.
	Tie bool `json:"tie"`
	// The sign of the winner.
	WinnerSign Sign `json:"winner_sign"`
	// The three fields which form a row.
	WinningRow []Field `json:"winning_row"`
}

// The two possible signs which can be placed on the board.
type Sign string

const (
	SignNone Sign = "none"
	SignX    Sign = "x"
	SignO    Sign = "o"
)

// One of the nine fields on the board.
type Field struct {
	Row    int  `json:"row"`
	Column int  `json:"column"`
	Sign   Sign `json:"sign"`
}

type GameConfig struct {
}
