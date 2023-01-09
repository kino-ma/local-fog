package core

import (
	"context"
	"fmt"
	"time"

	pb "local-fog/core/types"

	"google.golang.org/grpc"
)

type FogConsumer struct {
	Conn   *grpc.ClientConn
	Client pb.LocalFogClient
}

func Connect(host string, port int) (FogConsumer, error) {
	target := host + ":" + fmt.Sprint(port)
	conn, err := grpc.Dial(target)

	if err != nil {
		return FogConsumer{}, err
	}

	client := pb.NewLocalFogClient(conn)

	return FogConsumer{
		Conn:   conn,
		Client: client,
	}, nil
}

func (c FogConsumer) Ping(req *pb.PingRequest) (*pb.PingReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Client.Ping(ctx, req)
}

func (c FogConsumer) Sync(req *pb.SyncRequest) (*pb.SyncReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Client.Sync(ctx, req)
}

func (c FogConsumer) Call(req *pb.CallRequest) (*pb.CallReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Client.Call(ctx, req)

}

func (c FogConsumer) GetProgram(req *pb.GetProgramRequest) (*pb.GetProgramReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Client.GetProgram(ctx, req)

}
