package component

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"opencsg.com/csghub-server/builder/deploy"
	deployCommon "opencsg.com/csghub-server/builder/deploy/common"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/git/membership"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/errorx"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/common/utils/common"
)

const spaceGitattributesContent = modelGitattributesContent

var (
	streamlitConfigContent = `[server]
enableCORS = false
enableXsrfProtection = false
`
	streamlitConfig = ".streamlit/config.toml"
)

type SpaceComponent interface {
	Create(ctx context.Context, req types.CreateSpaceReq) (*types.Space, error)
	Show(ctx context.Context, namespace, name, currentUser string, needOpWeight bool) (*types.Space, error)
	Update(ctx context.Context, req *types.UpdateSpaceReq) (*types.Space, error)
	Index(ctx context.Context, repoFilter *types.RepoFilter, per, page int, needOpWeight bool) ([]*types.Space, int, error)
	OrgSpaces(ctx context.Context, req *types.OrgSpacesReq) ([]types.Space, int, error)
	// UserSpaces get spaces of owner and visible to current user
	UserSpaces(ctx context.Context, req *types.UserSpacesReq) ([]types.Space, int, error)
	UserLikesSpaces(ctx context.Context, req *types.UserCollectionReq, userID int64) ([]types.Space, int, error)
	ListByPath(ctx context.Context, paths []string) ([]*types.Space, error)
	AllowCallApi(ctx context.Context, spaceID int64, username string) (bool, error)
	Delete(ctx context.Context, namespace, name, currentUser string) error
	Deploy(ctx context.Context, namespace, name, currentUser string) (int64, error)
	Wakeup(ctx context.Context, namespace, name string) error
	Stop(ctx context.Context, namespace, name string, deleteSpace bool) error
	// FixHasEntryFile checks whether git repo has entry point file and update space's HasAppFile property in db
	FixHasEntryFile(ctx context.Context, s *database.Space) *database.Space
	Status(ctx context.Context, namespace, name string) (string, string, error)
	Logs(ctx context.Context, namespace, name, since string) (*deploy.MultiLogReader, error)
	// HasEntryFile checks whether space repo has entry point file to run with
	HasEntryFile(ctx context.Context, space *database.Space) bool
	GetByID(ctx context.Context, spaceID int64) (*database.Space, error)
	MCPIndex(ctx context.Context, repoFilter *types.RepoFilter, per, page int) ([]*types.MCPService, int, error)
	GetMCPServiceBySvcName(ctx context.Context, svcName string) (*types.MCPService, error)
}