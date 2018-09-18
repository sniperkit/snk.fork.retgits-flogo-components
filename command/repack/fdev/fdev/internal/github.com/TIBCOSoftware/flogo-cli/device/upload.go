package device

import (
	"github.com/sniperkit/snk.fork.retgits-flogo-components/command/repack/fdev/fdev/internal/github.com/TIBCOSoftware/flogo-cli/amalgomated_flag"

	"fmt"
	"github.com/sniperkit/snk.fork.retgits-flogo-components/command/repack/fdev/fdev/internal/github.com/TIBCOSoftware/flogo-cli/cli"
	"os"
)

var optUpload = &cli.OptionInfo{
	Name:		"upload",
	UsageLine:	"upload",
	Short:		"upload the device application",
	Long: `Upload the flogo device application.
`,
}

func init() {
	CommandRegistry.RegisterCommand(&cmdUpload{option: optUpload})
}

type cmdUpload struct {
	option *cli.OptionInfo
}

// HasOptionInfo implementation of cli.HasOptionInfo.OptionInfo
func (c *cmdUpload) OptionInfo() *cli.OptionInfo {
	return c.option
}

// AddFlags implementation of cli.Command.AddFlags
func (c *cmdUpload) AddFlags(fs *flag.FlagSet) {
}

// Exec implementation of cli.Command.Exec
func (c *cmdUpload) Exec(args []string) error {

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	return UploadDevice(SetupExistingProjectEnv(appDir))
}
