package conf

type Options struct {
	Users         map[string]string
	CleanRun      bool
	DebugMode     bool
	ForceRun      bool
	UseLegacyAuth bool
	AuthToken     string
	Tests         []string
	Tags          struct {
		Whitelist []string
		Blacklist []string
	}
	SiteUrl string
	Verbose int
	JvmPassthrough []string
	GradlePassthrough []string
	SiteName string
}
