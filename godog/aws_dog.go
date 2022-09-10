package godog

import (
	"log"
	"os"

	"github.com/DataDog/datadog-go/v5/statsd"
)

const (
	DEFAULT_BUFFER_SIZE int = 500
)

type AwsDogClient struct{}

var client *statsd.Client
var buffer *DDBuffer

func (a *AwsDogClient) RecordSimpleMetric(metricName string, value float64, tags ...string) {
	buffer.Count(metricName, value, getTags(tags...), 1)
}

func (a *AwsDogClient) RecordCompoundMetric(metricName string, value float64, tags ...string) {
	buffer.Gauge(metricName, value, getTags(tags...), 1)
}

func getTags(tags ...string) []string {
	result := make([]string, 0, len(tags)+1)

	if application := os.Getenv("APPLICATION"); application != "" {
		result = append(result, GetRawTag("application", application))
	}

	return append(result, tags...)
}

func init() {
	// Read datadog host from environment variables
	datadogHost, defined := os.LookupEnv("DD_AGENT_HOST")
	if !defined {
		datadogHost = "datadog"
	}

	c, error := statsd.New(datadogHost + ":8125")
	if error != nil {
		log.Print(error)
	}

	client = c
	buffer = CreateBuffer()
}
