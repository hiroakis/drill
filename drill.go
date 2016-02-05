package drill

import (
	"os"
	"path/filepath"
	"strings"
)

func Down(base string, depth int, abs bool, dir bool) ([]string, error) {

	var files []string

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() && !dir {
			return nil
		}

		rel, _ := filepath.Rel(base, path)
		if len(strings.Split(rel, "/")) > depth {
			return nil
		}
		if abs {
			absPath, _ := filepath.Abs(filepath.Dir(path))
			files = append(files, filepath.Join(absPath, info.Name()))
		} else {
			files = append(files, path)
		}

		return err
	})

	return files, err
}
