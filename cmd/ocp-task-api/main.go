package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/ozoncp/ocp-task-api/internal/api"
	desc "github.com/ozoncp/ocp-task-api/pkg/ocp-task-api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
	httpPort = ":7000"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
	httpEndpoint = flag.String("http-server-endpoint", "0.0.0.0"+httpPort, "HTTP server endpoint")

	swaggerDir = flag.String("swagger-dir", "swagger", "path to the directory which contains swagger definitions")
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Swagger JSON not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}

func runHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpTaskApiHandlerFromEndpoint(ctx, gwmux, *grpcEndpoint, opts)
	if err != nil {
		return err
	}

	mux.HandleFunc("/swagger/", serveSwagger)
	mux.Handle("/", gwmux)

	fmt.Printf("Server listening on %s\n", *httpEndpoint)
	return http.ListenAndServe(*httpEndpoint, mux)
}

func runGrpc() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpTaskApiServer(s, api.NewOcpTaskApi())

	fmt.Printf("Server listening on %s\n", *grpcEndpoint)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()

	go runGrpc()

	if err := runHttp(); err != nil {
		log.Fatal(err)
	}
}
