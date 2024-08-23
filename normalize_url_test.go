package main

import "testing"

func Test_normalizeUrl(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "remove scheme",
			args:    args{"https://blog.boot.dev/path"},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},

		{
			name:    "remove scheme",
			args:    args{"http://blog.boot.dev/path/"},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeUrl(tt.args.rawUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("normalizeUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("normalizeUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
