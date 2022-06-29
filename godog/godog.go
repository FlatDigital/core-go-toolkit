package godog

import (
	"strings"
)

type Client interface {
	RecordSimpleMetric(metricName string, value float64, tags ...string)
	RecordCompoundMetric(metricName string, value float64, tags ...string)
}

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

func RecordCompoundMetric(metricName string, value float64, tags ...string) {
	instance.RecordCompoundMetric(metricName, value, tags...)
}

func init() {
	instance = new(AwsDogClient)
}
