package cmd

import (
	"github.com/VEuPathDB/script-api-test-runner/internal/cmd/env"
	"github.com/VEuPathDB/script-api-test-runner/internal/conf"
	"os/exec"
	"strings"
)

func BuildGradleCommand(opts *conf.Options, tags []string) *exec.Cmd {
	cmd := exec.Command("./gradlew")
	cmd.Env = env.BuildEnv(opts)

	// The ordering for these is important!
	for _, opt := range opts.JvmPassthrough {
		cmd.Args = append(cmd.Args, "-D"+opt)
	}

	if opts.CleanRun {
		cmd.Args = append(cmd.Args, "clean")
	}

	cmd.Args = append(cmd.Args, "test")

	for _, test := range opts.Tests {
		cmd.Args = append(cmd.Args, "--tests "+test)
	}

	if len(opts.Tags.Whitelist) > 0 {
		cmd.Args = append(cmd.Args, "-PincludeTags="+strings.Join(opts.Tags.Whitelist, ","))
	} else {
		cmd.Args = append(cmd.Args, "-PexcludeTags="+strings.Join(OmitExclusions(opts, tags), ","))
	}

	if opts.Verbose {
		cmd.Args = append(cmd.Args, "--info")
	}

	if opts.ForceRun {
		cmd.Args = append(cmd.Args, "--rerun-tasks")
	}

	for _, arg := range opts.GradlePassthrough {
		cmd.Args = append(cmd.Args, arg)
	}

	return cmd
}

func OmitExclusions(opts *conf.Options, tags []string) []string {
	out := make([]string, 0, len(tags)-3)
	exc := map[string]bool{strings.ToLower(opts.SiteName): true}

	if len(opts.Users) > 0 {
		exc["auth"] = true
	}

	if len(opts.Users) > 1 {
		exc["multi-auth"] = true
	}

	for _, tag := range opts.Tags.Blacklist {
		exc[tag] = true
	}

	for _, tag := range tags {
		if !exc[tag] {
			out = append(out, tag)
		}
	}

	return out
}
