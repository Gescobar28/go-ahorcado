package players

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gp-practice/helpers"
	"github.com/gp-practice/words"
)

type Player struct {
	Username    string
	WinnerWords []words.Word
	LoserWords  []words.Word
}

func createPlayer(id int, reader *bufio.Reader) Player {
	var p Player

	fmt.Printf("Ingrese el usuario del jugador %d\n", id)
	input, _ := reader.ReadString('\n')

	fmt.Println("*************************")
	p.Username = strings.TrimSpace(input)
	return p
}

func AddPlayers(currentPlayers []Player) []Player {
	reader := bufio.NewReader(os.Stdin)
	continueAdd := "S"

	for {
		if strings.ToUpper(continueAdd) == "N" {
			break
		}
		helpers.ClearConsole()
		currentPlayers = append(currentPlayers, createPlayer(len(currentPlayers)+1, reader))
		fmt.Println("Desea agregar otro jugador? (S/N) ")
		fmt.Scanf("%s\n", &continueAdd)
	}

	return currentPlayers
}
