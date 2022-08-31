package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	_, err := sql.Open("mysql", "root:1111@tcp(127.0.0.1:3306)/data")
	if err != nil {
		fmt.Print("Error : ", err)
	}

	fmt.Println("Success")
}
