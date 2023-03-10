package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "local-fog/core/types"
	"local-fog/core/utils"

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

func Discover(maxCount int) ([]*pb.NodeInfoWrapper, error) {
	// We need to buffer data because mdns.Query will send data immediately after it starts
	ch := make(chan *mdns.ServiceEntry, maxCount)
	defer close(ch)

	queryParam := mdns.DefaultParams("_localfog._tcp")
	queryParam.Entries = ch
	queryParam.DisableIPv6 = true

	errCh := make(chan error)

	go func() {
		defer func() {
			err := recover()
			e, ok := err.(error)
			if err != nil && (!ok || e.Error() != "send on closed channel") {
				panic(err)
			}
		}()

		err := mdns.Query(queryParam)
		errCh <- err
		close(errCh)
	}()

	log.Printf("start lookup")

	nodes := make([]*pb.NodeInfoWrapper, 0, maxCount)

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

			info.AddrV4 = utils.IpToUint32(entry.AddrV4)
			info.AddrV6 = entry.AddrV6

			nodes = append(nodes, info)
		}
	}

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

func (c FogConsumer) UpdateNode(req *pb.UpdateNodeRequest) (*pb.UpdateNodeReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Client.UpdateNode(ctx, req)

}
