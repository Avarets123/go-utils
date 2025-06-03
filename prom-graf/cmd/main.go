package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"prom-graf/pkg"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace   = "test-app"
	nGoMetName  = "/sched/goroutines:goroutines"
	nMemMetName = "/memory/classes/heap/free:bytes"
)

func main() {

	nGo := pkg.NewGauge(namespace, "nGo", "Goroutines count")
	mUsage := pkg.NewGauge(namespace, "mUsage", "Memory usage")
	histogram := pkg.NewHistogram(namespace, "histogram", "this is my Histogram")

	prometheus.MustRegister(nGo, mUsage, histogram)

	rand.NewSource(time.Now().Unix())

	m := []metrics.Sample{
		{Name: nGoMetName}, {Name: nMemMetName},
	}

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {

			for range 4 {
				go func() {
					_ = make([]int, 1000000)
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			runtime.GC()
			metrics.Read(m)
			goCount := m[0].Value.Uint64()
			mUsCount := m[1].Value.Uint64()

			time.Sleep(time.Duration(rand.Intn(15)) * time.Second)

			nGo.Set(float64(goCount))
			mUsage.Set(float64(mUsCount))

		}
	}()

	fmt.Println("Port 2222")

	panic(http.ListenAndServe(":2222", nil))

}

func simpleMetrics() {

	const nGo = "/sched/goroutines:goroutines"
	for range 10000 {
		go func() {
			time.Sleep(4 * time.Second)
		}()
	}

	time.Sleep(2 * time.Second)
	m := make([]metrics.Sample, 1)
	m[0].Name = nGo

	metrics.Read(m)

	fmt.Printf("%+v\n", m)

	// f.Write(b)

}
