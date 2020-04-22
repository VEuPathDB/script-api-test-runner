package env

import "github.com/VEuPathDB/script-api-test-runner/internal/conf"

func BuildEnv(opts *conf.Options) []string {
	return []string {
		BuildUsersEnv(opts),
		BuildAuthEnv(opts),
		BuildSiteUrlEnv(opts),
		BuildTokenEnv(opts),
		BuildVerboseEnv(opts),
	}
}