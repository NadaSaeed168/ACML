package main
import (
	//"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	"io"
	"net/http"
	//"fmt"
	//"time"
)
type PageVariables struct {
	Date string
	Time string
}
func main() {
http.HandleFunc("/", index)
http.ListenAndServe(":80", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
io.WriteString(w, "hello from a docker container")
io.WriteString(w, "Go MYSQL Tutorial")
db, err := sql.Open("mysql", "root:password@tcp(:80)/database")
if err != nil{
	panic(err.Error())
}
defer db.Close()
io.WriteString(w,"MYSQL Tutorial")
insert, err := db.Prepare("INSERT INTO notes VALUES('5ara el senin','GUC')")
if err != nil{
	panic(err.Error())
	io.WriteString(w,"MYSQL")
}
io.WriteString(w,"MYSQL Table")
defer insert.Close()
io.WriteString(w,"Successfully inserted into user tables")
// now := time.Now() // find the time right now
// HomePageVars := PageVariables{ //store the date and time in a struct
// Date: now.Format("02-01-2006"),
// Time: now.Format("15:04:05"),
// }
// t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
// if err != nil { // if there is an error
// log.Print("template parsing error: ", err) // log it
// }
// err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
// if err != nil { // if there is an error
// log.Print("template executing error: ", err) //log it
// }
}