package check

import (
	"os"
	"testing"
)

func TestCheckImgLinks(t *testing.T) {

	testCases := []struct {
		name    string
		imgLink string
		path    string
		pathImg string
		wantErr bool
	}{
		{
			name:    "OK - Test 1 - Real file with good image link",
			imgLink: "![golang](testdata/golang.png)",
			path:    "good_img_link.*.md",
			pathImg: "",
			wantErr: false,
		},
		{
			name:    "KO - Test 2 - Real file with bad image link",
			imgLink: "![test](testdata/fake.png)",
			path:    "bad_img_link.*.md",
			pathImg: "",
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("testdata", tc.path)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tc.imgLink)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			err = CheckImgLinks(tmpfile.Name(), tc.pathImg)
			if (err != nil) != tc.wantErr {
				t.Errorf("CheckImgLinks() = %v, want %v", err, tc.wantErr)
			}
		})
	}
}
