module github.com/iotexproject/high-table

go 1.12

require (
	github.com/99designs/gqlgen v0.8.3
	github.com/agnivade/levenshtein v1.0.2 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/iotexproject/go-pkgs v0.1.1
	github.com/iotexproject/iotex-core v0.10.1
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.14.3
	github.com/stretchr/testify v1.3.0
	github.com/vektah/gqlparser v1.1.2
	go.opencensus.io v0.22.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	golang.org/x/tools v0.0.0-20190815235612-5b08f89bfc0c // indirect
	google.golang.org/appengine v1.6.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/ethereum/go-ethereum => github.com/iotexproject/go-ethereum v0.2.0

exclude github.com/dgraph-io/badger v2.0.0-rc.2+incompatible

exclude github.com/dgraph-io/badger v2.0.0-rc2+incompatible

exclude github.com/ipfs/go-ds-badger v0.0.3
