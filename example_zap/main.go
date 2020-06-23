package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func main() {
	initLogger()
	defer logger.Sync()

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":6060", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	logger.Infof("hello : %s", r.RemoteAddr)
	w.Write([]byte("hello " + r.RemoteAddr))
}

func initLogger() {
	opts := []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	logger = zap.New(zapcore.NewTee(mockCore()), opts...).Sugar()
	logger.Named("test")
}

//自定义时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func mockCore() zapcore.Core {
	// 写文件钩子
	// 文件格式模板
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //zapcore.CapitalColorLevelEncoder
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	// 文件core
	return zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(os.Stdout, newMockWriter()),
		zapcore.DebugLevel,
	)
}

type mockWriter struct {
	logs chan string
}

func newMockWriter() *mockWriter {
	w := &mockWriter{logs: make(chan string, 20)}
	go w.loop()
	return w
}

func (w *mockWriter) loop() {
	for {
		select {
		case <-w.logs:
			log.Println("before mock write")
			time.Sleep(time.Second * 5)
			log.Println("after mock write")
		}
	}
}

func (w *mockWriter) Write(b []byte) (int, error) {
	w.logs <- string(b)
	return len(b), nil
}
func (mockWriter) Sync() error {
	return nil
}
