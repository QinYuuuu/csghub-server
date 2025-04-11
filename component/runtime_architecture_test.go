package component

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/common/utils/common"
)

func TestRuntimeArchComponent_ListByRuntimeFrameworkID(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)

	data := []database.RuntimeArchitecture{
		{ID: 123, ArchitectureName: "arch"},
	}
	rc.mocks.stores.RuntimeArchMock().EXPECT().ListByRuntimeFrameworkID(ctx, int64(1)).Return(
		data, nil,
	)
	resp, err := rc.ListByRuntimeFrameworkID(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, data, resp)

}

func TestRuntimeArchComponent_SetArchitectures(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)

	rc.mocks.stores.RuntimeFrameworkMock().EXPECT().FindByID(ctx, int64(1)).Return(nil, nil)
	rc.mocks.stores.RuntimeArchMock().EXPECT().Add(ctx, database.RuntimeArchitecture{
		RuntimeFrameworkID: 1,
		ArchitectureName:   "foo",
	}).Return(nil)
	rc.mocks.stores.RuntimeArchMock().EXPECT().Add(ctx, database.RuntimeArchitecture{
		RuntimeFrameworkID: 1,
		ArchitectureName:   "bar",
	}).Return(errors.New(""))

	failed, err := rc.SetArchitectures(ctx, int64(1), []string{"foo", "bar"})
	require.Nil(t, err)
	require.Equal(t, []string{"bar"}, failed)

}

func TestRuntimeArchComponent_DeleteArchitectures(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)

	rc.mocks.stores.RuntimeFrameworkMock().EXPECT().FindByID(ctx, int64(1)).Return(nil, nil)
	rc.mocks.stores.RuntimeArchMock().EXPECT().DeleteByRuntimeIDAndArchName(ctx, int64(1), "foo").Return(nil)
	rc.mocks.stores.RuntimeArchMock().EXPECT().DeleteByRuntimeIDAndArchName(ctx, int64(1), "bar").Return(errors.New(""))

	failed, err := rc.DeleteArchitectures(ctx, int64(1), []string{"foo", "bar"})
	require.Nil(t, err)
	require.Equal(t, []string{"bar"}, failed)

}

func TestRuntimeArchComponent_AddRuntimeFrameworkTag(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)

	rc.mocks.stores.RuntimeFrameworkMock().EXPECT().FindByID(ctx, int64(2)).Return(
		&database.RuntimeFramework{
			FrameImage: "img",
		}, nil,
	)
	rc.mocks.stores.TagMock().EXPECT().UpsertRepoTags(ctx, int64(1), []int64{}, []int64{1}).Return(nil)

	err := rc.AddRuntimeFrameworkTag(ctx, []*database.Tag{
		{Name: "img", ID: 1},
	}, int64(1), int64(2))
	require.Nil(t, err)
}

func TestRuntimeArchComponent_AddResourceTag(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)

	rc.mocks.stores.ResourceModelMock().EXPECT().FindByModelName(ctx, "model").Return(
		[]*database.ResourceModel{
			{ResourceName: "r1"},
			{ResourceName: "r2"},
		}, nil,
	)
	rc.mocks.stores.TagMock().EXPECT().UpsertRepoTags(ctx, int64(1), []int64{}, []int64{1}).Return(nil)

	err := rc.AddResourceTag(ctx, []*database.Tag{
		{Name: "r1", ID: 1},
	}, "model", int64(1))
	require.Nil(t, err)
}

func TestRuntimeArchComponent_GetGGUFContent(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRuntimeArchComponent(ctx, t)
	rc.fileDownloadPath = "https://hub.opencsg.com/csg"
	req := types.GetFileReq{
		Lfs:       true,
		Namespace: "AIWizards",
		Name:      "Llama-2-7B-GGUF",
		Path:      "llama-2-7b.Q2_K.gguf",
		Ref:       "main",
		RepoType:  types.ModelRepo,
	}
	rc.mocks.components.repo.EXPECT().InternalDownloadFile(ctx, &req).Return(
		nil, 0, "https://hub.opencsg.com/csg/AIWizards/Llama-2-7B-GGUF/resolve/main/llama-2-7b.Q2_K.gguf", nil,
	)
	file, err := rc.GetGGUFContent(ctx, "llama-2-7b.Q2_K.gguf", &database.Repository{
		Path:          "AIWizards/Llama-2-7B-GGUF",
		DefaultBranch: "main",
	})
	require.Nil(t, err)
	meta := file.Metadata()
	require.Equal(t, "llama", meta.Architecture)
}

func TestRuntimeArchComponent_GetSafetensorsContent(t *testing.T) {
	fileList := []string{}
	//fileList append from 00001 to model-00001-of-00004.safetensors
	for i := 1; i <= 4; i++ {
		fileList = append(fileList, fmt.Sprintf("https://hub.opencsg.com/csg/Qwen/Qwen2.5-7B-Instruct/resolve/main/model-%05d-of-00004.safetensors", i))
	}
	modelInfo, err := common.GetModelInfo(fileList, "", 5120)
	modelInfo.HiddenSize = 3584
	modelInfo.NumHiddenLayers = 28
	modelInfo.NumAttentionHeads = 28
	kvcacheSize := common.GetKvCacheSize(modelInfo.ContextSize, modelInfo.BatchSize, modelInfo.HiddenSize, modelInfo.NumHiddenLayers, modelInfo.BytesPerParam)
	activateMemory := common.GetActivationMemory(modelInfo.BatchSize, modelInfo.ContextSize, modelInfo.NumHiddenLayers, modelInfo.HiddenSize, modelInfo.NumAttentionHeads, modelInfo.BytesPerParam)
	modelInfo.MiniGPUMemoryGB = kvcacheSize + modelInfo.ModelWeightsGB + activateMemory
	require.Nil(t, err)
	require.Equal(t, "BF16", modelInfo.TensorType)
	require.Equal(t, float32(7.62), modelInfo.ParamsBillions)
	require.Equal(t, 22, int(modelInfo.MiniGPUMemoryGB))
}
