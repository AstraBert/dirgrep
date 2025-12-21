package cmd

import (
	"os"
	"path"
	"regexp"
	"slices"
	"sync"
)

func grepOne(pattern string, file string, contextWindow int, ch chan<- map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	re, err := regexp.Compile(pattern)
	if err == nil {
		content, err := os.ReadFile(file)
		if err == nil {
			matches := re.FindAllStringIndex(string(content), -1)
			sts := []string{}
			for _, loc := range matches {
				for i := range loc {
					if i%2 == 1 {
						continue
					} else {
						if i < (len(loc) - 1) {
							minC := max(0, loc[i]-contextWindow)
							maxC := min(len(content), loc[i+1]+contextWindow)
							st := string(content[minC:maxC])
							highlighted := re.ReplaceAllString(st, "[bold red]$0[/]")
							sts = append(sts, highlighted)
						}
					}
				}
			}
			if len(sts) > 0 {
				mp := map[string][]string{file: sts}
				ch <- mp
			}
		}
	}
}

func GrepMany(pattern string, directory string, recursive bool, skipDirs []string, contextWindow int) (map[string][]string, error) {
	files, err := getFilesInDir(directory, recursive, skipDirs)
	if err != nil {
		return nil, err
	}
	ch := make(chan map[string][]string)
	var wg sync.WaitGroup
	for _, fl := range files {
		wg.Add(1)
		go grepOne(pattern, fl, contextWindow, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	finalRes := map[string][]string{}
	for item := range ch {
		for k := range item {
			finalRes[k] = item[k]
		}
	}
	return finalRes, nil
}

func getFilesInDir(directory string, recursive bool, skipDirs []string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, path.Join(directory, entry.Name()))
		} else {
			if recursive && !slices.Contains(skipDirs, entry.Name()) {
				subFiles, err := getFilesInDir(path.Join(directory, entry.Name()), recursive, skipDirs)
				if err != nil {
					continue
				} else {
					files = append(files, subFiles...)
				}
			}
		}
	}
	return files, nil
}
