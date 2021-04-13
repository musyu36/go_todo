package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// session関数を使ってログインしているか判定
	_, err := session(w, r)
	if err != nil {
		// ログインしていない
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		// ログインしている
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// session関数を使ってログインしているか判定
	sess, err := session(w, r)
	if err != nil {
		// ログインしていない
		http.Redirect(w, r, "/", 302)
	} else {
		// ログインしている
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
