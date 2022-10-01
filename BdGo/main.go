package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/db/models"
)

func main() {
	db.Connect()

	division := models.ListDivisions()
	fmt.Println(division)

	divi := models.GetDivision(2)
	fmt.Println("Divi  ", divi)

	divi.Name = "peso moscaa "
	divi.Save()
	fmt.Println(models.ListDivisions())

	divi.Delete()
	fmt.Println(models.ListDivisions())
	/*
			db.Ping()
			db.CreateTables(models.DivisionShema, "divisions")
			db.DeleteTable("MYTABLE22")


			divission := models.CreateDivision("Peso Gallo", 61.2)
		fmt.Println(divission)
	*/
	db.Close()
}
