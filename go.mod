module github.com/ozoncp/ocp-task-api

go 1.15

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.3
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	github.com/ozoncp/ocp-task-api/pkg/ocp-task-api v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210503173754-0981d6026fa6 // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-task-api/pkg/ocp-task-api => ./pkg/ocp-task-api
