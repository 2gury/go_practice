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
	Count    int
	Duration time.Duration
}

type ctxTimings struct {
	Data map[string]*Timing
	mx   sync.Mutex
}

func trackTime(ctx context.Context, workName string, start time.Time) {
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	elapsed := time.Since(start)
	timings.mx.Lock()
	defer timings.mx.Unlock()
	if ctxTiming, isTimingExist := timings.Data[workName]; !isTimingExist {
		timings.Data[workName] = &Timing{
			Count:    1,
			Duration: elapsed,
		}
	} else {
		ctxTiming.Count++
		ctxTiming.Duration += elapsed
	}
}

func emulateWork(ctx context.Context, workName string) {
	defer trackTime(ctx, workName, time.Now())
	time.Sleep(200 * time.Millisecond)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	emulateWork(ctx, "lol")
	emulateWork(ctx, "kek")
	emulateWork(ctx, "fek")
	emulateWork(ctx, "lol")
	emulateWork(ctx, "kek")
	emulateWork(ctx, "fek")
	emulateWork(ctx, "lol")
	emulateWork(ctx, "kek")
	emulateWork(ctx, "fek")
}

func logTimings(ctx context.Context, start time.Time) {
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	endTime := time.Since(start)
	buf := bytes.NewBufferString("timings")
	var totalTime time.Duration
	for timing, value := range timings.Data {
		totalTime += value.Duration
		buf.WriteString(fmt.Sprintf("\n\t [%s] %d %s ", timing, value.Count, value.Duration))
	}
	buf.WriteString(fmt.Sprintf("\n\t total   : %s", endTime))
	buf.WriteString(fmt.Sprintf("\n\t tracked : %s", totalTime))
	buf.WriteString(fmt.Sprintf("\n\t unknown : %s", endTime-totalTime))

	fmt.Println(buf.String())
}

func timingsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx,
			timingsKey, &ctxTimings{
				Data: make(map[string]*Timing),
				mx:   sync.Mutex{},
			})
		defer logTimings(ctx, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/", rootPage)

	handler := timingsMiddleware(m)
	http.ListenAndServe(":8080", handler)
}
