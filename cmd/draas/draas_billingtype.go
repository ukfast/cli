package draas

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/factory"
	flaghelper "github.com/ukfast/cli/internal/pkg/helper/flag"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/draas"
)

func draasBillingTypeRootCmd(f factory.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billingtype",
		Short: "sub-commands relating to billing types",
	}

	// Child commands
	cmd.AddCommand(draasBillingTypeListCmd(f))
	cmd.AddCommand(draasBillingTypeShowCmd(f))

	return cmd
}

func draasBillingTypeListCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists billing types",
		Long:    "This command lists billing types",
		Example: "ukfast draas billingtype list",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := f.NewClient()
			if err != nil {
				return err
			}

			return draasBillingTypeList(c.DRaaSService(), cmd, args)
		},
	}
}

func draasBillingTypeList(service draas.DRaaSService, cmd *cobra.Command, args []string) error {
	params, err := flaghelper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	billingtypes, err := service.GetBillingTypes(params)
	if err != nil {
		return fmt.Errorf("Error retrieving billing types: %s", err)
	}

	return output.CommandOutput(cmd, OutputDRaaSBillingTypesProvider(billingtypes))
}

func draasBillingTypeShowCmd(f factory.ClientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "show <billingtype: id>...",
		Short:   "Shows a billing type",
		Long:    "This command shows one or more billing types",
		Example: "ukfast draas billingtype show 00000000-0000-0000-0000-000000000000",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing billing type")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := f.NewClient()
			if err != nil {
				return err
			}

			return draasBillingTypeShow(c.DRaaSService(), cmd, args)
		},
	}
}

func draasBillingTypeShow(service draas.DRaaSService, cmd *cobra.Command, args []string) error {
	var billingtypes []draas.BillingType
	for _, arg := range args {
		billingtype, err := service.GetBillingType(arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving billing type [%s]: %s", arg, err)
			continue
		}

		billingtypes = append(billingtypes, billingtype)
	}

	return output.CommandOutput(cmd, OutputDRaaSBillingTypesProvider(billingtypes))
}
