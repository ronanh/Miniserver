package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"os"
	"path"
)

type NoArgs struct{}

type Reply struct {
	Result int
}

type Meteo struct{}

func (o *Meteo) GetRainInfo(r *http.Request, args *NoArgs, reply *Reply) error {
	log.Println("GetRainInfo")
	reply.Result = 5
	return nil
}

func StaticHandler(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	if "/" == url {
		url = "/index.html"
	}
	log.Println("Static: " + url)
	current_path, _ := os.Getwd()
	http.ServeFile(rw, req, path.Join(current_path, url))
}

func main() {
	http.HandleFunc("/", StaticHandler)

	//	m := Meteo{}
	//	log.Println(m.GetRainInfo())

	// Init gorilla RPC
	pRpcServer := rpc.NewServer()
	pRpcServer.RegisterCodec(json.NewCodec(), "application/json")
	pRpcServer.RegisterService(new(Meteo), "")
	http.Handle("/rpc", pRpcServer)

	if e := http.ListenAndServe(":8080", nil); e != nil {
		log.Fatal(e)
	}
}
