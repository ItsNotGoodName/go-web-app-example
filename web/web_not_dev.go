//go:build !dev

package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed all:dist
var distFS embed.FS

func AssetFS() http.FileSystem {
	return http.FS(echo.MustSubFS(distFS, "dist"))
}

func ReloadVite() {}

func HeadHTML() string {
	// Parse manifest
	file, err := distFS.Open("dist/.vite/manifest.json")
	if err != nil {
		panic(err)
	}

	manifestJSON, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	type Manifest struct {
		CSS     []string `json:"css"`
		File    string   `json:"file"`
		IsEntry bool     `json:"isEntry"`
		Src     string   `json:"src"`
	}

	var manifestMap map[string]Manifest
	if err := json.Unmarshal(manifestJSON, &manifestMap); err != nil {
		panic(err)
	}

	var manifest Manifest
	for _, man := range manifestMap {
		if man.IsEntry {
			manifest = man
			break
		}
	}

	host := ""
	var headTags string
	for _, v := range manifest.CSS {
		headTags += fmt.Sprintf(`<link rel="stylesheet" href="%s/%s" />`, host, v)
	}
	headTags += fmt.Sprintf(`<script type="module" src="%s/%s"></script>`, host, manifest.File)

	return string(headTags)
}
