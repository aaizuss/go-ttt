package console

import (
	"strconv"

	"github.com/aaizuss/go-ttt/board"
)

func (cli *Console) ShowBoard(board board.Board) {
	cli.Show(wrapNewLine(board.FormattedString()))
}

func (cli *Console) ShowOutcome(board board.Board) {
	if board.IsTie() {
		cli.Show(tie)
	} else {
		winner, _ := board.Winner()
		winMessage := "Player " + winner + " wins!\n"
		cli.Show(winMessage)
	}
}

func (cli *Console) GetMove(board board.Board) (move int) {
	for {
		cli.Show(chooseMovePrompt)
		move := cli.Read()

		if IsValidMoveChoice(board, move) {
			return toInt(move)
		} else {
			cli.Show(invalidMove)
		}
	}
}

func (cli *Console) GetGameChoice() string {
	cli.Show(welcome)
	for {
		cli.Show(gameMenu)
		choice := cli.Read()

		if IsValidGameChoice(choice) {
			return choice
		} else {
			cli.Show(invalidChoice)
		}
	}
}

func (cli *Console) ShowMoveRecap(marker string, move int) {
	cli.Show(marker + " marked " + (strconv.Itoa(move)) + "\n")
}

const (
	welcome  = "|----------------------------|\n|-- Welcome to Tic Tac Toe --|\n|----------------------------|\n\n"
	tie      = "It's a tie!\n"
	gameMenu = "Enter a number to select a game type.\n\n" +
		"    P1 (x)        P2 (o)\n" +
		"----------------------------\n" +
		"1  human     v   human\n" +
		"2  human     v   computer\n" +
		"3  computer  v   human\n\n"
	chooseMovePrompt = "Enter a number 0-8 to mark that position on the board: "
	invalidMove      = "You can't move there. "
	invalidChoice    = "That's not an option.\n"
)

func wrapNewLine(message string) string {
	return "\n" + message + "\n"
}

func toInt(inputMove string) int {
	move, _ := strconv.Atoi(inputMove)
	return move
}
