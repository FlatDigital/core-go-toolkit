package godog

import (
	"math"
	"strings"
	"time"
)

type Client interface {
	RecordSimpleMetric(metricName string, value float64, tags ...string)
	RecordSimpleTimeMetric(metricName string, fn action, tags ...string) (interface{}, error)
}

type action func() (interface{}, error)

type Tags struct {
	values map[string]string
}

func (t *Tags) Add(key string, value string) *Tags {
	t.init()

	if strings.TrimSpace(key) != "" && strings.TrimSpace(value) != "" {
		t.values[key] = value
	}

	return t
}

func (t *Tags) Remove(key string) *Tags {
	t.init()

	delete(t.values, key)

	return t
}

func (t *Tags) ToArray() []string {
	t.init()

	tags := make([]string, 0)

	for k, v := range t.values {
		tags = append(tags, GetRawTag(k, v))
	}

	return tags
}

func GetRawTag(key string, value string) string {
	return key + ":" + value
}

func (t *Tags) init() {
	if t.values == nil {
		t.values = make(map[string]string)
	}
}

var instance Client

func RecordSimpleMetric(metricName string, value float64, tags ...string) {
	instance.RecordSimpleMetric(metricName, value, tags...)
}

func RecordSimpleTimeMetric(metricName string, fn action, tags ...string) (interface{}, error) {
	return instance.RecordSimpleTimeMetric(metricName, fn, tags...)
}

func takeTime(fn action) (int64, interface{}, error) {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	result, error := fn()
	end := time.Now().UnixNano() / int64(time.Millisecond)
	return (end - start), result, error
}

func takeTimeFloat(fn action) (float64, interface{}, error) {
	initTime := time.Now()
	result, error := fn()
	benchTime := time.Now().Sub(initTime)
	roundedBenchTime := round(benchTime.Seconds()*1000, .5, 3)
	return roundedBenchTime, result, error
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func init() {
	instance = new(AwsDogClient)
}
