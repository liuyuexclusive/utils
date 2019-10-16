package configutil

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
