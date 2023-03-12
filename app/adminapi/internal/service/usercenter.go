package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-admin/api/adminapi/service/v1"
)

func (s *AdminApiService) Login(ctx context.Context, req *v1.LoginReq) (resp *v1.LoginResp, err error) {
	resp, err = s.userCenterRepo.Login(ctx, req)
	return
}
func (s *AdminApiService) Logout(ctx context.Context, req *emptypb.Empty) (resp *emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (s *AdminApiService) Captcha(ctx context.Context, req *emptypb.Empty) (resp *emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method Captcha not implemented")
}
