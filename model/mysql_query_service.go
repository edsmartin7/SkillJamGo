package model

import (
   "database/sql"
   "fmt"

   _ "github.com/go-sql-driver/mysql"
)

func OpenDB() sql.DB {

   db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
   if err != nil {
      fmt.Println("Error in opening local database: ", err)
   }

   return *db
}

func VerifyLogin(username string, password string) bool {

   db := OpenDB()

   verify, err := db.Query(CheckCredentials, username, password)
   if err != nil {
      fmt.Println("Error in query to verify login", err)
   }
   var un string
   var pw string
   for verify.Next() {
      err = verify.Scan(&un, &pw)
   }
   db.Close()

   return username==un && password==pw
}

