package env

import (
	"encoding/base64"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/script-api-test-runner/internal/conf"
)

const (
	envUsersPrefix   = "CREDENTIALS="
	envTokenPrefix   = "QA_AUTH="
	envUrlPrefix     = "SITE_PATH="
	envAuthPrefix    = "AUTH_TYPE="
	envVerbosePrefix = "PRINT_HTTP="

	authLegacy  = "LEGACY"
	authDefault = "OAUTH"
)

func BuildVerboseEnv(opts *conf.Options) string {
	if opts.Verbose {
		return envVerbosePrefix + "true"
	} else {
		return envVerbosePrefix + "false"
	}
}

func BuildAuthEnv(opts *conf.Options) string {
	if opts.UseLegacyAuth {
		return envAuthPrefix + authLegacy
	} else {
		return envAuthPrefix + authDefault
	}
}

func BuildSiteUrlEnv(opts *conf.Options) string {
	return envUrlPrefix + opts.SiteUrl
}

func BuildTokenEnv(opts *conf.Options) string {
	return envTokenPrefix + opts.AuthToken
}

func BuildUsersEnv(opts *conf.Options) string {
	return envUsersPrefix + BuildUsersJson(opts.Users)
}

type User struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func BuildUsersJson(users map[string]string) string {
	tmp := make([]User, 0, len(users))
	enc := base64.StdEncoding

	for k, v := range users {
		tmp = append(tmp, User{
			Email: k,
			Pass:  enc.EncodeToString([]byte(v)),
		})
	}

	out, err := json.Marshal(tmp)
	if err != nil {
		logrus.Fatal("failed to jsonify user credentials", err)
	}
	return string(out)
}
