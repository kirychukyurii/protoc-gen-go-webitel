package main

import (
	"flag"
	"fmt"
	"github.com/kirychukyurii/protoc-gen-go-webitel/cmd"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "1.0.0"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-webitel %v\n", version)
		return
	}

	protogen.Options{}.Run(cmd.Run(version))
}
