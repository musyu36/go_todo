package controllers

import (
	"fmt"
	"golang/todo_app/config"
	"golang/todo_app/models"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// filesを可変長引数として渡す
	templates := template.Must(template.ParseFiles(files...))
	// layout templateを明示的に使用
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// クッキーから値を受け取り
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		// 受け取ったクッキーのUUIDがデータベースに存在するか判定
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			// 存在しなければエラー生成
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static)) // 静的ファイルの読み込み
	// URLの設定、static という階層に設定にcss, js を設定したいが、実際には無いので、StripPrefix で static を取る
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URL登録(第１引数にURL , 第２引数にハンドラを登録)
	// "/" にアクセスすると top を実行する、(topはroute_mainに記述済み)
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	// サーバ立ち上げ
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
