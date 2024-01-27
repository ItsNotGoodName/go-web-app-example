//go:build dev

package web

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	cwd                string
	reloadViteFilePath string
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get package directory")
	}

	cwd = filepath.Dir(filename)
	reloadViteFilePath = path.Join(cwd, "reload-vite.local")
}

func ReloadVite() {
	os.Create(reloadViteFilePath)
}

func AssetFS() http.FileSystem {
	return http.FS(os.DirFS(filepath.Join(cwd, "public")))
}

func HeadHTML() string {
	host := os.Getenv("VITE_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	return fmt.Sprintf(`<script type="module" src="http://%s:5173/@vite/client"></script><script type="module" src="http://%s:5173/src/main.ts"></script>`, host, host)
}
