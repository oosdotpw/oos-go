package lib

import (
	"net/http"
	"oos-go/utils"
	"reflect"
)

type HandlerMap map[string]reflect.Type

type HttpServer struct {
	Maps HandlerMap
}

func (h HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	utils.Log(utils.INF, r.URL)

	typ, exist := h.Maps[r.URL.Path]
	if !exist {
		http.Error(w, "Page Not Found", 404)
		return
	}

	rHandler := reflect.New(typ).Elem()

	rHandler.FieldByName("Writer").Set(reflect.ValueOf(&w))
	rHandler.FieldByName("Request").Set(reflect.ValueOf(r))

	rHandler.MethodByName("Post").Call([]reflect.Value{})
}

func NewHttpServer(maps map[string]HandlerInterface) HttpServer {
	newMap := make(HandlerMap)
	for key, value := range maps {
		newMap[key] = reflect.TypeOf(value)
	}

	server := HttpServer{Maps: newMap}
	return server
}
