package auth

import (
	"context"
	ssov1 "github.com/JohnBilbo/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue = 0
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (token string, err error)
	Register(ctx context.Context, email string, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func RegisterGRPCServer(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &ServerAPI{auth: auth})
}
func (s *ServerAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing email")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing password")
	}
	if req.GetAppID() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "missing app ID")
	}
	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.AppID))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *ServerAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *ServerAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
