package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"log"
	"irmaethereumscheme/contracts"
	"irmaethereumscheme/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Creating scheme")

	const key = `{
      "address": "023e291a99d21c944a871adcc44561a58f99bdbc",
      "crypto": {
          "cipher": "aes-128-ctr",
          "cipherparams": {
              "iv": "57ef5312b9eb89d4a5d7f8cef6131987"
          },
          "ciphertext": "963839ae10dda6929a682c7e818b8c2af0c39f0eaeff7f1764a7d197e4a472a6",
          "kdf": "scrypt",
          "kdfparams": {
              "dklen": 32,
              "n": 262144,
              "p": 1,
              "r": 8,
              "salt": "836823419b773fcb0f8708f6129c9a7ce4cc49d1b3997be604b1679542fd03f7"
          },
          "mac": "7d3d92f46bf39729dea5060421f98143c79b23f42a7e5a35793956c614985413"
      },
      "id": "e04ebdfa-a96a-4581-bb4d-fb8bd54837ae",
      "version": 3
  }`

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "linux")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

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

	// addr, _, contract, err := DeployGreeter(auth, sim, "DIVA is cool")
	addr, _, _, err := contracts.DeployIRMAScheme(auth, client, "pbdf", metadataBuffer)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	fmt.Printf("Contract deployed to %s\n", addr.String())
}