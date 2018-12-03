package main
import (
	"database/sql"
    "log"
    "net/http"
    "text/template"
	"io"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)
type Notes struct {
    Id    int
    Description  string
    Location string
}
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "notes"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/show", Show)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
var tmpl = template.Must(template.ParseGlob("form/*"))
func index(w http.ResponseWriter, r *http.Request) {
io.WriteString(w, "hello from a docker container")
db := dbConn()
fmt.Println("MYSQL Tutorial")
//insert, err := db.Prepare("INSERT INTO notes VALUES(?,?,?)")
selDB, err := db.Query("SELECT * FROM notes")
if err != nil {
	panic(err.Error())
}
emp := Notes{}
res := []Notes{}
for selDB.Next() {
	var ID int
	var description0, location string
	err = selDB.Scan(&ID, &description0, &location)
	if err != nil {
		panic(err.Error())
	}
	emp.Id = ID
	emp.Description = description0
	emp.Location = location
	res = append(res, emp)
}
tmpl.ExecuteTemplate(w, "Index", res)
defer db.Close()
fmt.Println("Successfully inserted into user tables")
}
func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("ID")
    selDB, err := db.Query("SELECT * FROM notes WHERE ID=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Notes{}
    for selDB.Next() {
        var ID int
        var description0, location string
        err = selDB.Scan(&ID, &description0, &location)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = ID
        emp.Description = description0
        emp.Location = location
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}
func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("ID")
    selDB, err := db.Query("SELECT * FROM notes WHERE ID=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Notes{}
    for selDB.Next() {
        var ID int
        var description0, location string
        err = selDB.Scan(&ID, &description0, &location)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = ID
        emp.Description = description0
        emp.Location = location
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}
func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        description0 := r.FormValue("description0")
        location := r.FormValue("location")
        insForm, err := db.Prepare("INSERT INTO notes(description0, location) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(description0, location)
        log.Println("INSERT: Description: " + description0 + " | location: " + location)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        description0 := r.FormValue("description0")
        location := r.FormValue("location")
        ID := r.FormValue("Id")
        insForm, err := db.Prepare("UPDATE notes SET description0=?, location=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(description0, location, ID)
        log.Println("UPDATE: description: " + description0 + " | location: " + location)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("ID")
    delForm, err := db.Prepare("DELETE FROM Employee")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}