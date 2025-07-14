package main

import (
	"net/http"

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

	metrics.RegisterPrometheusMetrics()

	r := mux.NewRouter()
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	r.Handle("/metrics", metrics.GetPrometheusHandler())
	r.PathPrefix("/").Handler(handler)

	go func() {
		sugar.Debugln("pprof running on :6060")
		r.Handle(":6060", nil)
	}()

	sugar.Debugln("Service running on :8080")
	sugar.Fatal(http.ListenAndServe(":8080", r))
}
