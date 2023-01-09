package core

import (
	"fmt"

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
