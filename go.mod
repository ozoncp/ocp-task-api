module github.com/ozoncp/ocp-task-api

go 1.15

require (
	github.com/gogo/status v1.1.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.3
	github.com/klim0v/grpc-gateway-example v0.0.0-20200421090150-0b0c1702054d
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	github.com/ozoncp/ocp-task-api/pkg/ocp-task-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.10.0
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-task-api/pkg/ocp-task-api => ./pkg/ocp-task-api
