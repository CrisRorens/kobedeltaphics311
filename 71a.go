package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//sample push

func main() {
	for {
		var choice int
		fmt.Println("\n+========================+")
		fmt.Println("| CODEFORCES PROBLEM SET |")
		fmt.Println("|     KOBEDELTA PHI      |")
		fmt.Println("+========================+")
		fmt.Println("|       MAIN MENU        |")
		fmt.Println("|[1] Word Capitalization |")
		fmt.Println("|[2] Translation         |")
		fmt.Println("|[3] Problem 71A         |")
		fmt.Println("|[4] Strange Table       |")
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
			fmt.Println("|      Strange Table     |")
			fmt.Println("+========================+")
			strangeTable()
		case 5:
			fmt.Println("|    Telephone Number    |")
			fmt.Println("+========================+")
			telephoneNumber()
		case 6:
			fmt.Println("|       Buy Shovel       |")
			fmt.Println("+========================+")
			buyShovel()
		case 7:
			fmt.Println("            END           ")
			os.Exit(3)
		default:
			fmt.Println("|     INVALID INPUT!     |")
			fmt.Println("+========================+")

		}
	}
}

func telephoneNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	str := scanner.Text()
	j := int64(strings.Index(str, "8"))
	t := n - j
	if n < 11 {
		fmt.Println("NO")
	} else {
		if t < 11 || j < 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

func strangeTable() {
	var t int
	fmt.Print()
	fmt.Scan(&t)

	var i int
	for i = 0; i < t; i++ {
		var a, b, c, x, y, z int
		fmt.Print()
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)

		x = c % a
		if x == 0 {
			x = a
		}
		y = (c + a - 1) / a
		z = (x-1)*b + y

		fmt.Println(z)
	}
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

func buyShovel() {
	var k, r, i, count, value int
	fmt.Scan(&k)
	fmt.Scan(&r)

	i = 2
	count = 1
	value = k

	for true {
		if value%10 == 0 {
			fmt.Println(count)
			break
		}
		if value%10 == r {
			fmt.Println(count)
			break
		}
		value = k * i
		count++
		i++
	}
}

func wordCapitalization() {
	inputString := "i rEmEmbeR iT, aLl tOo wElL"
	tmp := []rune(inputString)
	outputString := string(unicode.ToUpper(tmp[0])) + string(strings.ToLower(inputString[1:]))
	fmt.Println(outputString)
}
