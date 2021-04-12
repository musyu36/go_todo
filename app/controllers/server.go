package controllers

import (
	"fmt"
	"golang/todo_app/config"
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

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static)) // 静的ファイルの読み込み
	// URLの設定、static という階層に設定にcss, js を設定したいが、実際には無いので、StripPrefix で static を取る
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URL登録(第１引数にURL , 第２引数にハンドラを登録)
	// "/" にアクセスすると top を実行する、(topはroute_mainに記述済み)
	http.HandleFunc("/", top)
	// サーバ立ち上げ
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
