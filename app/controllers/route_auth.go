package controllers

import (
	"golang/todo_app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		// 登録ボタンが押された時
		err := r.ParseForm() //form の解析
		if err != nil {
			log.Println(err)
		}
		// 値の取得
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		// ユーザー作成しつつ、エラーハンドリング
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		// リダイレクト
		http.Redirect(w, r, "/", 302)
	}
}
