package utils

import (
	"path/filepath"
	"runtime"
)

// GetProjectRoot fetches the root directory of the project
func GetProjectRoot() (root string) {
	_, callerFile, _, _ := runtime.Caller(0)

	root = filepath.Join(filepath.Dir(callerFile), "../")

	return
}
