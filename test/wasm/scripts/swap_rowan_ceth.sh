#!/bin/bash

rawmsg=`go run ./reflect/cmd/main.go`

sifnoded tx wasm execute sif14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s62cvu6 $rawmsg \
  --from sif \
  --keyring-backend test \
  --chain-id localnet \
  --broadcast-mode block \
  -y