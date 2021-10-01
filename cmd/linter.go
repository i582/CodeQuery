package cmd

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/VKCOM/noverify/src/cmd"
	"github.com/VKCOM/noverify/src/linter"
	"github.com/i582/CodeQuery/pkg/models"
	"github.com/i582/CodeQuery/pkg/pipes/collect"
	"github.com/i582/CodeQuery/pkg/pipes/collect/walkers"
	"github.com/i582/CodeQuery/pkg/pipes/data/preprocess"
	"github.com/i582/CodeQuery/pkg/utils"
)

func RunLinterForPaths(paths []string) *models.Database {
	log.SetFlags(0)
	log.Println("Collecting data for building a database")

	config := linter.NewConfig("7.4")
	context := walkers.NewGlobalContext(nil)

	oldArgs := os.Args
	newArgs := []string{os.Args[0]}
	newArgs = append(newArgs, "collect")
	os.Args = append(newArgs, paths...)

	_, err := cmd.Run(&cmd.MainConfig{
		LinterConfig: config,
		BeforeReport: func(report *linter.Report) bool {
			return false
		},
		DisableCriticalIssuesLog: true,
		AfterFlagParse: func(env cmd.InitEnvironment) {
			context.Info = env.MetaInfo
		},
		ModifyApp: func(app *cmd.App) {
			app.Name = "CodeQuery"
			app.Description = ""

			// Clear all the default commands.
			app.Commands = nil

			app.Commands = append(app.Commands, &cmd.Command{
				Name:        "collect",
				Description: "The command to start checking files",
				Arguments: []*cmd.Argument{
					{
						Name:        "targets",
						Description: "Folders or files for analysis",
					},
				},
				RegisterFlags: func(ctx *cmd.AppContext) (*flag.FlagSet, *cmd.FlagsGroups) {
					fs := flag.NewFlagSet("check", flag.ContinueOnError)

					// We don't need all the flags from NoVerify, so we only register some of them.

					ctx.ParsedFlags.MaxFileSize = 10 * 1024 * 1024
					ctx.ParsedFlags.MaxConcurrency = runtime.NumCPU()
					ctx.ParsedFlags.StubsDir = ""
					ctx.ParsedFlags.CacheDir = utils.DefaultCacheDir()
					ctx.ParsedFlags.DisableCache = false

					ctx.ParsedFlags.IndexOnlyFiles = ""
					ctx.ParsedFlags.PhpExtensionsArg = "php,inc,php5,phtml"

					ctx.ParsedFlags.PHP7 = true

					// Some values need to be set manually.
					ctx.ParsedFlags.AllowChecks = ""
					ctx.ParsedFlags.AllowAll = false
					ctx.ParsedFlags.ReportsCritical = ""

					return fs, cmd.NewFlagsGroups()
				},
				Action: func(ctx *cmd.AppContext) (int, error) {
					return collect.Collect(ctx, context)
				},
			})
		},
	})
	if err != nil {
		log.Println(err)
	}

	os.Args = oldArgs

	data := preprocess.Run(context)

	return data
}
