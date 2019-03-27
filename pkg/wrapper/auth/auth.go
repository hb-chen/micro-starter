package auth

import (
	"context"
	"encoding/base64"
	"strings"

	token "github.com/hb-go/micro/auth/srv/proto/token"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/vmware/vic/pkg/errors"
)

const (
	BASIC_SCHEMA  string = "Basic "
	BEARER_SCHEMA string = "Bearer "
)

type SkipperFunc func(ctx context.Context, req server.Request) bool
type TokenLookupFunc func(ctx context.Context) (string, error)

func DefaultSkipperFunc(ctx context.Context, req server.Request) bool {
	return false
}

func DefaultTokenLookupFunc(ctx context.Context) (string, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}

	// Grab the raw Authoirzation header
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BASIC_SCHEMA) && !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.New("Authorization requires Basic/Bearer scheme")
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, BASIC_SCHEMA) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(BASIC_SCHEMA):])
		if err != nil {
			return "", errors.New("Base64 encoding issue")
		}
		creds := strings.Split(string(str), ":")
		return creds[0], nil
	}

	return authHeader[len(BEARER_SCHEMA):], nil
}

// NewHandlerWrapper
func NewHandlerWrapper(service micro.Service, opts ...Option) server.HandlerWrapper {
	opt := newOptions(opts...)
	client := token.NewTokenService(opt.serviceName, service.Client())

	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if !opt.skipper(ctx, req) {
				if t, err := opt.tokenLookup(ctx); err != nil {
					return err
				} else {
					rt := &token.ReqToken{Token: t}
					if r, err := client.Verify(ctx, rt); err != nil {
						return err
					} else {
						if !r.Verified {
							return errors.New("token error")
						}
					}
				}
			}

			return h(ctx, req, rsp)
		}
	}
}
