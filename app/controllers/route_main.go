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

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	// セッションの取得
	sess, err := session(w, r)
	if err != nil {
		// ログインしていなければリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// フォームを解析
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// セッションからユーザーを取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// name アトリビュートで指定したcontent を取得
		content := r.PostFormValue("content")
		// Todo を作成し、エラーハンドリング
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		// リダイレクト
		http.Redirect(w, r, "/todos", 302)
	}
}
