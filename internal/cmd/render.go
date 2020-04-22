package cmd

import (
	"os/exec"
	"strings"
)

func RenderCommand(cmd *exec.Cmd) string {
	var out strings.Builder

	if len(cmd.Env) > 0 {
		out.WriteString("env ")
		out.WriteString(cmd.Env[0])
		out.WriteString(" \\\n")
		for i := 1; i < len(cmd.Env); i++ {
			out.WriteString("    ")
			out.WriteString(cmd.Env[i])
			out.WriteString(" \\\n")
		}
	}

	out.WriteString(cmd.Args[0])
	for i := 1; i < len(cmd.Args); i++ {
		out.WriteByte(' ')
		out.WriteString(cmd.Args[i])
	}

	return out.String()
}
