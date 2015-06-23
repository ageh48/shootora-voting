package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func render(w http.ResponseWriter, filename string, data interface{}) {
	t, _ := template.ParseFiles("tmpl/header.html")
	t.Execute(w, data)
	t, _ = template.ParseFiles(filename)
	t.Execute(w, data)
	t, _ = template.ParseFiles("tmpl/footer.html")
	t.Execute(w, data)
}

func renderError(w http.ResponseWriter, message string) {
	render(w, "tmpl/error500.html", map[string]string{
		"message": message,
	})
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	id := r.URL.Query().Get("id")
	user, err := GetUser(id)
	if err != nil || user.Id == "" {
		renderError(w, "IDがおかしいですよ！")
		return
	}
	log.Printf("%#v\n", user)

	if user.Num != 0 {
		render(w, "tmpl/thankyou.html", map[string]interface{}{
			"user": user,
		})
		return
	}

	numString := r.URL.Query().Get("num")
	if numString != "" {
		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			log.Println(err)
			renderError(w, "投票がおかしいですよ！")
			return
		}

		user.Num = int(num)
		if err := user.Save(); err != nil {
			log.Println(err)
			renderError(w, "投票がおかしいですよ！！")
			return
		}

		render(w, "tmpl/thankyou.html", map[string]interface{}{
			"user": user,
		})
	} else {
		render(w, "tmpl/vote.html", map[string]interface{}{
			"user": user,
		})
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetSummary()
	render(w, "tmpl/index.html", map[string]interface{}{
		"data": data,
	})
}

func main() {
	var err error

	db, err = sql.Open("sqlite3", "main.db")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
