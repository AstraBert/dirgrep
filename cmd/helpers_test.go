package cmd

import (
	"slices"
	"testing"
)

func TestGetFilesInDir(t *testing.T) {
	testCases := []struct {
		directory string
		recursive bool
		toSkip    []string
		expected  []string
	}{
		{"../testfiles/", true, []string{}, []string{"../testfiles/hello.txt", "../testfiles/loremipsum/part1.txt", "../testfiles/loremipsum/part2.txt", "../testfiles/skip/greetings.txt"}},
		{"../testfiles/", false, []string{}, []string{"../testfiles/hello.txt"}},
		{"../testfiles/", true, []string{"skip"}, []string{"../testfiles/hello.txt", "../testfiles/loremipsum/part1.txt", "../testfiles/loremipsum/part2.txt"}},
	}
	for _, tc := range testCases {
		res, err := getFilesInDir(tc.directory, tc.recursive, tc.toSkip)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if !slices.Equal(res, tc.expected) {
			t.Errorf("Expecting %v as result, got %v", tc.expected, res)
		}
	}
}

func TestGrepMany(t *testing.T) {
	testCases := []struct {
		directory     string
		recursive     bool
		toSkip        []string
		contextWindow int
		pattern       string
		expected      []string
	}{
		{"../testfiles/", true, []string{}, 0, `simple text file(\?|\!)`, []string{"../testfiles/hello.txt", "../testfiles/skip/greetings.txt"}},
		{"../testfiles/", false, []string{}, 0, `simple text file(\?|\!)`, []string{"../testfiles/hello.txt"}},
		{"../testfiles/", true, []string{}, 0, "in voluptate velit", []string{"../testfiles/loremipsum/part1.txt", "../testfiles/loremipsum/part2.txt"}},
		{"../testfiles/", true, []string{}, 1, "in voluptate velit", []string{" [bold red]in voluptate velit[/] "}},
	}
	for _, tc := range testCases {
		if tc.contextWindow == 0 {
			res, err := GrepMany(tc.pattern, tc.directory, tc.recursive, tc.toSkip, tc.contextWindow)
			if err != nil {
				t.Errorf("No expecting any error, got %s", err.Error())
			}
			fls := []string{}
			for k := range res {
				if len(res[k]) > 0 {
					fls = append(fls, k)
				}
			}
			for _, fl := range fls {
				if !slices.Contains(tc.expected, fl) {
					t.Errorf("Expected %s to be among the matched files, it is not", fl)
				}
			}
		} else {
			res, err := GrepMany(tc.pattern, tc.directory, tc.recursive, tc.toSkip, tc.contextWindow)
			if err != nil {
				t.Errorf("No expecting any error, got %s", err.Error())
			}
			matches := []string{}
			for k := range res {
				matches = append(matches, res[k]...)
			}
			for _, m := range matches {
				if m != tc.expected[0] {
					t.Errorf("Expecting to find match %s, got %s", tc.expected[0], m)
				}
			}
		}
	}
}
