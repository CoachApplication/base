package configconnector

import (
	"test"
)

var TESTVALUES = map[string]map[string]string{
	"one": map[string]string{
		"A": "A",
		"B": "B",
		"C": "C",
	},
	"two": map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	},
	"three": map[string]string{
		"1": "3",
		"2": "2",
		"3": "1",
	},
}

type TestConfigConnector struct {
	values map[string]map[string]string
}

// NewTestConfigConnect creates a testing config connector
func NewTestConfigConnector() *TestConfigConnector {
	return &TestConfigConnector{values: TESTVALUES}
}
