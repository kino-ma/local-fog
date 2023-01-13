package core

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "local-fog/core/types"

	"github.com/hashicorp/mdns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FogConsumer struct {
	Conn   *grpc.ClientConn
	Client pb.LocalFogClient
}

func Connect(host string, port int) (FogConsumer, error) {
	target := host + ":" + fmt.Sprint(port)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return FogConsumer{}, err
	}

	client := pb.NewLocalFogClient(conn)

	return FogConsumer{
		Conn:   conn,
		Client: client,
	}, nil
}

func Discover() (net.IP, error) {
	// We need to buffer data because mdns.Query will send data immediately after it starts
	ch := make(chan *mdns.ServiceEntry, 1)

	queryParam := mdns.DefaultParams("_localfog._tcp")
	queryParam.Entries = ch
	queryParam.DisableIPv6 = true

	err := mdns.Query(queryParam)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup the service: %v", err)
	}

	log.Printf("start lookup")

	entry := <-ch
	log.Printf("got entry: %v", entry)
	close(ch)

	return entry.AddrV4, nil
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
