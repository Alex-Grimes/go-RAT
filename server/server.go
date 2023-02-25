package server

type implantServer struct {
	work, output chan *grpcapi.Command
}
