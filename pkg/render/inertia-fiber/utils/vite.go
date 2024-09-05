package utils

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Vite - vite returns the HTML for the specified entrypoints.
func Vite(entrypoints []string, buildDirectory ...string) template.HTML {
	var builder strings.Builder

	log.Println(entrypoints, buildDirectory)

	log.Println(isRunningHot(buildDirectory))

	// If running in hot mode, return the HTML for the hot asset.
	if isRunningHot(buildDirectory) {
		html := makeTagForChunk(hotAsset("@vite/client", buildDirectory...))
		log.Println(html)
		for _, v := range entrypoints {
			log.Println("isRunningHot -> entrypoints", hotAsset(v, buildDirectory...))
			html += makeTagForChunk(hotAsset(v, buildDirectory...))
		}
		log.Println("hot", html)
		return html
	}

	// Otherwise, return the HTML for the specified entrypoints.
	manifest := manifest(buildDirectory)

	log.Println("Vite manifest:", manifest)

	log.Println("entrypoints:", entrypoints)

	for _, v := range entrypoints {
		m, exists := manifest[v]
		if !exists {
			continue
		}

		log.Println(m)

		for _, css := range m.Css {
			builder.WriteString(string(makeStylesheetTag(css)))
		}

		log.Println("entrypoint", m.File)

		builder.WriteString(string(makeScriptTag(m.File)))
	}

	return template.HTML(builder.String())
}

// makeTagForChunk returns the HTML tag for the specified URL.
func makeTagForChunk(url string) template.HTML {
	if isCssPath(url) {
		return template.HTML(makeStylesheetTag(url))
	}

	return template.HTML(makeScriptTag(url))
}

// makeStylesheetTag returns the HTML tag for the specified stylesheet URL.
func makeStylesheetTag(url string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, url)
}

// makeScriptTag returns the HTML tag for the specified script URL.
func makeScriptTag(url string) string {
	return fmt.Sprintf(`<script type="module" src="%s"></script>`, url)
}

// isCssPath returns true if the path is a CSS path.
func isCssPath(path string) bool {
	return regexp.MustCompile(`\.(css|less|sass|scss|styl|stylus|pcss|postcss)$`).MatchString(path)
}

// hotAsset returns the hot asset for the specified asset.
func hotAsset(asset string, buildDirectory ...string) string {
	data, err := os.ReadFile(hotFile(buildDirectory...))
	if err != nil {
		panic(err)
	}

	//log.Println(data)
	//log.Println(path.Join(buildDirectory...))
	//return strings.TrimSuffix(string(data), "\n") + "/" + asset
	if strings.Index(asset, "vite") == 1 {
		return strings.TrimSuffix(string(data), "\n") + "/" + asset
	}

	return strings.TrimSuffix(string(data), "\n") + "/" + strings.TrimPrefix(asset, "././internal/control_panel/")
	//return "/" + strings.TrimPrefix(asset, "././internal/control_panel/")
	//return "/" + strings.TrimPrefix(asset, "././internal/control_panel/")
	//return asset
}

// isRunningHot returns true if running in hot mode.
func isRunningHot(buildDirectory []string) bool {
	filename := hotFile(buildDirectory...)
	if filename == "" {
		return false
	}
	_, err := os.Stat(filename)

	return !os.IsNotExist(err)
}

// hotFile returns the hot file path.
func hotFile(buildDirectory ...string) string {
	sfs := append(buildDirectory, "public", "hot")
	return filepath.Join(sfs...)
}
