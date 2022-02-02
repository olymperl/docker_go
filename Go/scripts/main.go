package main

import(
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
	"net/http"
	"log"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\\nYour HTTP request method is %s\\n", r.Method)
}


func testHandler(w http.ResponseWriter, r *http.Request) {
	cnn, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
        if err != nil {
                log.Fatal(err)
        }

        id := 1
        var name string

        if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
                log.Fatal(err)
        }
	
        fmt.Fprintf(w, "%s is the result for the id %d", name, id)
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/", mainHandler)

        http.ListenAndServe(":8080", nil)
}

