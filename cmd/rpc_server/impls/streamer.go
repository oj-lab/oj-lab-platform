package impls

import (
	"log"
	"time"

	"github.com/oj-lab/platform/proto"
)

type StreamerServer struct {
	proto.UnimplementedStreamerServer
}

func (s *StreamerServer) StartStream(request *proto.StreamRequest, server proto.Streamer_StartStreamServer) error {
	tick := time.NewTicker(1 * time.Second)
	for range tick.C {
		if server.Context().Err() != nil {
			if server.Context().Err().Error() == "context canceled" {
				log.Printf("client closed stream")
				return nil
			}
			log.Printf("client closed stream with: %v", server.Context().Err().Error())
			return nil
		}

		err := server.Send(&proto.StreamResponse{Body: &proto.StreamResponse_Health{
			Health: &proto.ServerHealth{},
		}})
		if err != nil {
			log.Printf("send stream response failed: %v", err)
		}
	}

	return nil
}
