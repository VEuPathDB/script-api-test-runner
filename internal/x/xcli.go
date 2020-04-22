package x

import (
	"fmt"
	"github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"
	"github.com/VEuPathDB/lib-go-wdk-api/v0"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/VEuPathDB/script-api-test-runner/internal/conf"
)

const (
	fDescUser = "Site login email address and password formatted as:\n\n" +
		"    -a {email}:{password}\n\n" +
		"This can be specified more than once to provide multiple users.  Leaving" +
		" this out will disable tests that require a user session.  Only " +
		"specifying one set of credentials will disable tests requiring more than " +
		"one user.\n\nThis value can also be specified by setting the" +
		"\"CREDENTIALS\" environment variable to a JSON array of objects " +
		"containing the login email and password set as the properties \"email\" " +
		"and \"pass\" respectively."
	fDescClean = "Clean run.  This runs gradle clean to clear out any remnants " +
		"of previous builds or test runs."
	fDescDebug = "Debug test run script (prints the commands that would be " +
		"run).  If set the test suite will not be run."
	fDescJvm = "Use to set JVM flags."
	fDescForce = "Force tests to run even if they have not changed since the " +
		"last run"
	fDescLegacy = "Use legacy login scheme for authentication instead of OAuth." +
		" Primarily useful for running against local dev sites."
	fDescToken = "QA Token.  \"auth_tkt\" value used to bypass the extra login" +
		" step for QA sites.  This value can be attained by logging into a QA " +
		"site in a browser then retrieving the \"auth_tkt\" value from either the" +
		" query params (or cookies if not in query)"
	fDescTest = "Test class or method to run.  Allows wildcard matching with " +
		"'*'.  Can be specified multiple times to run multiple tests."
	fDescWlTag = "Only run tests annotated with a tag.  Can be specified " +
		"multiple times to create a tag whitelist.\n\nCannot be used with " +
		"--not-tag, if this is set, all --not-tag exclusions will be ignored."
	fDescBlTag = "Exclude tests annotated with a tag.  Can be specified " +
		"multiple times to create a tab blacklist"
	fDescVerbose = "Print verbose test info (including HTTP request details).  " +
		"NOTE: This produces a very large amount of output."
	fDescVersion = "Prints version details for this utility and exits."
	fDescListTags = "Lists the available static tags from the API test suite " +
		"and exits."
	aDescUrl = "URL of the site which will have it's API tested."
	examples = `Examples

  Run in interactive mode (or with env vars) and tell gradle to print
  stacktraces:

    run -- --stacktrace

  Run with single login and site:

    run -l -a some@email.addr:abc123 -u http://username.plasmodb.org/plasmo.username

  Run specific tests:

    run -s 'LoginTest' -s '*StepAnalysisTest$GetAnalysisDetails.invalidUserId'

  Exclude specific tags:

    run -T some,tags,to,exclude

  Run specific tags:

    run -t some,tags,to,run`
)


func ParseParams(version string) *conf.Options {
	out := new(conf.Options)

	com, err := cli.NewCommand().
		Description(examples).
		Flag(slFlag('a', "user", fDescUser).Bind(&out.Users, true)).
		Flag(slFlag('c', "clean-run", fDescClean).Bind(&out.CleanRun, false)).
		Flag(slFlag('d', "debug", fDescDebug).Bind(&out.DebugMode, false)).
		Flag(cli.NewFlag().Short('D').Description(fDescJvm).Bind(&out.JvmPassthrough, true)).
		Flag(slFlag('f', "force", fDescForce).Bind(&out.ForceRun, false)).
		Flag(slFlag('l', "use-legacy-auth", fDescLegacy).Bind(&out.UseLegacyAuth, false)).
		Flag(slFlag('q', "qa-auth-token", fDescToken).Bind(&out.AuthToken, true)).
		Flag(slFlag('s', "test", fDescTest).Bind(&out.Tests, true)).
		Flag(slFlag('t', "only-tag", fDescWlTag).Bind(&out.Tags.Whitelist, true)).
		Flag(slFlag('T', "not-tag", fDescBlTag).Bind(&out.Tags.Blacklist, true)).
		Flag(slFlag('v', "verbose", fDescVerbose).Bind(&out.Verbose, false)).
		Flag(slFlag('V', "version", fDescVersion).OnHit(func(argo.Flag) {
			fmt.Println(version)
			os.Exit(0)
		})).
		Flag(cli.NewFlag().Long("list-tags").Description(fDescListTags).OnHit(listTags)).
		Arg(cli.NewArg().Name("site-url").Bind(&out.SiteUrl).Description(aDescUrl).Require()).
		Parse()

	if err != nil {
		logrus.Fatal(err)
	}

	out.GradlePassthrough = com.Passthroughs()

	tmp, err := wdk.NewApiUrl(out.SiteUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	out.SiteUrl = tmp.String()

	dets, err := wdk.ForceNew(out.SiteUrl).
		UseAuthToken(out.AuthToken).
		GetServiceDetails()
	if err != nil {
		logrus.Fatal(err)
	}
	out.SiteName = dets.ProjectId

	return out
}

func slFlag(s byte, l, d string) argo.FlagBuilder {
	return cli.NewFlag().Short(s).Long(l).Description(d)
}

func listTags(argo.Flag) {
	tags := ReadConcreteTags()
	for _, tag := range tags {
		fmt.Println(tag)
	}
	os.Exit(0)
}