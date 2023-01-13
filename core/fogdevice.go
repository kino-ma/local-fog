package core

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func Discover(maxCount int) ([]*pb.NodeInfo, error) {
	// We need to buffer data because mdns.Query will send data immediately after it starts
	ch := make(chan *mdns.ServiceEntry, maxCount)

	queryParam := mdns.DefaultParams("_localfog._tcp")
	queryParam.Entries = ch
	queryParam.DisableIPv6 = true

	errCh := make(chan error)

	go func() {
		err := mdns.Query(queryParam)
		errCh <- err
	}()

	log.Printf("start lookup")

	nodes := make([]*pb.NodeInfo, 0, maxCount)

	for i := 0; i < maxCount; i++ {
		select {
		case err := <-errCh:
			return nodes, err
		case entry := <-ch:
			log.Printf("got entry: %v", entry)

			info, err := ParseTxt(entry.Info)
			if err != nil {
				if errors.Is(err, ErrNotLocalFogService) {
					continue
				} else {
					return nil, err
				}
			}
			if info.Id == 0 {
				log.Printf("Invalid record: %v", entry)
				continue
			}

			info.AddrV4 = IpToUint32(entry.AddrV4)
			info.AddrV6 = entry.AddrV6

			nodes = append(nodes, info)
		}
	}
	close(ch)

	return nodes, nil
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
