package util

import (
	"flag"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func FileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ListFiles(path string) []string {
	res := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			res = append(res, f.Name())
		}
	}

	return res
}

func GetFileAbsolutePath(relativePath string) (string, error) {
	var absolutePath string
	appPath := GetAppPath()
	absolutePath = filepath.Join(appPath, relativePath)

	defer func() {
		if err := recover(); err != nil {
			absolutePath = relativePath
			logs.Error("GetFileAbsolutePath error: ", err)
		}
	}()

	return strings.TrimSpace(absolutePath), nil
}

var appPath string

func GetAppPath() string {

	if StrIsEmpty(appPath) {
		flag.StringVar(&appPath, "app-path", "app-path", "get app-path")
		flag.Parse()
	}

	if StrIsEmpty(appPath) || appPath == "app-path" {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		index := strings.LastIndex(path, string(os.PathSeparator))
		return path[:index]
	}

	return appPath
}
