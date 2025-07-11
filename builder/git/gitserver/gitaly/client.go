package gitaly

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	gitalyauth "gitlab.com/gitlab-org/gitaly/v16/auth"
	gitalyclient "gitlab.com/gitlab-org/gitaly/v16/client"
	gitalypb "gitlab.com/gitlab-org/gitaly/v16/proto/go/gitalypb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/config"
)

var _ gitserver.GitServer = (*Client)(nil)

type Client struct {
	config              *config.Config
	sidechannelRegistry *gitalyclient.SidechannelRegistry
	repoClient          gitalypb.RepositoryServiceClient
	commitClient        gitalypb.CommitServiceClient
	blobClient          gitalypb.BlobServiceClient
	refClient           gitalypb.RefServiceClient
	diffClient          gitalypb.DiffServiceClient
	operationClient     gitalypb.OperationServiceClient
	smartHttpClient     gitalypb.SmartHTTPServiceClient
	remoteClient        gitalypb.RemoteServiceClient
	timeout             time.Duration
	repoStore           database.RepoStore
}

func NewClient(config *config.Config) (*Client, error) {
	var sidechannelRegistry *gitalyclient.SidechannelRegistry
	accessLogger := log.New()
	accessLogger.SetLevel(log.InfoLevel)
	sidechannelRegistry = gitalyclient.NewSidechannelRegistry(log.NewEntry(accessLogger))
	connOpts := append(gitalyclient.DefaultDialOpts,
		grpc.WithPerRPCCredentials(gitalyauth.RPCCredentialsV2(config.GitalyServer.Token)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		gitalyclient.WithGitalyDNSResolver(gitalyclient.DefaultDNSResolverBuilderConfig()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)

	conn, connErr := gitalyclient.DialSidechannel(context.Background(), config.GitalyServer.Address, sidechannelRegistry, connOpts)
	repoClient := gitalypb.NewRepositoryServiceClient(conn)
	commitClient := gitalypb.NewCommitServiceClient(conn)
	blobClient := gitalypb.NewBlobServiceClient(conn)
	refClient := gitalypb.NewRefServiceClient(conn)
	diffClient := gitalypb.NewDiffServiceClient(conn)
	operationClient := gitalypb.NewOperationServiceClient(conn)
	smartHttpClient := gitalypb.NewSmartHTTPServiceClient(conn)
	remoteClient := gitalypb.NewRemoteServiceClient(conn)

	if connErr != nil {
		return nil, connErr
	}
	timeoutTime := time.Duration(config.Git.OperationTimeout) * time.Second

	return &Client{
		config:              config,
		sidechannelRegistry: sidechannelRegistry,
		repoClient:          repoClient,
		commitClient:        commitClient,
		blobClient:          blobClient,
		refClient:           refClient,
		diffClient:          diffClient,
		operationClient:     operationClient,
		smartHttpClient:     smartHttpClient,
		remoteClient:        remoteClient,
		timeout:             timeoutTime,
		repoStore:           database.NewRepoStore(),
	}, nil
}
