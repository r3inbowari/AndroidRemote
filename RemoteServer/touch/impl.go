package touch

import "time"

type Server struct {
	UnimplementedTouchServer
}

func (s Server) TouchReq(request *TouchRequest, server Touch_TouchReqServer) error {
	if request.Message == "attach touch" {
		// ctx := server.Context()
		server.Send(&TouchReply{
			Type:          0,
			X:             110,
			Y:             110,
			Contact:       0,
			Ts:            10,
		})

		time.Sleep(time.Millisecond *20)
		server.Send(&TouchReply{
			Type:          TouchReply_RELEASE,
			Contact:       0,
			Ts:            10,
		})
	}
	return nil
}

func (s Server) mustEmbedUnimplementedTouchServer() {
	// panic("implement me")
}
