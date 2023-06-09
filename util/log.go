package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Info(number string) {
	// 建立 logrus 的 Logger 物件
	logger := logrus.New()

	// 設定日誌輸出至標準輸出（stdout）
	logger.SetOutput(os.Stdout)

	// 設定日誌層級
	logger.SetLevel(logrus.DebugLevel)

	// 記錄不同層級的日誌訊息
	logger.Trace("This is a trace message" + number)
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
