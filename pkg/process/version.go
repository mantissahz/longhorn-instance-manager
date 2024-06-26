package process

import (
	rpc "github.com/longhorn/types/pkg/generated/imrpc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/longhorn/longhorn-instance-manager/pkg/meta"
)

func (pm *Manager) VersionGet(ctx context.Context, empty *emptypb.Empty) (*rpc.VersionResponse, error) {
	v := meta.GetVersion()
	return &rpc.VersionResponse{
		Version:   v.Version,
		GitCommit: v.GitCommit,
		BuildDate: v.BuildDate,

		InstanceManagerAPIVersion:    int64(v.InstanceManagerAPIVersion),
		InstanceManagerAPIMinVersion: int64(v.InstanceManagerAPIMinVersion),

		InstanceManagerProxyAPIVersion:    int64(v.InstanceManagerProxyAPIVersion),
		InstanceManagerProxyAPIMinVersion: int64(v.InstanceManagerProxyAPIMinVersion),
	}, nil
}
