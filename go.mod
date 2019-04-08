module github.com/hashgard/hashgard

go 1.12

require (
	bou.ke/monkey v1.0.1 // indirect
	cloud.google.com/go v0.34.0 // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d // indirect
	github.com/btcsuite/btcutil v0.0.0-20190207003914-4c204d697803 // indirect
	github.com/cosmos/cosmos-sdk v0.33.0
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8 // indirect
	github.com/cosmos/ledger-cosmos-go v0.9.8 // indirect
	github.com/ethereum/go-ethereum v0.0.0-20190313125352-1a29bf0ee2c5 // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/websocket v0.0.0-20190306004257-0ec3d1bd7fe5 // indirect
	github.com/hashicorp/hcl v0.0.0-20180906183839-65a6292f0157 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/magiconair/properties v0.0.0-20190110142458-7757cc9fdb85 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/otiai10/copy v0.0.0-20180813032824-7e9a647135a1
	github.com/otiai10/curr v0.0.0-20150429015615-9b4961190c95 // indirect
	github.com/otiai10/mint v1.2.3 // indirect
	github.com/pelletier/go-toml v0.0.0-20190313045714-690ec00a4b7e // indirect
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/procfs v0.0.0-20190306233201-d0f344d83b0c // indirect
	github.com/rakyll/statik v0.1.4
	github.com/rs/cors v0.0.0-20190116175910-76f58f330d76 // indirect
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v0.0.0-20181223182923-24fa6976df40 // indirect
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20180708030551-c4c61651e9e3 // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/iavl v0.12.0 // indirect
	github.com/tendermint/tendermint v0.31.0-dev0
	github.com/zondax/hid v0.9.0 // indirect
	github.com/zondax/ledger-go v0.8.0 // indirect
	golang.org/x/crypto v0.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect

)

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.37.2
	git.apache.org/thrift => github.com/apache/thrift v0.12.0
	github.com/cosmos/cosmos-sdk => github.com/hashgard/cosmos-sdk v0.33.0-hashgard
	golang.org/x/build => github.com/golang/build v0.0.0-20190402050623-31fad79bef7f
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190321205749-f0864edee7f3
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190328230028-74de082e2cca
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190319182350-c85d3e98c914
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190312170614-0655857e383f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190402054613-e4093980e83e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190401205534-4c644d7e323d
	google.golang.org/api => github.com/golang/go v0.0.0-20190402054533-56517216c052
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190401181712-f467c93bbac2
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1

)
