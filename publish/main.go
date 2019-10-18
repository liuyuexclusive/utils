package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"text/template"
	"time"

	"github.com/liuyuexclusive/utils/executil"
)

type Publish struct {
	ProjectName string
	Type        string
	AppName     string
	Version     string
}

func main() {
	now := time.Now()

	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := path.Base(dir)

	reg := regexp.MustCompile(`(.+).(srv|web|front).(.+)`)

	if !reg.MatchString(baseDir) {
		panic("错误的项目名称，请使用[project_name].[srv|web|front].[app_name]的格式")
	}

	res := reg.FindAllStringSubmatch(baseDir, -1)

	var publish Publish

	publish.ProjectName = res[0][1]
	publish.Type = res[0][2]
	publish.AppName = res[0][3]
	publish.Version = version

	tem, err := template.ParseFiles("publish.txt")

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)

	tem.Execute(buf, &publish)

	cmd := buf.String()

	fmt.Println(cmd)

	executil.Run("sh", "-c", cmd)

	fmt.Printf("总耗时:%f秒\n", time.Since(now).Seconds())
}

var version string

func init() {
	flag.StringVar(&version, "v", "latest", "镜像版本号")
}
