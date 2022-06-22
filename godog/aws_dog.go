/**
 * @author mlabarinas
 */

package godog

import (
	"os"

	"github.com/DataDog/datadog-go/statsd"
)

const (
	ENDPOINT            string = "datadog:8125"
	DEFAULT_BUFFER_SIZE int    = 500
)

type AwsDogClient struct{}

var client *statsd.Client
var buffer *DDBuffer

func (a *AwsDogClient) RecordSimpleMetric(metricName string, value float64, tags ...string) {
	buffer.Count(metricName, value, getTags(tags...), 1)
}

func (a *AwsDogClient) RecordSimpleTimeMetric(metricName string, fn action, tags ...string) (interface{}, error) {
	time, result, error := takeTime(fn)

	buffer.Count(metricName, float64(time), getTags(tags...), 1)

	return result, error
}

func getTags(tags ...string) []string {
	result := make([]string, 0, len(tags)+3)

	if platform := os.Getenv("PLATFORM"); platform != "" {
		result = append(result, GetRawTag("platform", platform))
	}
	if application := os.Getenv("APPLICATION"); application != "" {
		result = append(result, GetRawTag("application", application))
	}
	if dataCenter := os.Getenv("DATACENTER"); dataCenter != "" {
		result = append(result, GetRawTag("datacenter", dataCenter))
	}

	return append(result, tags...)
}

func init() {
	c, error := statsd.NewBuffered(ENDPOINT, DEFAULT_BUFFER_SIZE)

	if error != nil {
		panic(error)
	}

	client = c
	buffer = CreateBuffer()
	return
}
