package check

import (
	"os"
	"testing"
)

func TestCheckWebLinks(t *testing.T) {

	testCases := []struct {
		name    string
		path    string
		weblink string
		wantErr bool
	}{
		{
			name:    "KO - Test 1 - Bad web link",
			path:    "bad_weblink.*.md",
			weblink: "[unknow](https://fake_website_for_test_my_program.com)",
			wantErr: true,
		},
		{
			name:    "OK - Test 2 - Good web link",
			path:    "good_weblink.*.md",
			weblink: "[dimops](https://dimops.com)",
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", tc.path)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tc.weblink)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			err = CheckWebLinks(tmpfile.Name())
			if (err != nil) != tc.wantErr {
				t.Errorf("CheckWebLinks() = %v, want %v", err, tc.wantErr)
			}
		})
	}
}
