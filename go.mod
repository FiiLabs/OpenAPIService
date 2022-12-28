module github.com/FiiLabs/OpenAPIService

go 1.16

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/bianjieai/iritamod-sdk-go v0.0.0-20220708032705-9e8e301da3a8
	github.com/bianjieai/opb-sdk-go v0.2.0
	github.com/containerd/fifo v1.0.0 // indirect
	github.com/docker/docker v1.4.2-0.20180625184442-8e610b2b55bf
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/irisnet/core-sdk-go v0.0.0-20220720085949-4d825adb8054
	github.com/irisnet/irismod-sdk-go/mt v0.0.0-20221014104619-6f27c71cd5e4 // indirect
	github.com/irisnet/irismod-sdk-go/nft v0.0.0-20221014104619-6f27c71cd5e4
	github.com/irisnet/irismod-sdk-go/random v0.0.0-20221014104619-6f27c71cd5e4 // indirect
	github.com/irisnet/irismod-sdk-go/record v0.0.0-20221014104619-6f27c71cd5e4 // indirect
	github.com/irisnet/irismod-sdk-go/service v0.0.0-20221014104619-6f27c71cd5e4 // indirect
	github.com/irisnet/irismod-sdk-go/token v0.0.0-20221014104619-6f27c71cd5e4 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.14.0
	go.mongodb.org/mongo-driver v1.11.0
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/net v0.3.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/prometheus/common => github.com/prometheus/common v0.26.0
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
