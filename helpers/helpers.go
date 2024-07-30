package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error al limpiar la consola: %v\n", err)
	}
}

func RemoveAccents(s string) (string, bool) {
	accents := "áéíóúÁÉÍÓÚñÑ"
	for _, r := range s {
		if strings.ContainsRune(accents, r) {
			return s, true
		}
	}
	return s, false
}

func ShowHangman(steps, difficult int) {
	var count int

	if difficult == 1 {
		count = steps - 4
	} else {
		count = steps
	}

	if count == 6 {
		fmt.Println("_________")
		fmt.Println("|/    	 |")
		fmt.Print("|       ")
		fmt.Println("(x)")
		fmt.Print("|	")
		fmt.Println("/|\\")
		fmt.Print("|       ")
		fmt.Println("/ \\")
		fmt.Println("|__________")
	}

	switch count {
	case 5:
		fmt.Println("_________")
		fmt.Println("|/    	 |")
		fmt.Print("|       ")
		fmt.Println("(x)")
		fmt.Print("|	")
		fmt.Println("/|\\")
		fmt.Println("|       ")
		fmt.Println("|__________")
	case 4:
		fmt.Println("_________")
		fmt.Println("|/    	 |")
		fmt.Print("|       ")
		fmt.Println("(x)")
		fmt.Println("|	")
		fmt.Println("|       ")
		fmt.Println("|__________")
	case 3:
		fmt.Println("_________")
		fmt.Println("|/    	 |")
		fmt.Println("|       ")
		fmt.Println("|	")
		fmt.Println("|       ")
		fmt.Println("|__________")
	case 2:
		fmt.Println()
		fmt.Println()
		fmt.Println("|       ")
		fmt.Println("|	")
		fmt.Println("|       ")
		fmt.Println("|__________")
	case 1:
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("__________")
	case 0:
		if difficult == 1 {

			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("  ________")
		} else {

			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	case -1:
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("    ______")
	case -2:
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("      ____")
	case -3:
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("        __")
	case -4:
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func ShowSavedman() {
	fmt.Println("_________")
	fmt.Println("|/    	 |")
	fmt.Println("|  ")
	fmt.Print("|     ")
	fmt.Println("\\( )/")
	fmt.Print("|       ")
	fmt.Println("|")
	fmt.Println("|______/_\\__")
}
