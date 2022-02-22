#!/bin/bash

set -eu

######################################
# "SEND" | "ACK" | "TIMEOUT" | "UPDATE"
TYPE="TIMEOUT"

# "TERRA" | "SIF"
CHAIN="SIF"

CONNECTION="18" ## NOTE: 21 for Sif Terra-Sif connection; 19 for Terra Terra-Sif connection
#####################################

case $CHAIN in

  SIF)
    CLI="sifnoded"
    OUTPUT_DIR=sif
    NODE="https://rpc.sifchain.finance:443/"
    ;;

  TERRA)
    CLI="terrad"
    OUTPUT_DIR=terra
    NODE="http://public-node.terra.dev:26657"
    ;;

  *)
    echo -n "Unknown chain: $CHAIN"
    exit 1
    ;;
esac

OUTPUT_DIR=$OUTPUT_DIR/$CONNECTION


case $TYPE in

  SEND)
    QUERY="send_packet.packet_connection=connection-$CONNECTION"
    OUTPUT_DIR=$OUTPUT_DIR/send
    ;;

  ACK)
    QUERY="acknowledge_packet.packet_connection=connection-$CONNECTION"
    OUTPUT_DIR=$OUTPUT_DIR/ack
    ;;

  TIMEOUT)
    QUERY="timeout_packet.packet_connection=connection-$CONNECTION"
    OUTPUT_DIR=$OUTPUT_DIR/timeout
    ;;

  UPDATE)
    QUERY="update_client.client_id=07-tendermint-$CONNECTION"
    OUTPUT_DIR=$OUTPUT_DIR/update_client
    ;;

  *)
    echo -n "Unknown query type: $TYPE"
    exit 1
    ;;
esac



get_num_pages () {
    echo "Calculating number of pages"
    $CLI query txs --events $QUERY --node $NODE --output json  --page 1 --limit=1 > tmp
    TOTAL_COUNT=$(cat tmp | jq '.total_count|tonumber')
    NUM_PAGES=$(( ($TOTAL_COUNT + (30 - 1)) / 30))
    rm tmp
}

get_num_pages
echo "Total number of pages: $NUM_PAGES"

mkdir -p data/$OUTPUT_DIR

for ((i=1; i <= $NUM_PAGES; i++));
    do
       echo "Getting page $i"
       $CLI query txs --events $QUERY --node $NODE --page $i --output json > data/$OUTPUT_DIR/$i.json
done

