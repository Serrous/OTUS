package main

import (
	"reflect"
	"testing"
)

func Test_wordCount(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordCount(tt.args.inputStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
