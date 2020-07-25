package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	VIDEO_DIR             = getAppPath() + "\\videos\\"
	MAX_UPLOAD_SIZE int64 = 500 * 1024 * 1024
)

// 获取绝对路径（go build 所在目录）
func getAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
