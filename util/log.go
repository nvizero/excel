package util

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Info(number, brand string) {
	now := time.Now()
	// 格式化時間為字串
	timeStr := now.Format("200601021504")
	// 創建一個新的 FileHook 實例
	log.Out = os.Stdout
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("./logs/log_"+timeStr+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	defer file.Close()
	log.WithFields(logrus.Fields{
		"PartsNumber": number,
		"brand":       brand,
	}).Info("有錯誤的行數")
}
