package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gp-practice/helpers"
	"github.com/gp-practice/players"
	"github.com/gp-practice/words"
)

func configGame(difficultGame *int, reader *bufio.Reader) {
	helpers.ClearConsole()

	fmt.Println("Configuración del Juego")
	fmt.Println()

	fmt.Println("Intentos: Es la cantidad permitida de letras incorrectas, es decir, que no se encuentren en la palabra.")

	fmt.Println("Mostrar Letras: El jugador puede ver las letras que ya ingresó para evitar repetirlas.")
	fmt.Println()

	fmt.Println("Seleccione una opción:")
	fmt.Println()

	fmt.Println("1. Fácil - Diez intentos, se muestran las letras ingresadas que son incorrectas")
	fmt.Println("2. Intermedio - Seis intentos, se muestran las letras ingresadas que son incorrectas")
	fmt.Println("3. Difícil - Seis intentos, NO se muestran las letras ingresadas que son incorrectas")
	fmt.Println()

	difficult, _ := reader.ReadString('\n')
	optionDifficult, err := strconv.Atoi(strings.TrimSpace(difficult))
	if err != nil {
		fmt.Println(err)
	}

	*difficultGame = optionDifficult
}

func main() {
	var currentPlayers []players.Player
	var difficult int

	reader := bufio.NewReader(os.Stdin)

	helpers.ClearConsole()

	for {

		keepPlaying := true
		helpers.ClearConsole()

		fmt.Println(`¡Bienvenido al juego: El Ahorcado! 
		
En este juego, tu objetivo es adivinar una palabra oculta ingresando letras. 
Cada letra que adivines se revelará en la palabra. 

Tienes un número limitado de intentos para adivinar la palabra completa antes 
de que el juego termine.`)

		fmt.Println()
		fmt.Println("Selecciona una opción ingresando el número correspondiente:")
		fmt.Println()

		fmt.Println("1. Jugar")
		fmt.Println("2. Configurar Juego")
		fmt.Println("3. Ayuda")
		fmt.Println("4. Salir")
		fmt.Println()

		input, _ := reader.ReadString('\n')
		option, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Error al leer la opción. Por favor ingrese un número válido.")
			continue
		}

		switch option {
		case 6:
			currentPlayers = players.AddPlayers(currentPlayers)
			helpers.ClearConsole()

			fmt.Println("")
			fmt.Println("Jugadores agregados")
			fmt.Println("")

			for k, v := range currentPlayers {
				fmt.Printf("%d - %s \n", k+1, v.Username)
			}

			fmt.Println("___________________________________________")
			fmt.Println("")

		case 1:
			helpers.ClearConsole()

			fmt.Println("Cargando...")

			for {

				option := words.GuessWord(difficult)

				if option == "n" {
					break
				}
			}

		case 2:
			configGame(&difficult, reader)
		case 3:

			helpers.ClearConsole()
			fmt.Println(`Objetivo del Juego: Adivina la palabra oculta antes de que se te acaben los intentos permitidos.

Cómo Jugar:

- La palabra oculta se representará con guiones bajos (_) para las letras no adivinadas.
- Ingresarás una letra a la vez.
- Si la letra está en la palabra oculta, se mostrará en su posición correspondiente.
- Si la letra no está en la palabra, perderás un intento.

Intentos:

- Tienes un número limitado de intentos para adivinar la palabra.
- El juego termina si adivinas la palabra correctamente o si te quedas sin intentos.

Instrucciones:

- Ingresa una letra y presiona Enter.
- Puedes intentar adivinar cualquier letra en cualquier momento del juego.
- Asegúrate de ingresar solo una letra a la vez.

Ejemplo:

Palabra oculta: golang

Supongamos que ingresas la letra g, la palabra mostrada podría ser g__a_g.

Continúa ingresando letras hasta que adivines la palabra o se te acaben los intentos.

¡Buena suerte y diviértete!

Presiona enter para volver al Menú Principal`)

			reader.ReadLine()

		case 4:
			keepPlaying = false
		default:
			fmt.Println("Opción no válida.")
		}

		if !keepPlaying {
			break
		}

	}

	helpers.ClearConsole()

	fmt.Println("Gracias por jugar!!")
	fmt.Println("")
	fmt.Println("Presiona enter para cerrar la ventana")
	fmt.Scanln()

}
