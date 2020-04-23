package env

import (
	"github.com/VEuPathDB/script-api-test-runner/internal/conf"
	"os"
)

func BuildEnv(opts *conf.Options) []string {
	out := []string {
		BuildUsersEnv(opts),
		BuildAuthEnv(opts),
		BuildSiteUrlEnv(opts),
		BuildTokenEnv(opts),
		BuildVerboseEnv(opts),
	}

	if val, ok := os.LookupEnv("JAVA_HOME"); ok {
		out = append(out, "JAVA_HOME=" + val)
	}

	return out
}