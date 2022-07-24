package main

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/yimi-go/version/verflag"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	pflag.Parse()
	verflag.PrintAndExitIfRequested()

	var flags flag.FlagSet
	options := protogen.Options{
		ParamFunc: flags.Set,
	}
	options.Run(genFiles)
}
