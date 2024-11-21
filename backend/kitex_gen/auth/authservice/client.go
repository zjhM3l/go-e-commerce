// Code generated by Kitex v0.11.3. DO NOT EDIT.

package authservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	auth "github.com/zjhM3l/go-e-commerce/kitex_gen/auth"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DeliverTokenByRPC(ctx context.Context, request *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliverTokenResp, err error)
	VerifyTokenByRPC(ctx context.Context, request *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kAuthServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAuthServiceClient struct {
	*kClient
}

func (p *kAuthServiceClient) DeliverTokenByRPC(ctx context.Context, request *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliverTokenResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeliverTokenByRPC(ctx, request)
}

func (p *kAuthServiceClient) VerifyTokenByRPC(ctx context.Context, request *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyTokenByRPC(ctx, request)
}
