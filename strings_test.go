package _strings_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zakon47/_strings"
	"testing"
)

func TestNumberLeft(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test1", metka: " 1m", result: []string{"1", "m"}},
		{name: "test1", metka: "2 1m", result: []string{"2", " 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"4", "dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"", "dsd4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := _strings.NumberLeft(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestNumberRight(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test1", metka: " 1m", result: []string{"", "1m"}},
		{name: "test1", metka: "2 1m", result: []string{"", "2 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"", "4dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"4", "dsd"}},
		{name: "test1", metka: "zak47", result: []string{"47", "zak"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := _strings.NumberRight(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestStringLeft(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test1", metka: " 1m", result: []string{"", "1m"}},
		{name: "test1", metka: "2 1m", result: []string{"", "2 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"", "4dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"dsd", "4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := _strings.StringLeft(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestStringRight(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test1", metka: " 1m", result: []string{"m", "1"}},
		{name: "test1", metka: "2 1m", result: []string{"m", "2 1"}},
		{name: "test1", metka: "4dsd", result: []string{"dsd", "4"}},
		{name: "test1", metka: "dsd4", result: []string{"", "dsd4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := _strings.StringRight(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
