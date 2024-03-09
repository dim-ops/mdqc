package mdextract

import (
	"os"
	"regexp"
	"testing"
)

func TestGetLinks(t *testing.T) {

	tmpfile, err := os.CreateTemp("", "example.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // nettoyer après

	text := `[unknow](https://example.com)`
	if _, err := tmpfile.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	regex := regexp.MustCompile(`\[(.*?)\]\((https?://[^\s\)]+)\)`)

	links, err := GetLinks([]string{tmpfile.Name()}, regex)
	if err != nil {
		t.Fatal(err)
	}

	// Vérifier le résultat
	if len(links) != 1 || links[0] != "https://example.com" {
		t.Errorf("Expected [\"https://example.com\"], got %v", links)
	}
}

func TestIsValidExtension(t *testing.T) {

	testCases := []struct {
		name          string
		filename      string
		goodExtension bool
	}{
		{
			name:          "KO - Test 1 - Txt file",
			filename:      "file.txt",
			goodExtension: false,
		},
		{
			name:          "OK - Test 2 - Markdown file",
			filename:      "realfile.md",
			goodExtension: true,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			err := isValidExtension(tc.filename)
			if err != tc.goodExtension {
				t.Errorf("isValidExtension() = %v, want %v", err, tc.goodExtension)
			}
		})
	}
}
