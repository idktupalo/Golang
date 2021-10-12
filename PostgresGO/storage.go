package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "*******"
	dbname   = "postgresgo"
)

var idSlice = []int{}

// type databaseUpdate interface {
// 	scanUpdateValues(id int)
// }

// type databaseRowID struct {
// 	id   int
// 	name string
// 	age  int
// }

func startConnToDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func pushValuesToDB(db *sql.DB, id int, val1 string, val2 int) {
	insertData := `insert into "Students"("Id", "Name","Age") values($1, $2, $3)`
	_, err := db.Exec(insertData, id, val1, val2)
	CheckError(err)
}

func deleteValuesDB(db *sql.DB, id int) {
	deleteValue := `delete from "Students" where "Id"=$1`
	_, err := db.Exec(deleteValue, id)
	CheckError(err)
}

func getAllItemsDB(db *sql.DB) {
	rows, err := db.Query(`select * from "Students"`)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var name string
		var id, age int
		err = rows.Scan(&id, &name, &age)
		CheckError(err)
		idSlice = append(idSlice, id)
		fmt.Println(id, name, age)
	}
}

func checkErrId(slice []int, id int) bool {
	for _, val := range slice {
		if val == id {
			return false
		}
	}
	return true
}

func inputRequest(db *sql.DB) {
	var check_str, id int
	var flag bool = true
	for flag {
		printRequests()
		fmt.Scan(&check_str)

		switch check_str {
		case 1:
			id, name, age := scanInsertValues()
			pushValuesToDB(db, id, name, age)
		case 2:
			updateData(db)
		case 3:
			fmt.Println("Input id[type int]")
			fmt.Scan(&id)
			if checkErrId(idSlice, id) {
				fmt.Println("[ Bad request ---> <id not exist> ]")
				os.Exit(1)
			}
			deleteValuesDB(db, id)
		case 4:
			getAllItemsDB(db)
		case 5:
			flag = false
		}
	}
}

func printRequests() {
	fmt.Println("_________________________________________________")
	fmt.Println("|         |         |         |         |       |")
	fmt.Println("|insert(1)|update(2)|delete(3)|select(4)|exit(5)|")
	fmt.Println("|         |         |         |         |       |")
	fmt.Println("_________________________________________________")
	fmt.Print("INPUT REQUEST NUMBER --->  ")
}

func scanInsertValues() (int, string, int) {
	var id, age int
	var name string
	fmt.Println("Input id[type int],name[type string],age[type int]")
	fmt.Scan(&id, &name, &age)
	if !checkErrId(idSlice, id) {
		fmt.Println("[ Bad request ---> <id exist> ]")
		os.Exit(1)
	}
	return id, name, age
}

func updateData(db *sql.DB) {
	printUpdateRequest()
	fmt.Println("Input request number:")
	var requestNum, id, oldId, age int
	var name string
	fmt.Scan(&requestNum)
	switch requestNum {
	case 1:
		fmt.Println("Input Id[type int],Name[type string]")
		fmt.Scan(&id, &age)
		updateData := `update "Students" set "Name"=` + name + ` where "Id"=` + strconv.Itoa(id)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 2:
		fmt.Println("Input Id[type int],Age[type int]")
		fmt.Scan(&id, &age)
		updateData := `update "Students" set "Age"=` + strconv.Itoa(age) + ` where "Id"=` + strconv.Itoa(id)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 3:
		fmt.Println("Input Id[type int],OldID[type int]")
		fmt.Scan(&id, &age)
		updateData := `update "Students" set "Id"=` + strconv.Itoa(id) + ` where "Id"=` + strconv.Itoa(oldId)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 4:
		fmt.Println("Input Id[type int],Name[type string],oldID[type int]")
		fmt.Scan(&id, &name, &oldId)
		updateData := `update "Students" set "Name"=` + name + `,"Id"=` + strconv.Itoa(id) + ` where "Id"=` + strconv.Itoa(oldId)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 5:
		fmt.Println("Input Id[type int],Age[type int],oldID[type int]")
		fmt.Scan(&id, &age, &oldId)
		updateData := `update "Students" set "Name"=` + name + `,"Age"=` + strconv.Itoa(age) + ` where "Id"=` + strconv.Itoa(oldId)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 6:
		fmt.Println("Input Name[type int],Age[type int],ID[type int]")
		fmt.Scan(&name, &age, &id)
		updateData := `update "Students" set "Name"=` + name + `,"Age"=` + strconv.Itoa(age) + ` where "Id"=` + strconv.Itoa(id)
		_, err := db.Exec(updateData)
		CheckError(err)
	case 7:
		fmt.Scan(&id, &name, &age)
		updateAllValueDB(db, id, name, age)
	}
}

func printUpdateRequest() {
	fmt.Println("1)Name")
	fmt.Println("2)Age")
	fmt.Println("3)Id")
	fmt.Println("4)Id,Name")
	fmt.Println("5)Id,Age")
	fmt.Println("6)Name,Age")
	fmt.Println("7)Id,Name,Age")
}

// func (param databaseRowID) updateOneValueDB(db *sql.DB,id int){
// 	updateData:=`update "Students" set "Id"=$1 where "Id"=$2`
// 	_,err:=db.Exec(updateData,param,id)
// 	CheckError(err)
// }

func updateAllValueDB(db *sql.DB, id int, val1 string, val2 int) {
	updateData := `update "Students" set "Name"=$1,"Age"=$2 where "Id"=$3`
	_, err := db.Exec(updateData, val1, val2, id)
	CheckError(err)
}
