package compress

import "testing"

func TestCompressImg(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "OK - Test 1 - Compress PNG Image",
			path:    "testdata/golang.png",
			wantErr: false,
		},
		{
			name:    "OK - Test 3 - Compress JPG Image",
			path:    "testdata/golang.png",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompressImg(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("CompressImg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
