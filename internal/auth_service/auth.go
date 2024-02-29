package authservice

import (
	proto "github.com/Futturi/AuthSer/protos"
	"google.golang.org/grpc"
)

func InitAuth(host, port string) (proto.AuthClient, error) {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := proto.NewAuthClient(conn)
	return c, nil
}
