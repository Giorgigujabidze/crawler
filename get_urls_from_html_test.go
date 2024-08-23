package main

import (
	"reflect"
	"testing"
)

func Test_getUrlsFromHtml(t *testing.T) {
	type args struct {
		htmlBody   string
		rawBaseUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Some Test",
			args: args{`
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`, "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			wantErr: false,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUrlsFromHtml(tt.args.htmlBody, tt.args.rawBaseUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUrlsFromHtml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUrlsFromHtml() got = %v, want %v", got, tt.want)
			}
		})
	}
}
