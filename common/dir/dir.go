package dir

import (
	"os"
	"path/filepath"
)

var (
	root   string
	public string
)

func Init() {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root = cwd
	public = filepath.Join(root, "public")
}

func Public(path string) string {
	return filepath.Join(public, path)
}

func Root(path string) string {
	return filepath.Join(root, path)
}
