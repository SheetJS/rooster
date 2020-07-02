package filter

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestExtensionsFromReader(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "All lowercase, with dot",
			args: args{
				rd: strings.NewReader("TXT, PDF, DOC, DOCX"),
			},
			want:    []string{".txt", ".pdf", ".doc", ".docx"},
			wantErr: false,
		},
		{
			name: "All caps, with dot",
			args: args{
				rd: strings.NewReader(".TXT, .PDF, .DOC, .DOCX"),
			},
			want:    []string{".txt", ".pdf", ".doc", ".docx"},
			wantErr: false,
		},
		{
			name: "Mixed case, mixed dots",
			args: args{
				rd: strings.NewReader(".txT, pdf, .doc, DOCX"),
			},
			want:    []string{".txt", ".pdf", ".doc", ".docx"},
			wantErr: false,
		},
		{
			name: "Too many headers",
			args: args{
				rd: strings.NewReader("HEADER, OTHER \n .TXT, .PDF"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Not a csv",
			args: args{
				rd: strings.NewReader("Hey there bud,\nHow's it going?"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtensionsFromReader(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtensionsFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtensionsFromReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
