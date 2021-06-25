package loadbalancer

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/loadbalancer"
)

func loadbalancerTargetGroupACLRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acl",
		Short: "sub-commands relating to ACLs",
	}

	// Child commands
	cmd.AddCommand(loadbalancerTargetGroupACLListCmd(f))

	return cmd
}

func loadbalancerTargetGroupACLListCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "list <acl: id>",
		Short:   "Lists ACLs",
		Long:    "This command lists ACLs",
		Example: "ukfast loadbalancer targetgroup acl list 123",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing target group")
			}

			return nil
		},
		RunE: loadbalancerCobraRunEFunc(f, loadbalancerTargetGroupACLList),
	}
}

func loadbalancerTargetGroupACLList(service loadbalancer.LoadBalancerService, cmd *cobra.Command, args []string) error {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	groupID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid target group ID")
	}

	acls, err := service.GetTargetGroupACLs(groupID, params)
	if err != nil {
		return fmt.Errorf("Error retrieving ACLs: %s", err)
	}

	return output.CommandOutput(cmd, OutputLoadBalancerACLsProvider(acls))
}
