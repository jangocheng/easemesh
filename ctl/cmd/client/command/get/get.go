package get

import (
	"time"

	"github.com/megaease/easemeshctl/cmd/client/command/meshclient"
	"github.com/megaease/easemeshctl/cmd/client/command/printer"
	"github.com/megaease/easemeshctl/cmd/client/resource"
	"github.com/megaease/easemeshctl/cmd/client/util"
	"github.com/megaease/easemeshctl/cmd/common"

	"github.com/spf13/cobra"
)

type Arguments struct {
	Server       string
	Timeout      time.Duration
	OutputFormat string
}

func Run(cmd *cobra.Command, args *Arguments) {
	switch args.OutputFormat {
	case "table", "yaml", "json":
	default:
		common.ExitWithErrorf("unsupported output format %s (support table, yaml, json)",
			args.OutputFormat)
	}

	visitorBulder := util.NewVisitorBuilder()

	cmdArgs := cmd.Flags().Args()

	switch len(cmdArgs) {
	case 0:
		common.ExitWithErrorf("no resource specified")
	case 1:
		visitorBulder.CommandParam(&util.CommandOptions{
			Kind: cmdArgs[0],
		})
	case 2:
		visitorBulder.CommandParam(&util.CommandOptions{
			Kind: cmdArgs[0],
			Name: cmdArgs[1],
		})
	default:
		common.ExitWithErrorf("invalid command args: support <resource kind> [resource name]")
	}

	vss, err := visitorBulder.Do()
	if err != nil {
		common.ExitWithErrorf("build visitor failed: %s", err)
	}

	printer := printer.New(args.OutputFormat)
	var errs []error
	for _, vs := range vss {
		vs.Visit(func(mo resource.MeshObject, e error) error {
			if e != nil {
				common.OutputErrorf("visit failed: %v", e)
				errs = append(errs, e)
				return nil
			}

			resourceID := mo.Kind()
			if mo.Name() != "" {
				resourceID += "/" + mo.Name()
			}

			objects, err := WrapGetterByMeshObject(mo, meshclient.New(args.Server), args.Timeout).Get()
			if err != nil {
				errs = append(errs, err)
				common.OutputErrorf("%s get failed: %s\n", resourceID, err)
			}

			printer.PrintObjects(objects)

			return nil
		})
	}

	if len(errs) > 0 {
		common.ExitWithErrorf("getting resources has errors occurred")
	}
}