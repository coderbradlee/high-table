module github.com/iotexproject/high-table

go 1.13

require (
	github.com/golang/protobuf v1.3.1
	github.com/iotexproject/iotex-core v0.10.1
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/pkg/errors v0.8.1
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/grpc v1.21.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/ethereum/go-ethereum => github.com/iotexproject/go-ethereum v0.2.0

exclude github.com/dgraph-io/badger v2.0.0-rc.2+incompatible

exclude github.com/dgraph-io/badger v2.0.0-rc2+incompatible

exclude github.com/ipfs/go-ds-badger v0.0.3
