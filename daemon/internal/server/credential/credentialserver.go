// Code generated by goctl. DO NOT EDIT.
// Source: credential.proto

package server

import (
	"context"

	credentiallogic "github.com/jaronnie/jzero/daemon/internal/logic/credential"
	"github.com/jaronnie/jzero/daemon/internal/pb/credentialpb"
	"github.com/jaronnie/jzero/daemon/internal/svc"
)

type CredentialServer struct {
	svcCtx *svc.ServiceContext
	credentialpb.UnimplementedCredentialServer
}

func NewCredentialServer(svcCtx *svc.ServiceContext) *CredentialServer {
	return &CredentialServer{
		svcCtx: svcCtx,
	}
}

func (s *CredentialServer) CredentialVersion(ctx context.Context, in *credentialpb.Empty) (*credentialpb.CredentialVersionResponse, error) {
	l := credentiallogic.NewCredentialVersionLogic(ctx, s.svcCtx)
	return l.CredentialVersion(in)
}

func (s *CredentialServer) CreateCredential(ctx context.Context, in *credentialpb.CreateCredentialRequest) (*credentialpb.CreateCredentialResponse, error) {
	l := credentiallogic.NewCreateCredentialLogic(ctx, s.svcCtx)
	return l.CreateCredential(in)
}
