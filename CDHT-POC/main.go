package main

import (
	_ "fmt"
	"poc/app/DNS"
)

func main(){
	app := DNS.TerminalUI{}
	app.Init()

	app.UserUI()
}