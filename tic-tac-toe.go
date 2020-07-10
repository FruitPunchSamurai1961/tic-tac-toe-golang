package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func drawBoard(gameBoard *[3][3]string) {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			if j == 2 {
				fmt.Print(gameBoard[i][j])
				fmt.Println()
			} else {
				fmt.Print(gameBoard[i][j])
				fmt.Print("|")
			}
		}
		if i != 2 {
			fmt.Print("-----\n")
		}
	}
}

type Player struct {
	name  string
	value string
}

func gameLoop(player Player, gameBoard *[3][3]string, reader *bufio.Reader) error {
	fmt.Printf("Enter where you would like to place your %v, %v:", player.value, player.name)
	ans, _ := reader.ReadString('\n')
	row, _ := strconv.Atoi(ans[:len(ans)-3])
	col, _ := strconv.Atoi(ans[len(ans)-2 : len(ans)-1])
	if row > 2 || col > 2 {
		return errors.New("out")
	}
	if gameBoard[row][col] != "." {
		return errors.New("taken")
	}
	gameBoard[row][col] = player.value
	drawBoard(gameBoard)
	if checkWin(gameBoard, player.value) {
		return errors.New("win")
	}
	return nil
}

func checkWin(gameBoard *[3][3]string, value string) bool {
	//horizontal and vertical
	for i := range gameBoard {
		if gameBoard[i][0] == value && gameBoard[i][1] == value && gameBoard[i][2] == value {
			return true
		}
		if gameBoard[0][i] == value && gameBoard[1][i] == value && gameBoard[2][i] == value {
			return true
		}
	}
	//diagonal
	if gameBoard[0][0] == value && gameBoard[1][1] == value && gameBoard[2][2] == value {
		return true
	}
	if gameBoard[0][2] == value && gameBoard[1][1] == value && gameBoard[2][0] == value {
		return true
	}
	return false
}
func isEmpty(gameBoard *[3][3]string) bool {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			if gameBoard[i][j] == "." {
				return true
			}
		}
	}
	return false
}

func main() {
gameLoop:
	for {
		gameBoard := [3][3]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		}
		fmt.Print("Welcome to Golang Tic-Tac-Toe!\nSome small note:\n1.Enter your input in the [row],[column] format using numbers from 0-2\nExample: 1,2\n")
		drawBoard(&gameBoard)

		reader := bufio.NewReader(os.Stdin)
		player1 := Player{
			name:  "Player 1",
			value: "X"}

		player2 := Player{
			name:  "Player 2",
			value: "O",
		}

		turn := "Player1"
		for isEmpty(&gameBoard) {
			if turn == "Player1" {
				err := gameLoop(player1, &gameBoard, reader)
				if err != nil {
					if err.Error() == "out" {
						fmt.Print("1.Enter your input in the [row],[column] format using numbers from 0-2\nExample: 1,2\n")
						drawBoard(&gameBoard)
						continue
					} else if err.Error() == "taken" {
						fmt.Print("Please choose a square that is not taken\n")
						drawBoard(&gameBoard)
						continue
					} else if err.Error() == "win" {
						fmt.Printf("%v has won!", player1.name)
						break
					}
				}
				turn = "Player2"
			} else {
				err := gameLoop(player2, &gameBoard, reader)
				if err != nil {
					if err.Error() == "out" {
						fmt.Print("1.Enter your input in the [row],[column] format using numbers from 0-2\nExample: 1,2\n")
						drawBoard(&gameBoard)
						continue
					} else if err.Error() == "taken" {
						fmt.Print("Please choose a square that is not taken\n")
						drawBoard(&gameBoard)
						continue
					} else if err.Error() == "win" {
						fmt.Printf("%v has won!", player2.name)
						break
					}
				}
				turn = "Player1"
			}
		}
		if !isEmpty(&gameBoard) && !checkWin(&gameBoard, player2.value) && !checkWin(&gameBoard, player1.value) {
			fmt.Println("The game was a draw!")
		}
		fmt.Println("Would you like to play again?(y/n)")
		answer,_ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer != "y" {
			break gameLoop
		}
	}
}
