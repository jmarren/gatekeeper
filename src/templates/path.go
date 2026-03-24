package templates

import (
	"strings"

	"github.com/jmarren/gatekeeper/src/util"
)

// returns the path to this directory
func DirPath() string {
	path, err := util.GetBasePath()

	if err != nil {
		panic(err)
	}

	basePath, _ := strings.CutSuffix(path, "path.go")
	return basePath
}
