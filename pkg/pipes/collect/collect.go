package collect

import (
	"github.com/VKCOM/noverify/src/cmd"
	"github.com/i582/CodeQuery/pkg/pipes/collect/walkers"
)

// Collect is the function that starts the analysis of the project.
func Collect(ctx *cmd.AppContext, globalContext *walkers.GlobalContext) (status int, err error) {
	// Registering custom walkers for collecting the call graph.
	walkers.Register(ctx.MainConfig.LinterConfig, globalContext, "kphp-color")

	// If there are no arguments, then we interpret this as
	// an analysis of the current directory.
	if len(ctx.ParsedArgs) == 0 {
		ctx.ParsedArgs = append(ctx.ParsedArgs, "./")
	}

	// The main function for analyzing in NoVerify,
	// in it, we collect all the functions of the project.
	_, err = cmd.Check(ctx)
	if err != nil {
		return 2, err
	}

	// If the status is not zero, it means that there are
	// some errors at the stage of data collection.
	//
	// No further analysis is needed.
	if status != 0 {
		return status, nil
	}

	return 0, nil
}
