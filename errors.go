package flip

import (
	"fmt"
	"strings"
)

type Error struct {
	NoCommitments  []Player
	NoReveals      []Player
	BadCommitments []Player
	Duplicates     []Player
}

func (e *Error) addNoCommitment(p Player) {
	e.NoCommitments = append(e.NoCommitments, p)
}

func (e *Error) addNoReveal(p Player) {
	e.NoReveals = append(e.NoReveals, p)
}

func (e *Error) addBadCommitment(p Player) {
	e.BadCommitments = append(e.BadCommitments, p)
}

func (e *Error) addDuplicate(p Player) {
	e.Duplicates = append(e.Duplicates, p)
}

func (e Error) IsNil() bool {
	return len(e.NoCommitments)+len(e.NoReveals)+len(e.BadCommitments) == 0
}

func (e Error) format(out []string, what string, players []Player) []string {
	if len(players) == 0 {
		return out
	}
	var playerStrings []string
	for _, p := range players {
		playerStrings = append(playerStrings, string(p))
	}
	s := fmt.Sprintf("Players %s: %s", what, strings.Join(playerStrings, ","))
	return append(out, s)
}

func (e Error) Error() string {
	var parts []string
	parts = e.format(parts, "without commitments", e.NoCommitments)
	parts = e.format(parts, "without reveals", e.NoReveals)
	parts = e.format(parts, "with bad commitments", e.BadCommitments)
	parts = e.format(parts, "with duplicated IDs", e.Duplicates)
	return fmt.Sprintf("Errors in flip: %s", strings.Join(parts, ";"))
}

type GameAlreadyStartedError GameKey

func (g GameAlreadyStartedError) Error() string {
	return fmt.Sprintf("Game already started: %s", g)
}

type GameFinishedError GameKey

func (g GameFinishedError) Error() string {
	return fmt.Sprintf("Game is finisehd: %s", g)
}

type TimeoutError struct {
	Key   GameKey
	Stage Stage
}

func (t TimeoutError) Error() string {
	return fmt.Sprintf("Game %s timed out in stage: %d", t.Stage)
}

type BadMessageForStageError struct {
	MessageStage Stage
	GameStage    Stage
}

func (b BadMessageForStageError) Error() string {
	return fmt.Sprintf("Message received (%s) was for wrong stage (%s)", b.MessageStage, b.GameStage)
}
