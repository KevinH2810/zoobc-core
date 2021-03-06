// ZooBC Copyright (C) 2020 Quasisoft Limited - Hong Kong
// This file is part of ZooBC <https://github.com/zoobc/zoobc-core>
//
// ZooBC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ZooBC is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ZooBC.  If not, see <http://www.gnu.org/licenses/>.
//
// Additional Permission Under GNU GPL Version 3 section 7.
// As the special exception permitted under Section 7b, c and e,
// in respect with the Author’s copyright, please refer to this section:
//
// 1. You are free to convey this Program according to GNU GPL Version 3,
//     as long as you respect and comply with the Author’s copyright by
//     showing in its user interface an Appropriate Notice that the derivate
//     program and its source code are “powered by ZooBC”.
//     This is an acknowledgement for the copyright holder, ZooBC,
//     as the implementation of appreciation of the exclusive right of the
//     creator and to avoid any circumvention on the rights under trademark
//     law for use of some trade names, trademarks, or service marks.
//
// 2. Complying to the GNU GPL Version 3, you may distribute
//     the program without any permission from the Author.
//     However a prior notification to the authors will be appreciated.
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Core developers that contributed to the current implementation of the
// software are:
//             Ahmad Ali Abdilah ahmad.abdilah[at]blockchainzoo.com
//             Allan Bintoro allan.bintoro[at]blockchainzoo.com
//             Andy Herman
//             Gede Sukra
//             Ketut Ariasa
//             Nawi Kartini nawi.kartini[at]blockchainzoo.com
//             Stefano Galassi stefano.galassi[at]blockchainzoo.com
//
// IMPORTANT: The above copyright notice and this permission notice
// shall be included in all copies or substantial portions of the Software.
package interceptor

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/zoobc/zoobc-core/common/crypto"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	rpcModel "github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	mockOwnerAddress = []byte{0, 0, 0, 0, 4, 38, 68, 24, 230, 247, 88, 220, 119, 124, 51, 149, 127, 214, 82, 224, 72, 239, 56, 139, 255,
		81, 229, 184, 77, 80, 80, 39, 254, 173, 28, 169}
	mockOwnerSeed = "concur vocalist rotten busload gap quote stinging undiluted surfer goofiness deviation starved"
)

type (
	mockServerStreamSuccess struct{}
	mockServerStreamNoAuth  struct {
		mockServerStreamSuccess
	}
	mockServerStreamNoMetadata struct {
		mockServerStreamSuccess
	}
	mockServerStreamInvalidAuth struct {
		mockServerStreamSuccess
	}
)

func (*mockServerStreamSuccess) SetHeader(metadata.MD) error  { return nil }
func (*mockServerStreamSuccess) SendHeader(metadata.MD) error { return nil }
func (*mockServerStreamSuccess) SetTrailer(metadata.MD)       {}
func (*mockServerStreamSuccess) Context() context.Context {
	currentTime := uint64(time.Now().Unix())
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(util.ConvertUint32ToBytes(uint32(rpcModel.RequestType_GetNodeHardware)))
	buffer.Write(util.ConvertUint64ToBytes(currentTime))
	sig, _ := (crypto.NewSignature()).Sign(
		buffer.Bytes(),
		rpcModel.AccountType_ZbcAccountType,
		mockOwnerSeed,
	)
	buffer.Write(sig)
	ctx := context.Background()
	md := metadata.Pairs("authorization", base64.StdEncoding.EncodeToString(buffer.Bytes()))
	ctx = metadata.NewIncomingContext(ctx, md)
	return ctx
}
func (*mockServerStreamSuccess) SendMsg(m interface{}) error { return nil }
func (*mockServerStreamSuccess) RecvMsg(m interface{}) error { return nil }

func (*mockServerStreamNoAuth) Context() context.Context {
	ctx := context.Background()
	md := metadata.Pairs("foo", "bar")
	ctx = metadata.NewIncomingContext(ctx, md)
	return ctx
}

func (*mockServerStreamNoMetadata) Context() context.Context {
	ctx := context.Background()
	return ctx
}

func (*mockServerStreamInvalidAuth) Context() context.Context {
	// sign with invalid request type parameter
	currentTime := uint64(time.Now().Unix())
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(util.ConvertUint32ToBytes(uint32(1435)))
	buffer.Write(util.ConvertUint64ToBytes(currentTime))
	sig, _ := (crypto.NewSignature()).Sign(
		buffer.Bytes(),
		rpcModel.AccountType_ZbcAccountType,
		mockOwnerSeed,
	)
	buffer.Write(sig)
	ctx := context.Background()
	md := metadata.Pairs("authorization", base64.StdEncoding.EncodeToString(buffer.Bytes()))
	ctx = metadata.NewIncomingContext(ctx, md)
	return ctx
}

func TestNewServerInterceptor(t *testing.T) {
	type args struct {
		logger              *logrus.Logger
		ownerAccountAddress []byte
		ignoredErrCodes     map[codes.Code]string
	}
	tests := []struct {
		name        string
		args        args
		want        grpc.UnaryServerInterceptor
		wantRecover bool
	}{
		{
			name: "wantRecover",
			args: args{
				logger: logrus.New(),
				ownerAccountAddress: []byte{0, 0, 0, 0, 185, 226, 12, 96, 140, 157, 68, 172, 119, 193, 144, 246, 76, 118, 0, 112,
					113, 140, 183, 229, 116, 202, 211, 235, 190, 224, 217, 238, 63, 223, 225, 162},
			},
			want: func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				return nil, status.Errorf(codes.Internal, "there's something wrong")
			},
			wantRecover: true,
		},
		{
			name: "wantNotRecover",
			args: args{
				logger: logrus.New(),
			},
			want: func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				return nil, status.Errorf(codes.Internal, "there's something wrong")
			},
			wantRecover: false,
		},
		{
			name: "wantNotRecover:IgnoredLog",
			args: args{
				logger: logrus.New(),
				ignoredErrCodes: map[codes.Code]string{
					codes.InvalidArgument: "invalid args",
				},
			},
			want: func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				return nil, status.Errorf(codes.InvalidArgument, "invalid args")
			},
			wantRecover: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewServerInterceptor(tt.args.logger, tt.args.ownerAccountAddress, tt.args.ignoredErrCodes)
			if cmp.Equal(got, tt.want) {
				t.Errorf("NewInterceptor() = %v, want %v", got, tt.want)
			}
			for k := range tt.args.ignoredErrCodes {
				testServerInterceptor(got, tt.wantRecover, k)
			}
		})
	}
}

func testServerInterceptor(fn grpc.UnaryServerInterceptor, wantRecover bool, errCode codes.Code) {
	var (
		handler grpc.UnaryHandler
	)
	if wantRecover {
		handler = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			panic(handler)
		}
	} else {
		handler = func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			return nil, status.Errorf(errCode, "there's something wrong")
		}
	}
	_, _ = fn(context.Background(), nil, &grpc.UnaryServerInfo{}, handler)
}

func TestNewClientInterceptor(t *testing.T) {
	type args struct {
		logger        *logrus.Logger
		ignoredErrors map[codes.Code]string
	}
	tests := []struct {
		name        string
		args        args
		want        grpc.UnaryClientInterceptor
		wantRecover bool
	}{
		{
			name: "wantNotRecover",
			args: args{
				logger:        logrus.New(),
				ignoredErrors: nil,
			},
			want: func(
				ctx context.Context,
				method string,
				req, reply interface{},
				cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker,
				opts ...grpc.CallOption) error {
				return status.Errorf(codes.Internal, "there's something wrong")
			},
			wantRecover: false,
		},
		{
			name: "wantNotRecover:ignoredLog",
			args: args{
				logger: logrus.New(),
				ignoredErrors: map[codes.Code]string{
					codes.InvalidArgument: "i want to ignored log for this err code",
				},
			},
			want: func(
				ctx context.Context,
				method string,
				req, reply interface{},
				cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker,
				opts ...grpc.CallOption) error {
				return status.Errorf(codes.Internal, "there's something wrong")
			},
			wantRecover: false,
		},
		{
			name: "wantRecover",
			args: args{
				logger: logrus.New(),
			},
			want: func(
				ctx context.Context,
				method string,
				req, reply interface{},
				cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker,
				opts ...grpc.CallOption) error {
				return status.Errorf(codes.Internal, "there's something wrong")
			},
			wantRecover: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClientInterceptor(tt.args.logger, tt.args.ignoredErrors)
			if cmp.Equal(got, tt.want) {
				t.Errorf("NewClientInterceptor() = %v, want %v", got, tt.want)
			}

			for k := range tt.args.ignoredErrors {
				testClientInterceptor(got, tt.wantRecover, k)
			}
		})
	}
}

func testClientInterceptor(fn grpc.UnaryClientInterceptor, wantRecover bool, ignoredCode codes.Code) {
	var (
		invoker grpc.UnaryInvoker
	)
	if wantRecover {
		invoker = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
			panic(invoker)
		}
	} else {
		invoker = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
			return status.Errorf(ignoredCode, "there's something wrong")
		}
	}

	cc, _ := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	_ = fn(context.Background(), "testMethod", nil, nil, cc, invoker, nil)

}

func TestNewNodeAdminAuthStreamInterceptor(t *testing.T) {

	type args struct {
		ownerAddress []byte
		fullMethod   string
		handler      grpc.StreamHandler
		serverStream grpc.ServerStream
	}
	tests := []struct {
		name         string
		args         args
		want         grpc.StreamServerInterceptor
		wantInnerErr bool
	}{
		{
			name: "wantErr:noMetadata",
			args: args{
				ownerAddress: mockOwnerAddress,
				fullMethod:   "/service.NodeHardwareService/GetNodeHardware",
				serverStream: &mockServerStreamNoMetadata{},
			},
			want: func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				return nil
			},
			wantInnerErr: true,
		},
		{
			name: "wantErr:noAuth",
			args: args{
				ownerAddress: mockOwnerAddress,
				fullMethod:   "/service.NodeHardwareService/GetNodeHardware",
				serverStream: &mockServerStreamNoAuth{},
			},
			want: func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				return nil
			},
			wantInnerErr: true,
		},
		{
			name: "wantSuccess:GetNodeHardware",
			args: args{
				ownerAddress: mockOwnerAddress,
				fullMethod:   "/service.NodeHardwareService/GetNodeHardware",
				serverStream: &mockServerStreamSuccess{},
				handler: func(srv interface{}, stream grpc.ServerStream) error {
					return nil
				},
			},
			want: func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				return nil
			},
			wantInnerErr: false,
		},
		{
			name: "wantFail:invalidSignature",
			args: args{
				ownerAddress: mockOwnerAddress,
				fullMethod:   "/service.NodeHardwareService/GetNodeHardware",
				serverStream: &mockServerStreamInvalidAuth{},
				handler: func(srv interface{}, stream grpc.ServerStream) error {
					return nil
				},
			},
			want: func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				return nil
			},
			wantInnerErr: true,
		},
		{
			name: "wantSuccess:unprotectedService",
			args: args{
				ownerAddress: mockOwnerAddress,
				fullMethod:   "/nonProtectedService",
				serverStream: &mockServerStreamSuccess{},
				handler: func(srv interface{}, stream grpc.ServerStream) error {
					return nil
				},
			},
			want: func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				return nil
			},
			wantInnerErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStreamInterceptor(tt.args.ownerAddress)
			if cmp.Equal(got, tt.want) {
				t.Errorf("NewInterceptor() = %v, want %v", got, tt.want)
			}
			info := &grpc.StreamServerInfo{
				FullMethod:     tt.args.fullMethod,
				IsClientStream: true,
				IsServerStream: true,
			}
			err := got(nil, tt.args.serverStream, info, tt.args.handler)
			if err != nil && tt.wantInnerErr != true {
				t.Error("unexpected error occurred")
			}
		})
	}
}
