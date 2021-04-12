package controllers

import (
	"golang/todo_app/config"
	"net/http"
)

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
