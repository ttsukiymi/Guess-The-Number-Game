package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

var secretnumber int
var hint int
var mNumber int
var levelName string

type Result struct {
	Date    string
	Outcome string
	Tries   int
	Level   string
}

func main() {
	for {

		lvl()

		rand.Seed(time.Now().UnixNano())
		secretnumber = rand.Intn(mNumber) + 1

		var gnumber int   // эта залупа вводится
		var guesses []int //попытки прошлые

		fmt.Printf("Игра 'Угадай число' - от 1 до %d началась!\n", mNumber)
		fmt.Printf("Нужно угадать число, заганное случайно за %d попыток. рескнешь?\n", hint)

		win := false
		triesUSed := 0

		for i := 1; i <= hint; i++ {
			fmt.Printf("Попытка № %d:", i)
			if len(guesses) > 0 {
				fmt.Println("")
				fmt.Println("введенные ранее числа: ")
				for _, v := range guesses {
					fmt.Print(v, " ")
				}
				fmt.Println()
			}

			_, err := fmt.Scan(&gnumber)
			if err != nil {
				fmt.Println(Red + "ошибка! ввести нужно число" + Reset)

				var buff string
				fmt.Scanln(&buff)

				i--
				continue
			}

			triesUSed = i
			guesses = append(guesses, gnumber)

			distance := int(math.Abs(float64(secretnumber - gnumber)))

			if gnumber == secretnumber {
				fmt.Println(Green + "поздравляю! ты угадал! правильно число: " + fmt.Sprint(secretnumber) + Reset)
				win = true
				break
			}

			if gnumber < secretnumber {
				fmt.Println("загаданное число больше")
			} else {
				fmt.Println("загаданное число меньше")
			}

			if distance <= 5 {
				fmt.Println(Yellow + "горячоо" + Reset)
			} else if distance <= 15 {
				fmt.Println(Yellow + "тепленько" + Reset)
			} else {
				fmt.Println(Yellow + "холодняк" + Reset)
			}
		}

		outcome := "проигрыш"
		if win {
			outcome = "победа"
		} else {
			fmt.Println("увы, но ты проиграл(" + Reset)
			fmt.Println("загаданное число было - ", secretnumber, Reset)

		}

		saveResult(outcome, triesUSed, levelName)

		var again string
		fmt.Printf("\n хочешь сыграть еще? да/нет: ")
		fmt.Scan(&again)

		if again != "да" {
			fmt.Println("спасибо за игру!")
			break
		}
	}
}

func lvl() {

	var level int

	fmt.Println("выбери уровень сложности")
	fmt.Println("1 - easy (1-50, 15 попыток)")
	fmt.Println("2 - normal (1-100, 10 попыток)")
	fmt.Println("3 - hard (1-200, 5 попыток)")
	fmt.Println("выбираешь:")
	fmt.Scan(&level)

	switch level {
	case 1:
		mNumber = 50
		hint = 15
		levelName = "easy"
	case 2:
		mNumber = 100
		hint = 10
		levelName = "normal"
	case 3:
		mNumber = 200
		hint = 5
		levelName = "hard"
	default:
		fmt.Println("неккоректно введен уровень, значит будет normal")
		mNumber = 100
		hint = 10
		levelName = "normal"

	}
}

func saveResult(outcome string, tries int, level string) {

	newResult := Result{

		Date:    time.Now().Format("2006-01-02 15:04"),
		Outcome: outcome,
		Tries:   tries,
		Level:   level,
	}

	var results []Result

	data, err := os.ReadFile("results.json")

	if err == nil {
		json.Unmarshal(data, &results)
	}

	results = append(results, newResult)

	file, _ := os.Create("results.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(results)
}
