package main

import (
	"flag"
	"fmt"
	"github.com/c0dehvb/go-runlike/runlike"
	"github.com/c0dehvb/go-runlike/runlike/utils"
	"os"
	"runtime"
	"strings"
)

func init() {
	flag.StringVar(&container, "c", "", "[Required] ContainerID or ContainerName")
	flag.BoolVar(&pretty, "pretty", false, "Print pretty format (default false)")
	flag.StringVar(&platform, "os", "", "Specific OS platform format (windows or linux)")
}

var (
	container string
	pretty bool
	platform string
)

func main() {
	flag.Parse()
	if container == "" {
		fmt.Println("flag 'container' cannot be empty")
		flag.Usage()
		os.Exit(1)
	}

	result, err := utils.ContainerInspectMap(container)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if platform == "" {
		platform = runtime.GOOS
	}
	fmt.Println(runlike.Inspect(result, pretty, strings.ToLower(platform)))
}