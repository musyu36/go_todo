package controllers

import (
	"golang/todo_app/config"
	"net/http"
)

func StartMainServer() error {
	// URL登録(第１引数にURL , 第２引数にハンドラを登録)
	// "/" にアクセスすると top を実行する、(topはroute_mainに記述済み)
	http.HandleFunc("/", top)
	// サーバ立ち上げ
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
