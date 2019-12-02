package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/cli/test/test_output"
	"github.com/ukfast/sdk-go/pkg/service/storage"
)

func Test_storageHostList(t *testing.T) {
	t.Run("DefaultRetrieve", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		service.EXPECT().GetHosts(gomock.Any()).Return([]storage.Host{}, nil).Times(1)

		storageHostList(service, &cobra.Command{}, []string{})
	})

	t.Run("MalformedFlag_OutputsFatal", func(t *testing.T) {
		defer func() { flagFilter = nil }()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)
		flagFilter = []string{"invalidfilter"}

		test_output.AssertFatalOutput(t, "Missing value for filtering\n", func() {
			storageHostList(service, &cobra.Command{}, []string{})
		})
	})

	t.Run("GetHostsError_OutputsFatal", func(t *testing.T) {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		service.EXPECT().GetHosts(gomock.Any()).Return([]storage.Host{}, errors.New("test error")).Times(1)

		test_output.AssertFatalOutput(t, "Error retrieving hosts: test error\n", func() {
			storageHostList(service, &cobra.Command{}, []string{})
		})
	})
}

func Test_storageHostShowCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := storageHostShowCmd().Args(nil, []string{"123"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := storageHostShowCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing host", err.Error())
	})
}

func Test_storageHostShow(t *testing.T) {
	t.Run("SingleHost", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		service.EXPECT().GetHost(123).Return(storage.Host{}, nil).Times(1)

		storageHostShow(service, &cobra.Command{}, []string{"123"})
	})

	t.Run("MultipleHosts", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		gomock.InOrder(
			service.EXPECT().GetHost(123).Return(storage.Host{}, nil),
			service.EXPECT().GetHost(456).Return(storage.Host{}, nil),
		)

		storageHostShow(service, &cobra.Command{}, []string{"123", "456"})
	})

	t.Run("GetHostID_OutputsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		test_output.AssertErrorOutput(t, "Invalid host ID [abc]\n", func() {
			storageHostShow(service, &cobra.Command{}, []string{"abc"})
		})
	})

	t.Run("GetHostError_OutputsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockStorageService(mockCtrl)

		service.EXPECT().GetHost(123).Return(storage.Host{}, errors.New("test error"))

		test_output.AssertErrorOutput(t, "Error retrieving host [123]: test error\n", func() {
			storageHostShow(service, &cobra.Command{}, []string{"123"})
		})
	})
}