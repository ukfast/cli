package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/cli/internal/pkg/resource"
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/ltaas"
)

func loadtestRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "loadtest",
		Short: "Commands relating to load testing service",
	}

	// Child root commands
	cmd.AddCommand(loadtestDomainRootCmd())
	cmd.AddCommand(loadtestTestRootCmd())
	cmd.AddCommand(loadtestJobRootCmd())

	return cmd
}

// Currently non-functional, as domains aren't yet filterable server-side
type LoadTestDomainLocatorProvider struct {
	service ltaas.LTaaSService
}

func NewLoadTestDomainLocatorProvider(service ltaas.LTaaSService) *LoadTestDomainLocatorProvider {
	return &LoadTestDomainLocatorProvider{service: service}
}

func (p *LoadTestDomainLocatorProvider) SupportedProperties() []string {
	return []string{"name"}
}

func (p *LoadTestDomainLocatorProvider) Locate(property string, value string) (interface{}, error) {
	params := connection.APIRequestParameters{}
	params.WithFilter(connection.APIRequestFiltering{Property: property, Operator: connection.EQOperator, Value: []string{value}})

	return p.service.GetDomains(params)
}

func getLoadTestDomainByNameOrID(service ltaas.LTaaSService, nameOrID string) (ltaas.Domain, error) {
	_, err := uuid.Parse(nameOrID)
	if err != nil {
		locator := resource.NewResourceLocator(NewLoadTestDomainLocatorProvider(service))

		domain, err := locator.Invoke(nameOrID)
		if err != nil {
			return ltaas.Domain{}, fmt.Errorf("Error locating domain [%s]: %s", nameOrID, err)
		}

		return domain.(ltaas.Domain), nil
	}

	domain, err := service.GetDomain(nameOrID)
	if err != nil {
		return ltaas.Domain{}, fmt.Errorf("Error retrieving domain by ID [%s]: %s", nameOrID, err)
	}

	return domain, nil
}

func outputLoadTestDomains(domains []ltaas.Domain) {
	err := Output(NewGenericOutputHandlerProvider(domains, []string{"id", "name"}))
	if err != nil {
		output.Fatalf("Failed to output domains: %s", err)
	}
}

func outputLoadTestTests(tests []ltaas.Test) {
	err := Output(NewGenericOutputHandlerProvider(tests, []string{"id", "name", "number_of_users", "duration", "protocol", "path"}))
	if err != nil {
		output.Fatalf("Failed to output tests: %s", err)
	}
}

func outputLoadTestJobs(jobs []ltaas.Job) {
	err := Output(NewGenericOutputHandlerProvider(jobs, []string{"id", "status", "job_start_timestamp", "job_end_timestamp"}))
	if err != nil {
		output.Fatalf("Failed to output jobs: %s", err)
	}
}

func outputLoadTestJobResults(results []ltaas.JobResults) {
	err := Output(NewGenericOutputHandlerProvider(results, []string{"id", "status", "job_start_timestamp", "job_end_timestamp"}))
	if err != nil {
		output.Fatalf("Failed to output job results: %s", err)
	}
}
