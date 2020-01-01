package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/cli/test/test_output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func Test_ecloudPodApplianceListCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := ecloudPodApplianceListCmd().Args(nil, []string{"123"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := ecloudPodApplianceListCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing pod", err.Error())
	})
}

func Test_ecloudPodApplianceList(t *testing.T) {
	t.Run("DefaultRetrieve", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetPodAppliances(123, gomock.Any()).Return([]ecloud.Appliance{}, nil).Times(1)

		ecloudPodApplianceList(service, &cobra.Command{}, []string{"123"})
	})

	t.Run("InvalidPodID_OutputsFatal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		test_output.AssertFatalOutput(t, "Invalid pod ID [abc]\n", func() {
			ecloudPodApplianceList(service, &cobra.Command{}, []string{"abc"})
		})
	})

	t.Run("MalformedFlag_OutputsFatal", func(t *testing.T) {
		defer func() { flagFilter = nil }()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)
		cmd := &cobra.Command{}
		cmd.Flags().StringArray("filter", []string{"invalidfilter"}, "")

		test_output.AssertFatalOutput(t, "Missing value for filtering\n", func() {
			ecloudPodApplianceList(service, cmd, []string{"123"})
		})
	})

	t.Run("GetAppliancesError_OutputsFatal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetPodAppliances(123, gomock.Any()).Return([]ecloud.Appliance{}, errors.New("test error 1")).Times(1)

		test_output.AssertFatalOutput(t, "Error retrieving pod appliances: test error 1\n", func() {
			ecloudPodApplianceList(service, &cobra.Command{}, []string{"123"})
		})
	})
}
