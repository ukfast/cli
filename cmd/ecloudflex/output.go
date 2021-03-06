package ecloudflex

import (
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ecloudflex"
)

func OutputECloudFlexProjectsProvider(projects []ecloudflex.Project) output.OutputHandlerDataProvider {
	return output.NewSerializedOutputHandlerDataProvider(projects).WithDefaultFields([]string{"id", "name", "created_at"})
}
