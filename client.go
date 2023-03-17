package cmdb_sdk

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client interface {
	GetAddr() string
	GetDefaultTimeOut() time.Duration
}

func Dial(ctx context.Context, c Client, timeout ...time.Duration) (context.Context, context.CancelFunc, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.GetAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, nil, err
	}
	var t time.Duration
	if len(timeout) == 0 {
		t = c.GetDefaultTimeOut()
	} else {
		t = timeout[0]
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*t)
	return ctx, cancel, conn, nil
}
