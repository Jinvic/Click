package log

import (
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	currentLevel = LevelInfo
	fileLogger   *log.Logger
)

func init() {
	// 设置控制台日志格式
	log.SetFlags(log.LstdFlags)

	// 创建日志目录
	if err := os.MkdirAll("./log", 0755); err != nil {
		log.Fatalln("Failed to create log directory:", err)
	}

	// 创建日志文件
	today := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile("./log/"+today+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	// 创建文件日志器
	fileLogger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func SetLevel(level LogLevel) {
	currentLevel = level
}

func SetLevelFromString(level string) {
	switch strings.ToLower(level) {
	case "debug":
		currentLevel = LevelDebug
	case "info":
		currentLevel = LevelInfo
	case "warn":
		currentLevel = LevelWarn
	case "error":
		currentLevel = LevelError
	default:
		currentLevel = LevelInfo
	}
}

func logToBoth(prefix string, v ...any) {
	msg := append([]any{prefix}, v...)
	log.Println(msg...)
	fileLogger.Println(msg...)
}

func Debug(v ...any) {
	if currentLevel <= LevelDebug {
		logToBoth("[DEBUG]", v...)
	}
}

func Info(v ...any) {
	if currentLevel <= LevelInfo {
		logToBoth("[INFO]", v...)
	}
}

func Warn(v ...any) {
	if currentLevel <= LevelWarn {
		logToBoth("[WARN]", v...)
	}
}

func Error(v ...any) {
	if currentLevel <= LevelError {
		logToBoth("[ERROR]", v...)
	}
}
