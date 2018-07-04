package main

import (
	"bytes"
	"fmt"
	"flag"

	"github.com/hb-go/micro/pkg/file"
	"os/exec"
	"strings"
	"os"
)

var (
	cmdHelp = flag.Bool("h", false, "帮助")
	root    = flag.String("r", "", "需要查找.proto文件的路径，多路径使用\":\"分隔")
)

func init() {
	flag.Parse()
}

func main() {
	if *cmdHelp {
		flag.PrintDefaults()
		return
	}

	roots := strings.Split(*root, ":")
	if *root == "" || len(roots) <= 0 {
		fmt.Println("-r 请输入正确的路径")
		return
	}

	files := file.FileSlice{}
	nbytes := int64(0)
	file.WalkDirs(roots, ".proto", &files, &nbytes)

	fmt.Printf("proto files count:%d, size:%.3f kb\n", files.Len(), float64(nbytes)/1e3)

	arg := []string{}

	arg = append(arg, "-I="+os.Getenv("GOPATH")+"/src:.")
	arg = append(arg, "--micro_out=.", "--go_out=.")

	var out bytes.Buffer
	var stderr bytes.Buffer
	for _, v := range files {
		fmt.Println("proto file path:" + v.Path)

		cmd := exec.Command("protoc", append(arg, v.Path)...)
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("exec cmd error:" + fmt.Sprint(err) + "\n" + stderr.String())
			continue
		}

		fmt.Println("exec cmd success\n" + out.String())
	}
}
