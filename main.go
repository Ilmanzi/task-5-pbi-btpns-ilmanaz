package main

import (
	db "github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/database"
	"github.com/Ilmanzi/task-5-pbi-btpns-ilmanaz.git/routes"
)

func main() {
	db.Init()
	r := routes.Routes()

	// Listen and Server in 0.0.0.0:7000
	r.Run(":7000")
}
