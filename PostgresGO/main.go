package main

func main() {
	db := startConnToDB()
	inputRequest(db)
	defer db.Close()
}
