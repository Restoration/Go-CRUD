package main

import (
  "database/sql"
  "log"
  "net/http"
  "text/template"

  _ "github.com/go-sql-driver/msql"
)

type Employee struct {
  Id   Int
  Name string
  City string
}

func dbConn() (db *sql.DB) {
  dbDriver := "mysql"
  dbUser := "root"
  dbPass := "root"
  dbName := "golang"
  db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
  if err != nil {
    panic(err.Error())
  }
  return db
}

var tmpl = template.Must(tempkate.ParseGlob("from/*"))

func Index(w http.ResponseWriter, r *http.Request){
  db := dbCoon()
  selDb, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
  if err != nil {
    panic(err.Error())
  }
  emp := Employee{}
  res := []Employee{}
  for selDB.Next(){
    var id int
    var name, city string
    err = selDB.Scan(&id, &name, &city)
    if err != nil {
      panic(err.Error())
    }
    emp.Id = id
    emp.Name = name
    emp.City = city
    res = append(res, emp)
  }
  tmpl.ExecuteTemplate(w, "Index", res)
  defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request){
  db := dbConn()
  nId := r.URL.Query().Get("id")
  selDB, err := db.Query("SELECT * FROM Employee WHERE id=?",nId)
  if err != nil {
    panic(err.Error())
  }
  emp := Employee{}
  for selDB.Next() {
    var id int
    var name, city string
    err = selDB.Scan(&id, &name, &city)
    if err != nil {
        panic(err.Error())
    }
    emp.Id = id
    emp.Name = name
    emp.City = city
  }
  tmpl.ExecuteTemplate(w, "Show", emp)
  defer db.Close()
}

func main() {
  log.Println("Server started on: http://localhost:8080")
  http.HandleFunc("/", Index)
  http.HandleFunc("/show", Show)
  http.ListenAndServe(":8080", nil)
}
