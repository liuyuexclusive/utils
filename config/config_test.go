package config

import (
	"reflect"
	"testing"
)

func TestMustGet(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{"TestMustGet", &Config{Name: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustGet(); !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("want name = %v, got %v", tt.want.Name, got.Name)
			}
		})
	}
}