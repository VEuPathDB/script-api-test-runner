package main

import (
	"fmt"
	"os"

	"github.com/VEuPathDB/script-api-test-runner/internal/cmd"
	"github.com/VEuPathDB/script-api-test-runner/internal/x"
)

var version string

func main() {
	opts := x.ParseParams(version)
	tags := x.ReadConcreteTags()
	test := cmd.BuildGradleCommand(opts, tags)

	if opts.DebugMode {
		fmt.Println(cmd.RenderCommand(test))
	} else {
		test.Stdout = os.Stdout
		test.Stderr = os.Stderr
		if err := test.Run(); err != nil {
			panic(err)
		}
	}
}
