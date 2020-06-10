package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"text/template"
	"time"

	"github.com/liuyuexclusive/utils/command"
	"github.com/spf13/cobra"
)

type Publish struct {
	ProjectName string
	Type        string
	AppName     string
	Version     string
	Host        string
}

var golangTemplate string = `
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
docker build . -t registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker push registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
ssh root@{{.Host}} "
docker pull registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker stop future.{{.Type}}.basic_1
docker rm future.{{.Type}}.basic_1
docker run -d --network=future_default --name=future.{{.Type}}.basic_1 registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
"
`

var vueTemplate string = `
npm install
npm run build
docker build . -t registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker push registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
ssh root@{{.Host}} "
docker stop future.{{.Type}}.admin_1
docker rm future.{{.Type}}.admin_1
docker pull registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker run -d -p 9090:80 -v /root/future/nginx.conf:/etc/nginx/nginx.conf --network=future_default --name=future.{{.Type}}.admin_1 registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
"
`

var version, host string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		publish()
	},
}

func publish() {
	now := time.Now()

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
	publish.Host = host

	_, err = os.Stat("./publish.txt")
	if err != nil {
		fileInfo, err := os.Create("publish.txt")

		if err != nil {
			log.Fatal(err)
		}
		defer fileInfo.Close()
		var content string
		switch publish.Type {
		case "srv", "web":
			content = golangTemplate
		case "front":
			content = vueTemplate

		}
		fileInfo.WriteString(content)
	}

	tem, err := template.ParseFiles("publish.txt")

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)

	tem.Execute(buf, &publish)

	cmd := buf.String()

	fmt.Println(cmd)

	command.Run("sh", "-c", cmd)

	fmt.Printf("总耗时:%f秒\n", time.Since(now).Seconds())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.publish.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringVar(&version, "v", "latest", "docker image version")
	rootCmd.Flags().StringVar(&host, "h", "49.232.166.55", "host")

}
