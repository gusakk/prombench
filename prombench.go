package prombench

import (
	"context"
	"fmt"
	"github.com/ncabatoff/prombench/harness"
	"github.com/ncabatoff/prombench/loadgen"
	"github.com/prometheus/client_golang/api/prometheus"
	"github.com/prometheus/common/model"
	"log"
	"time"
)

const (
	SdCfgDir = "sd_configs"
)

func Run(cfg harness.Config) {
	mainctx := context.Background()
	genbuilder := func() loadgen.MetricsGenerator {
		return loadgen.NewTestCollector(100, 100)
	}
	le := loadgen.NewLoadExporterInternal(mainctx, SdCfgDir, genbuilder)
	for i := 0; i < cfg.NumExporters; i++ {
		if err := le.AddTarget(cfg.FirstPort + i); err != nil {
			log.Fatalf("Error starting exporter: %v", err)
		}
	}

	harness.SetupDataDir("data", cfg.Rmdata)
	harness.SetupPrometheusConfig(SdCfgDir, cfg.ScrapeInterval)
	stopPrometheus := harness.StartPrometheus(mainctx, cfg.PrometheusPath, cfg.ExtraArgs)

	startTime := time.Now().Truncate(time.Minute)
	time.Sleep(cfg.TestDuration)
	expectedSum, err := le.Stop()
	log.Printf("sum=%d, err=%v", expectedSum, err)
	// TODO make delay before query configurable
	time.Sleep(5 * time.Second)
	ttime := time.Since(startTime)
	ttimestr := fmt.Sprintf("%ds", int(1+ttime.Seconds()))
	query := fmt.Sprintf(`sum(sum_over_time({__name__=~"test.+"}[%s]))`, ttimestr)
	vect := queryPrometheusVector("http://localhost:9090", query)
	actualSum := -1
	if len(vect) > 0 {
		actualSum = int(vect[0].Value)
	}
	if expectedSum != actualSum {
		log.Printf("Expected %d, got %d", expectedSum, actualSum)
	}
	stopPrometheus()
}

func queryPrometheusVector(url, query string) model.Vector {
	cfg := prometheus.Config{Address: url, Transport: prometheus.DefaultTransport}
	client, err := prometheus.New(cfg)
	if err != nil {
		log.Fatalf("error building client: %v", err)
	}
	qapi := prometheus.NewQueryAPI(client)
	log.Printf("issueing query: %s", query)
	result, err := qapi.Query(context.TODO(), query, time.Now())
	if err != nil {
		log.Printf("error performing query: %v", err)
		return nil
	}
	log.Printf("prometheus query result: %v", result)
	return result.(model.Vector)
}
