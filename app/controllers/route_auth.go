package controllers

import (
	"golang/todo_app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// サインアップページはログインしていない状態なら表示する
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
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

func login(w http.ResponseWriter, r *http.Request) {
	// ログインページはログインしていない状態なら表示する
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

}

// ログインページに入力されたemail, password で認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	// ユーザーをemailから取得
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	// 暗号化して保存しているので、暗号化してからパスワード比較
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		// 成功したらセッションとクッキーを作成し、トップにリダイレクト
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		// 失敗したらログイン画面にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	// クッキーの取得
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	// ErrNoCookie でなければ
	if err != http.ErrNoCookie {
		// cookie.Value つまり UUID から session struct を作成しておく
		session := models.Session{UUID: cookie.Value}
		session.DeleteSEssionByUUID()

	}
	http.Redirect(w, r, "/login", 302)
}
