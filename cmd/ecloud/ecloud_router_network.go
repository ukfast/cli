package ecloud

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func ecloudRouterNetworkRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network",
		Short: "sub-commands relating to router networks",
	}

	// Child commands
	cmd.AddCommand(ecloudRouterNetworkListCmd(f))

	return cmd
}

func ecloudRouterNetworkListCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists router networks",
		Long:    "This command lists router networks",
		Example: "ukfast ecloud router network list rtr-abcdef12",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing router")
			}

			return nil
		},
		RunE: ecloudCobraRunEFunc(f, ecloudRouterNetworkList),
	}

	cmd.Flags().String("name", "", "Network name for filtering")

	return cmd
}

func ecloudRouterNetworkList(service ecloud.ECloudService, cmd *cobra.Command, args []string) error {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd, helper.NewStringFilterFlagOption("name", "name"))
	if err != nil {
		return err
	}

	networks, err := service.GetRouterNetworks(args[0], params)
	if err != nil {
		return fmt.Errorf("Error retrieving router networks: %s", err)
	}

	return output.CommandOutput(cmd, OutputECloudNetworksProvider(networks))
}
