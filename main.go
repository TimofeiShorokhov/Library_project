package main

import (
	"Library_project/email"
	"Library_project/other"
	"Library_project/routers"
)

func main() {
	other.ConnectDB()
	email.CheckForSend()
	go email.TickerForEmail()
	routers.Routers()
}
