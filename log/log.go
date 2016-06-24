package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/op/go-logging"
)

var (
	logger   *logging.Logger
	fileName string
)

/// short for
func L() *logging.Logger {
	return Logger()
}

func Logger() *logging.Logger {
	return logger
}

func FileName() string {
	return fileName
}

func init() {
	logger = logging.MustGetLogger("chef_go.log")

	var (
		consoleBackend logging.LeveledBackend
		fileBackend    logging.LeveledBackend
	)

	consoleLogBackend := logging.NewLogBackend(os.Stdout, "", 0)

	consoleFormat := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x} | %{message} |%{color:reset} %{longfile} %{callpath}`,
	)
	consoleFormatter := logging.NewBackendFormatter(consoleLogBackend, consoleFormat)
	consoleBackend = logging.AddModuleLevel(consoleFormatter)

	consoleBackend.SetLevel(logging.DEBUG, "")

	fileName = formatFileName()
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open log file fail", "filename ->", fileName, "err ->", err)
		return
	}
	fileLogBackend := logging.NewLogBackend(f, "", 0)

	fileFormat := logging.MustStringFormatter(
		`%{time:0102 15:04:05} | %{level:.4s} | %{message}`,
	)
	fileFormatter := logging.NewBackendFormatter(fileLogBackend, fileFormat)
	fileBackend = logging.AddModuleLevel(fileFormatter)

	fileBackend.SetLevel(logging.INFO, "")

	logging.SetBackend(consoleBackend, fileBackend)
	return
}

func formatFileName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return fmt.Sprintf(
		"%s.%s.%s.%d.log.chef",
		filepath.Base(os.Args[0]),
		time.Now().Format("20060102T150405"),
		hostname,
		os.Getpid(),
	)
}
