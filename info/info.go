package info

import "fmt"

func Success(msg string) {
	fmt.Println("[\033[32m\033[1m+\033[0m] " + msg)
}

func Info(msg string) {
	fmt.Println("[\033[34m\033[1m*\033[0m] " + msg)
}

func Error(msg string) {
	fmt.Println("[\033[31m\033[1mx\033[0m] " + msg)
}

func GreenMsg(msg string) {
	fmt.Println("\033[32m\033[1m" + msg + "\033[0m")
}

func BlueMsg(msg string) {
	fmt.Println("\033[34m\033[1m" + msg + "\033[0m")
}

func RedMsg(msg string) {
	fmt.Println("\033[31m\033[1m" + msg + "\033[0m")
}
