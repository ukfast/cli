package ecloud

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	flaghelper "github.com/ukfast/cli/internal/pkg/helper/flag"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func ecloudSolutionFirewallRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "firewall",
		Short: "sub-commands relating to solution firewalls",
	}

	// Child commands
	cmd.AddCommand(ecloudSolutionFirewallListCmd(f))

	return cmd
}

func ecloudSolutionFirewallListCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists solution firewalls",
		Long:    "This command lists solution firewalls",
		Example: "ukfast ecloud solution firewall list 123",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing solution")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := f.NewClient()
			if err != nil {
				return err
			}

			return ecloudSolutionFirewallList(c.ECloudService(), cmd, args)
		},
	}
}

func ecloudSolutionFirewallList(service ecloud.ECloudService, cmd *cobra.Command, args []string) error {
	solutionID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid solution ID [%s]", args[0])
	}

	params, err := flaghelper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	firewalls, err := service.GetSolutionFirewalls(solutionID, params)
	if err != nil {
		return fmt.Errorf("Error retrieving solution firewalls: %s", err)
	}

	return output.CommandOutput(cmd, OutputECloudFirewallsProvider(firewalls))
}
