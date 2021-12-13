package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for {
		var choice int
		fmt.Println("\n+========================+")
		fmt.Println("| CodeForces Problem Set |")
		fmt.Println("|     KOBEDELTA PHI      |")
		fmt.Println("+========================+")
		fmt.Println("|       MAIN MENU        |")
		fmt.Println("|[1] Word Capitalization |")
		fmt.Println("|[2] Translation         |")
		fmt.Println("|[3] Problem 71A         |")
		fmt.Println("|[4] Exit                |")
		fmt.Println("+========================+")
		fmt.Print("Option: ")
		fmt.Scanf("%d", &choice)
		fmt.Println("\n+========================+")
		switch choice {
		case 1:
			fmt.Println("|   Word Capitalization  |")
			fmt.Println("+========================+")
			wordCapitalization()
		case 2:
			fmt.Println("|       Translation      |")
			fmt.Println("+========================+")
			translation()
		case 3:
			fmt.Println("|       Problem 71A      |")
			fmt.Println("+========================+")
			seventyOneA()
		case 4:
			fmt.Println("            END           ")
			os.Exit(3)
		default:
			fmt.Println("|     INVALID INPUT!     |")
			fmt.Println("+========================+")

		}
	}

}

func wordCapitalization() {
	inputString := "i rEmEmbeR iT, aLl tOo wElL"
	tmp := []rune(inputString)
	outputString := string(unicode.ToUpper(tmp[0])) + string(strings.ToLower(inputString[1:]))
	fmt.Println(outputString)
}

func translation() {
	var s, t string
	fmt.Scan(&s, &t)
	for i := range s {
		if s[i] != t[len(t)-1-i] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}

func seventyOneA() {
	var a int
	fmt.Scanf("%d\n", &a)

	wrd := make([]string, a)
	for i := 0; i < a; i++ {
		fmt.Scanf("%s\n", &wrd[i])
	}
	for i := 0; i < a; i++ {
		if len(wrd[i]) > 10 {
			fmt.Printf("%c%d%c\n", wrd[i][0], (len(wrd[i]) - 2), wrd[i][len(wrd[i])-1])
		} else {
			fmt.Printf("%s\n", wrd[i])
		}
	}
}
