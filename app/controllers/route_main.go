package controllers

import (
	"golang/todo_app/models"
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

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	// セッションの取得
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// フォームを解析
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// 内容を取得
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
