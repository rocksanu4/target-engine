package main

import (
	"net/http"

	"github.com/go-kit/kit/tree/master/log"
	"go.uber.org/zap"
	"github.com/gorilla/mux"
	_ "net/http/pprof"

	"target-engine/delivery/endpoint"
	"target-engine/delivery/service"
	"target-engine/delivery/transport"
	"target-engine/internal/metrics"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()


	svc := service.NewDeliveryService()
	eps := endpoint.MakeDeliveryEndpoints(svc)
	handler := transport.NewHTTPHandler(eps, logger)

	// Prometheus metrics
	metrics.RegisterPrometheusMetrics()

	// pprof
	go func() {
		sugar.Debugln("pprof running on :6060")
		http.ListenAndServe(":6060", nil)
	}()

	r := mux.NewRouter()
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	r.Handle("/metrics", metrics.GetPrometheusHandler())
	r.PathPrefix("/").Handler(handler)

	sugar.Debugln("Service running on :8080")
	sugar.Fatal(http.ListenAndServe(":8080", r))
}
