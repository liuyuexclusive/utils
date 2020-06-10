package appconfig

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
				t.Errorf("MustGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{"TestGet", &Config{Name: "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
