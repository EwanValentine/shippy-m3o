module vessels

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-micro/v3 v3.0.0-beta.0.20200825081046-bf8b3aeac796
	github.com/micro/micro/v3 v3.0.0-beta.3
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	google.golang.org/protobuf v1.25.0
)
