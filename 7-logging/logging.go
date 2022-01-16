package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

type AccessLogger struct {
	StdLogger    *log.Logger
	ZapLogger    *zap.SugaredLogger
	LogrusLogger *logrus.Entry
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (l *AccessLogger) accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := r.FormValue("log")

		next.ServeHTTP(w, r)

		switch logger {
		case "fmt":
			fmt.Printf("FMT [%s] %s %s %s\n",
				r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		case "std":
			l.StdLogger.Printf("FMT [%s] %s %s %s\n",
				r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		case "logrus":
			l.LogrusLogger.WithFields(logrus.Fields{
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
				"work_time":   time.Since(start),
			}).Info(r.URL.Path)
		case "zap":
			l.ZapLogger.Info(r.URL.Path,
				zap.String("method", r.Method),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("url", r.URL.Path),
				zap.Duration("work_time", time.Since(start)))
		default:
		}
	})
}

func main() {
	host := "localhost"
	port := 8080

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	zapLogger.Info("starting server",
		zap.String("logger", "zap"),
		zap.String("host", host),
		zap.Int("port", port))

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false})
	logrus.WithFields(logrus.Fields{
		"logger": "logrus",
		"host":   host,
		"port":   port,
	}).Info("starting server")

	logger := AccessLogger{}
	logger.StdLogger = log.New(os.Stdout, "STD", log.LUTC|log.Lshortfile)
	logger.ZapLogger = zapLogger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)
	logger.LogrusLogger = logrus.WithFields(logrus.Fields{
		"mode":   "[access_log]",
		"logger": "LOGRUS",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := mux.NewRouter()
	r.HandleFunc("/", RootPage)
	handler := logger.accessLogMiddleware(r)
	http.ListenAndServe(":8080", handler)
}
