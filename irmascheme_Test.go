package main

import (
	"fmt"
	"log"
	"irmaethereumscheme/contracts"
	"irmaethereumscheme/proto"
	"github.com/golang/protobuf/proto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// Generate keys
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)




	// Create simulator with prefunded account for generated keys
	alloc := core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}}
	alloc[crypto.PubkeyToAddress(key.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(10000000000)}
	sim := backends.NewSimulatedBackend(alloc)




	// Deploy contract including metadata
	metadata := &irmaproto.IRMASchemeMetadata{
		Version: 7,
		Id: "pbdf",
		Url: "https://privacybydesign.foundation/schememanager/pbdf",
		Name: []*irmaproto.IRMASchemeMetadata_LocalizedName{
			{ Lang: "en", Name: "Privacy by Design Foundation" },
			{ Lang: "nl", Name: "Stichting Privacy by Design" },
		},
		Description: []*irmaproto.IRMASchemeMetadata_LocalizedDescription{
			{
				Lang: "en",
				Name: "The Privacy by Design Foundation develops the IRMA app and the IRMA " +
					"infrastructure, and issues basic attributes for free.",
			},{
				Lang: "nl",
				Name: "De stichting Privacy by Design ontwikkelt de IRMA app en de IRMA" +
					"infrastructuur, en geeft gratis een set basisattributen uit.",
			},
		},
		Keyshareserver: "https://privacybydesign.foundation/tomcat/irma_keyshare_server/api/v1",
		Keysharewebsite: "https://privacybydesign.foundation/mijnirma/",
		Keyshareattribute: "pbdf.pbdf.mijnirma.email",
		Contact: "https://privacybydesign.foundation/",
	}
	metadataBuffer, err := proto.Marshal(metadata);
	fmt.Printf("Size of metadata: %d, %d\n", len(metadataBuffer), cap(metadataBuffer));

	addr, createTransaction, createdContract, err := contracts.DeployIRMAScheme(auth, sim, "pbdf", metadataBuffer)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	fmt.Printf("Contract deployed to %s\n", addr.String())
	fmt.Printf("Transaction cost %d gas\n", createTransaction.Gas())
	sim.Commit()
	fmt.Printf("Created contract: %+v\n\n", createdContract.IRMASchemeCaller)


	// Opening the new contract
	contractAddr := addr.String();
	schemeContractAddress := common.HexToAddress(contractAddr);
	fmt.Printf("Reading scheme from: %s\n", contractAddr)
	contract, err := contracts.NewIRMAScheme(schemeContractAddress, sim)
	if err != nil {
		log.Fatalf("could not load contract: %v", err)
	}
	fmt.Printf("Contract: %+v\n", contract)



	// Reading metadata (proto)buffer
	schemeMetadataBuffer, err := contract.Metadata(nil);
	if err != nil {
		log.Fatalf("could not get metadata: %v", err);
	}
	schemeMetadata := &irmaproto.IRMASchemeMetadata{}
	proto.Unmarshal(schemeMetadataBuffer, schemeMetadata)
	fmt.Printf("Metadata: %+v\n", schemeMetadata)



	// Add Issuer
	issuerMetadata := &irmaproto.IRMAIssuerMetadata{
		Version: 4,
		Id: "pbdf",
		Shortname: []*irmaproto.IRMAIssuerMetadata_LocalizedName{
			{ Lang: "en", Name: "PbD.f" },
			{ Lang: "nl", Name: "PbD.f" },
		},
		Name: []*irmaproto.IRMAIssuerMetadata_LocalizedName{
			{ Lang: "en", Name: "Privacy by Design Foundation" },
			{ Lang: "nl", Name: "Stichting Privacy by Design" },
		},
		Schememanager: "pbdf",
		Contactaddress: "",
		Contactemail: "info@privacybydesign.foundation",
		Baseurl: "",
	}
	issuerMetadataBuffer, err := proto.Marshal(issuerMetadata);
	fmt.Printf("Size of metadata: %d\n", len(issuerMetadataBuffer));

	transaction, err := contract.AddIssuer(auth, "pbdf", "logo.png", issuerMetadataBuffer);
	if err != nil {
		log.Fatalf("could not add issuer to contract: %v", err)
	}
	fmt.Printf("Transaction cost %d gas\n", transaction.Gas())
	sim.Commit();




}