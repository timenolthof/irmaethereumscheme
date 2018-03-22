package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"irmaethereumscheme/contracts"
	"irmaethereumscheme/proto"
	"github.com/golang/protobuf/proto"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	contractAddr := "0x9c939f2ab2c3940e7d8747a8fde73330c7442a6c";
	schemeContractAddress := common.HexToAddress(contractAddr);
	fmt.Printf("Reading scheme %s\n", contractAddr)

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Opening existing contract

	contract, err := contracts.NewIRMAScheme(schemeContractAddress, client)
	if err != nil {
		log.Fatalf("could not load contract: %v", err)
	}

	// Reading metadata (proto)buffer
	schemeMetadataBuffer, err := contract.Metadata(nil);
	if err != nil {
		log.Fatalf("could not get metadata: %v", err);
	}
	fmt.Printf("Metadata: %+v\n", schemeMetadataBuffer)
	schemeMetadata := &irmaproto.IRMASchemeMetadata{}
	proto.Unmarshal(schemeMetadataBuffer, schemeMetadata)
	fmt.Printf("Metadata: %+v\n", schemeMetadata)
}
