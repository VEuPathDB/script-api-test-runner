package main

import (
	"fmt"

	"github.com/VEuPathDB/script-api-test-runner/internal/cmd"
	"github.com/VEuPathDB/script-api-test-runner/internal/x"
)

func main() {
	opts := x.ParseParams()
	tags := x.ReadConcreteTags()
	test := cmd.BuildGradleCommand(opts, tags)

	if opts.DebugMode {
		fmt.Println(cmd.RenderCommand(test))
	}
	//fmt.Println(test)
}
