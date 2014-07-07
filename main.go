package main

import (
	flag "github.com/dotcloud/docker/pkg/mflag"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"os/exec"
)

type Server struct {
	Bind      string
	Container *restful.Container
}

// Initiate routing
func (s *Server) Init() {
	s.Container = restful.NewContainer()
	s.Container.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(s.envDisplay))

	s.Container.Add(ws)
}

// Start the Server
func (s *Server) Start() {
	server := &http.Server{
		Addr:    s.Bind,
		Handler: s.Container,
	}
	log.Printf("Listening on %s ...", s.Bind)
	log.Fatal(server.ListenAndServe())
}

// os.exec env command to display environment variables
// as text/plain
func (s *Server) envDisplay(request *restful.Request, response *restful.Response) {
	out, err := exec.Command("env").Output()
	if err != nil {
		log.Fatal(err)
	}
	response.AddHeader("Content-Type", "text/plain")
	response.Write(out)
}

var (
	Address string
)

// Juste add a nice flag (using docker mflag) and start it
func main() {
	flag.StringVar(&Address, []string{"l", "-listen"}, "127.0.0.1:8081", "IPaddress to listen from ex: (127.0.0.1:8081)")
	flag.Parse()
	s := &Server{
		Bind: Address,
	}
	s.Init()
	s.Start()
}
