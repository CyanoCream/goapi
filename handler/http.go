package handler

import "sesi_8/service"

type HttpServer struct {
	app service.ServiceInterface
}

func NewHttpServer(app service.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
