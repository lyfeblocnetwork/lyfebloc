#!/bin/bash

KEY="dev0"
CHAINID="lyfebloc_1775-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t lyfebloc-datadir.XXXXX)

echo "create and add new keys"
./lyfeblocd keys add $KEY --home "$DATA_DIR" --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Lyfebloc with moniker=$MONIKER and chain-id=$CHAINID"
./lyfeblocd init $MONIKER --chain-id "$CHAINID" --home "$DATA_DIR"
echo "prepare genesis: Allocate genesis accounts"
./lyfeblocd add-genesis-account \
	"$(./lyfeblocd keys show "$KEY" -a --home "$DATA_DIR" --keyring-backend test)" 1000000000000000000alyfe,1000000000000000000stake \
	--home "$DATA_DIR" --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./lyfeblocd gentx "$KEY" 1000000000000000000stake --keyring-backend test --home "$DATA_DIR" --keyring-backend test --chain-id "$CHAINID"
echo "prepare genesis: Collect genesis tx"
./lyfeblocd collect-gentxs --home "$DATA_DIR"
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./lyfeblocd validate-genesis --home "$DATA_DIR"

echo "starting lyfebloc node in background ..."
./lyfeblocd start --pruning=nothing --rpc.unsafe \
	--keyring-backend test --home "$DATA_DIR" \
	>"$DATA_DIR"/node.log 2>&1 &
disown

echo "started lyfebloc node"
tail -f /dev/null
