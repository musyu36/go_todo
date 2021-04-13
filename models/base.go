package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang/todo_app/config"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

// const (
// 	tableNameUser    = "users"
// 	tableNameTodo    = "todos"
// 	tableNameSession = "sessions"
// )

func init() {
	// heroku の環境変数の値を取り出し
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // users table がなければ作成する
	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME
	// )`, tableNameUser) // tableNameUserは %s に入る
	// // コマンド実行
	// Db.Exec(cmdU)

	// // todos table が無ければ作成する
	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	content TEXT,
	// 	user_id INTEGER,
	// 	created_at DATETIME
	// )`, tableNameTodo)
	// Db.Exec(cmdT)

	// // session table(ログイン情報を保持する)
	// cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	email STRING,
	// 	user_id INTEGER,
	// 	created_at DATETIME
	// )`, tableNameSession)

	// Db.Exec(cmdS)
}

// UUIDの作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードの暗号化
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
