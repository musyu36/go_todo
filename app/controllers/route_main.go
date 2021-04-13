package controllers

import (
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
	_, err := session(w, r)
	if err != nil {
		// ログインしていない
		http.Redirect(w, r, "/", 302)
	} else {
		// ログインしている
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
