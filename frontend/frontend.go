//go:generate npm install
//go:generate echo npm run build
package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist
var staticFilesFs embed.FS
var staticFilesSubFs fs.FS

func init() {
	subFs, err := fs.Sub(staticFilesFs, "dist")
	if err != nil {
		panic(err)
	}
	staticFilesSubFs = subFs
}

func GetStaticFilesRoot() fs.FS {
	return staticFilesSubFs
}
