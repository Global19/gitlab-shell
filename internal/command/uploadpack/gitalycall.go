package uploadpack

import (
	"context"
	"os"

	"google.golang.org/grpc"

	"gitlab.com/gitlab-org/gitaly/client"
	pb "gitlab.com/gitlab-org/gitaly/proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitlab-shell/internal/command/commandargs"
	"gitlab.com/gitlab-org/gitlab-shell/internal/gitlabnet/accessverifier"
	"gitlab.com/gitlab-org/gitlab-shell/internal/handler"
	"gitlab.com/gitlab-org/labkit/correlation"
)

func (c *Command) performGitalyCall(response *accessverifier.Response) error {
	gc := &handler.GitalyCommand{
		Config:      c.Config,
		ServiceName: string(commandargs.UploadPack),
		Address:     response.Gitaly.Address,
		Token:       response.Gitaly.Token,
		Features:    response.Gitaly.Features,
	}

	request := &pb.SSHUploadPackRequest{
		Repository:       &response.Gitaly.Repo,
		GitProtocol:      os.Getenv(commandargs.GitProtocolEnv),
		GitConfigOptions: response.GitConfigOptions,
	}

	return gc.RunGitalyCommand(func(ctx context.Context, conn *grpc.ClientConn) (int32, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		gc.LogExecution(request.Repository, response, request.GitProtocol)

		if response.CorrelationID != "" {
			ctx = correlation.ContextWithCorrelation(ctx, response.CorrelationID)
		}

		rw := c.ReadWriter
		return client.UploadPack(ctx, conn, rw.In, rw.Out, rw.ErrOut, request)
	})
}
