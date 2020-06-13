package db

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestOpen(t *testing.T) {
	type args struct {
		f func(*DB) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {"TestOpen", args{f: func(db *gorm.DB) error { return db.DB().Ping() }}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Open(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
