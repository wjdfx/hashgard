module github.com/hashgard/hashgard

go 1.12

require (
	bou.ke/monkey v1.0.1 // indirect
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
	github.com/gorilla/mux v0.0.0-20190228181203-15a353a63672
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
	github.com/prometheus/common v0.2.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190306233201-d0f344d83b0c // indirect
	github.com/rakyll/statik v0.1.4
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
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
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
	google.golang.org/grpc v0.0.0-20190313171052-9c3a9595696a // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect

)

replace (
	github.com/cosmos/cosmos-sdk => github.com/hashgard/cosmos-sdk v0.33.0-hashgard
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
)
