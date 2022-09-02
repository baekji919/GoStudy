package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main()  {
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		marshal, err := json.Marshal(map[string]string{"text": fmt.Sprintf("%+v", err)})
		if err != nil {
			log.Fatal("Json Marshal Error: ",err)
		}
		requestBody := bytes.NewBuffer(marshal)

		resp, err := http.Post("https://hooks.slack.com/", "application/json", requestBody)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Print("Error : ", err)
	}
	db.Close()
}
