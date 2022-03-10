package client

import (
	"reflect"
	"testing"
)

func TestTagList(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TagList(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagList() = %v, want %v", got, tt.want)
			}
		})
	}
}
