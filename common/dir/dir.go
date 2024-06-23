package dir

import (
	"os"
	"path/filepath"
)

var (
	root   string
	static string
)

func Init() {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root = cwd
	static = filepath.Join(root, "static")
}

func Static(path string) string {
	return filepath.Join(static, path)
}

func Root(path string) string {
	return filepath.Join(root, path)
}
