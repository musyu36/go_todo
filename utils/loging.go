package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// O_RDWR 読み書き
	// O_CREATE なければ作成
	// O_APPEND 追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //0666 permition
	if err != nil {
		log.Fatalln(err)

	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // ログの書き込み先Stdout と logfile に設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // フォーマット指定
	log.SetOutput(multiLogFile)                          // ログの出力先を設定
}
