package ecloud_v2

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func ecloudDHCPRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dhcp",
		Short: "sub-commands relating to dhcps",
	}

	// Child commands
	cmd.AddCommand(ecloudDHCPListCmd(f))
	cmd.AddCommand(ecloudDHCPShowCmd(f))

	return cmd
}

func ecloudDHCPListCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists dhcps",
		Long:    "This command lists dhcps",
		Example: "ukfast ecloud dhcp list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := f.NewClient()
			if err != nil {
				return err
			}

			return ecloudDHCPList(c.ECloudService(), cmd, args)
		},
	}

	cmd.Flags().String("name", "", "DHCP name for filtering")

	return cmd
}

func ecloudDHCPList(service ecloud.ECloudService, cmd *cobra.Command, args []string) error {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	helper.HydrateAPIRequestParametersWithStringFilterFlag(&params, cmd, "name", "name")

	dhcps, err := service.GetDHCPs(params)
	if err != nil {
		return fmt.Errorf("Error retrieving dhcps: %s", err)
	}

	return output.CommandOutput(cmd, OutputECloudDHCPsProvider(dhcps))
}

func ecloudDHCPShowCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "show <dhcp: id>...",
		Short:   "Shows a dhcp",
		Long:    "This command shows one or more dhcps",
		Example: "ukfast ecloud dhcp show 123",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing dhcp")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := f.NewClient()
			if err != nil {
				return err
			}

			return ecloudDHCPShow(c.ECloudService(), cmd, args)
		},
	}
}

func ecloudDHCPShow(service ecloud.ECloudService, cmd *cobra.Command, args []string) error {
	var dhcps []ecloud.DHCP
	for _, arg := range args {
		dhcp, err := service.GetDHCP(arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving dhcp [%s]: %s", arg, err)
			continue
		}

		dhcps = append(dhcps, dhcp)
	}

	return output.CommandOutput(cmd, OutputECloudDHCPsProvider(dhcps))
}
