package component

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"opencsg.com/csghub-server/builder/deploy"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/git/membership"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/errorx"
	"opencsg.com/csghub-server/common/types"
)

// 添加GetMCPServiceBySvcName的测试用例
func TestSpaceComponent_GetMCPServiceBySvcName(t *testing.T) {
	ctx := context.TODO()
	sc := initializeTestSpaceComponent(ctx, t)

	t.Run("Success", func(t *testing.T) {
		expectedSvc := &types.MCPService{
			SvcName:  "test-svc",
			Endpoint: "http://test-endpoint",
		}
		sc.mocks.stores.SpaceMock().EXPECT().GetMCPServiceBySvcName(ctx, "test-svc").Return(expectedSvc, nil)

		result, err := sc.GetMCPServiceBySvcName(ctx, "test-svc")
		require.Nil(t, err)
		require.Equal(t, expectedSvc, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		sc.mocks.stores.SpaceMock().EXPECT().GetMCPServiceBySvcName(ctx, "non-existent").Return(nil, errorx.ErrNotFound)

		result, err := sc.GetMCPServiceBySvcName(ctx, "non-existent")
		require.Nil(t, result)
		require.True(t, errors.Is(err, errorx.ErrNotFound))
	})

	t.Run("Error", func(t *testing.T) {
		expectedErr := fmt.Errorf("database error")
		sc.mocks.stores.SpaceMock().EXPECT().GetMCPServiceBySvcName(ctx, "error-svc").Return(nil, expectedErr)

		result, err := sc.GetMCPServiceBySvcName(ctx, "error-svc")
		require.Nil(t, result)
		require.ErrorIs(t, err, expectedErr)
	})
}