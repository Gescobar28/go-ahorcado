package words

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gp-practice/helpers"
)

type Word struct {
	Word   string
	Length int
	Points int
}

func LoadWord(w, letter, gw string) string {

	currentLetter := ""

	for i, l := range w {
		if string(l) == " " {
			fmt.Print(" ")
			currentLetter += string(l)
		} else if len(gw) > 0 && string(gw[i]) != " " && string(gw[i]) != "_" {
			currentLetter += string(gw[i])
		} else if strings.EqualFold(string(l), string(letter)) {
			fmt.Print(strings.ToUpper(letter))
			currentLetter += strings.ToUpper(string(l))
		} else {
			fmt.Print("_")
			currentLetter += string("_")
		}
	}

	return currentLetter

}

func getWord() (Word, bool) {
	var w Word

	url := `https://clientes.api.greenborn.com.ar/public-random-word`
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error leyendo la respuesta: %v\n", err)
	}

	defer response.Body.Close()

	var word []string
	err = json.Unmarshal(body, &word)
	if err != nil {
		fmt.Printf("Error deserializando la respuesta: %v\n", err)

	}

	apiWord, exists := helpers.RemoveAccents(word[0])

	w.Word = apiWord
	w.Length = len(apiWord)
	w.Points = len(apiWord) * 10

	return w, exists
}

func GuessWord(difficult int) string {
	var apiWord Word
	var secretWord string
	var guessWord string

	var showSecret int
	var steps int
	var lettersEntered string
	trys := 6
	level := "Difícil"

	if difficult == 1 {
		trys = 10
	}

	if difficult == 1 {
		level = "Fácil"
	} else if difficult == 2 {
		level = "Intermedio"
	} else {
		level = "Difícil"
	}

	for {
		word, exists := getWord()

		if !exists {
			apiWord = word
			break
		}

	}

	for _, l := range apiWord.Word {
		if string(l) == " " {
			secretWord += " "
		} else {
			secretWord += "_"

		}
	}

	fmt.Println()

	for {
		reader := bufio.NewReader(os.Stdin)
		var letter string
		helpers.ClearConsole()

		fmt.Printf("Nivel: %s.\n", level)
		fmt.Printf("La palabra contiene %d letras.\n", apiWord.Length)

		helpers.ShowHangman(steps, difficult)
		fmt.Println()

		if showSecret >= 1 {
			fmt.Print("--->     ")
			fmt.Println(guessWord)
		} else {
			fmt.Print("--->     ")
			fmt.Println(secretWord)
		}

		fmt.Println()
		if difficult == 1 || difficult == 2 {
			fmt.Printf("Letras ingresadas: %s\n", strings.ToUpper(lettersEntered))
			fmt.Println()
		}
		fmt.Printf("Numero de intentos restantes: %d\n", trys)
		fmt.Println()

		fmt.Println("Ingresa una letra o la palabra oculta")

		input, _ := reader.ReadString('\n')
		letter = strings.TrimSpace(input)

		fmt.Println()

		if len(letter) > 1 && strings.EqualFold(strings.ToUpper(letter), strings.ToUpper(apiWord.Word)) {
			return winnerRound(apiWord, reader)
		}

		guessWord = LoadWord(apiWord.Word, letter, guessWord)
		showSecret++
		if !strings.Contains(apiWord.Word, letter) {
			trys--
			steps++
			if len(letter) == 1 && letter != " " && !strings.Contains(lettersEntered, letter) {
				lettersEntered += letter + " "
			}
		}

		helpers.ClearConsole()

		//Si se terminan los intentos finalizamos la ronda y mostramos un mensaje
		if trys == 0 {
			return loserRound(apiWord, reader, steps, difficult)
		}

		//Si adivina la palabra finalizamos la ronda y mostramos un mensaje
		if len(guessWord) > 1 && !strings.Contains(guessWord, "_") {
			return winnerRound(apiWord, reader)
		}

	}

}

func winnerRound(w Word, r *bufio.Reader) string {
	helpers.ClearConsole()
	fmt.Println("Felicidades!! Has Ganado!!")
	fmt.Println()

	helpers.ShowSavedman()
	fmt.Println()

	fmt.Printf("La palabra oculta era '%s'!!\n", strings.ToUpper(w.Word))
	fmt.Println()

	option := ""

	fmt.Println("Deseas iniciar otra ronda? S/N")
	fmt.Scanf("%s", &option)
	r.ReadString('\n')

	return option
}

func loserRound(w Word, r *bufio.Reader, steps, difficult int) string {
	helpers.ClearConsole()
	fmt.Println("No lo has conseguido esta vez, se terminaron los intentos.")
	helpers.ShowHangman(steps, difficult)

	fmt.Println()
	fmt.Printf("La palabra oculta era '%s' \n", strings.ToUpper(w.Word))
	fmt.Println()

	option := ""

	fmt.Println("Te gustaría intentarlo de nuevo? S/N")
	fmt.Scanf("%s", &option)
	r.ReadString('\n')

	return option
}
