package mdextract

import (
	"testing"
)

func TestGetFiles(t *testing.T) {

	testCases := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "KO - Test 1 - File doesn't exist",
			path:    "fakePath/fakefile.md",
			wantErr: true,
		},
		{
			name:    "OK - Test 2 - File exists",
			path:    "testdata/realfile.md",
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {

			_, err := GetFiles(tc.path)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetFiles() = %v, want %v", err, tc.wantErr)
			}
		})
	}
}
