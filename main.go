package main

import (
	"Library_project/other"
	"Library_project/routers"
)

func main() {
	other.ConnectDB()
	routers.Routers()
}
