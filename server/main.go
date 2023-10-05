package main

import (
	"todo/controller"
	"todo/data"
)

func main() {
	data.SetupDB()

	controller.Router()

}
