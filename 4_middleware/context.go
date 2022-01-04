package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

type key int

const timingsKey key = 1

type Timing struct {
	Count int
	Duration time.Duration
}

type ctxTimings struct {
	sync.Mutex
	Data map[string]*Timing
}

func trackContextTimings(ctx context.Context, metricName string, start time.Time) {
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	elapsed := time.Since(start)
	timings.Lock()
	defer timings.Unlock()
	if metric, isMetricExist := timings.Data[metricName]; !isMetricExist {
		timings.Data[metricName] = &Timing{
			Count: 1,
			Duration: elapsed,
		}
	} else {
		metric.Count++
		metric.Duration += elapsed
	}
}


func emulateWork(ctx context.Context, workName string) {
	defer trackContextTimings(ctx, workName, time.Now())
	rnd := time.Duration(100 * time.Millisecond)
	time.Sleep(rnd)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	emulateWork(ctx, "lol")
	emulateWork(ctx, "kek")
	emulateWork(ctx, "foo")
	emulateWork(ctx, "foo")
	emulateWork(ctx, "kek")
}

func logContextTimings(ctx context.Context, path string, start time.Time) {
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	totalReal := time.Since(start)
	buf := bytes.NewBufferString(path)
	var total time.Duration
	for timing, value := range timings.Data {
		total += value.Duration
		buf.WriteString(fmt.Sprintf("\n\t(%s) %d %s", timing, value.Count, value.Duration))
	}
	buf.WriteString(fmt.Sprintf("\n\t total: %s", totalReal))
	buf.WriteString(fmt.Sprintf("\n\t tracked: %s", total))
	buf.WriteString(fmt.Sprintf("\n\t unkn: %s", totalReal - total))

	fmt.Println(buf.String())
}

func timingsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx,
			timingsKey,
			&ctxTimings{
			Data: make(map[string]*Timing),
			})
		defer logContextTimings(ctx, r.URL.Path, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootPage)

	handler := timingsMiddleware(r)

	http.ListenAndServe(":8080", handler)
}