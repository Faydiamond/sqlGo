package models

import (
	"gomysql/db"
)

type division struct {
	id_division   int64
	Name          string
	weigthFighter float32
	date_creation string
}

type Divisions []division

const Divisionshema string = `
create table  Divisions 
(
	id_division int unsigned auto_increment primary key,
    name varchar(120) not null,
    weigthFighter float not null,
    date_creation timestamp default current_timestamp
)
`

func NewDivision(name string, weiight float32) *division {
	addDivision := &division{Name: name, weigthFighter: weiight}
	return addDivision
}

//insertar registro
func (divi *division) insert() {
	sql := "insert Divisions set name=?, weigthFighter=?"
	db.Exec(sql, divi.Name, divi.weigthFighter)
	//division.id, _ = result.LastInsertId()

}

///crear
func CreateDivision(name string, weiight float32) *division {
	division := NewDivision(name, weiight)
	division.insert()
	return division
}

//listar divisiones

func ListDivisions() Divisions {
	sql := "select id_division,name, weigthFighter, date_creation from Divisions"
	divisiones := Divisions{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		division := division{}
		rows.Scan(&division.id_division, &division.Name, &division.weigthFighter, &division.date_creation)
		divisiones = append(divisiones, division)
	}
	return divisiones
}

func GetDivision(id int) *division {
	divi := NewDivision("", 0)
	sql := " select id_division,name, weigthFighter, date_creation from Divisions where  id_division=?"
	row, _ := db.Query(sql, id)

	for row.Next() {
		row.Scan(&divi.id_division, &divi.Name, &divi.weigthFighter, &divi.date_creation)
	}
	return divi
}

func (divi *division) update() {
	sql := "update  Divisions set name=?, weigthFighter=? where id_division=?"
	db.Exec(sql, divi.Name, divi.weigthFighter, divi.id_division)
	//division.id, _ = result.LastInsertId()

}

//save o edit
func (divi *division) Save() {
	if divi.id_division == 0 {
		divi.insert()
	} else {
		divi.update()
	}
}

//eliminar
func (divi *division) Delete() {
	sql := "delete from Divisions where id_division=?"
	db.Exec(sql, divi.id_division)

}
