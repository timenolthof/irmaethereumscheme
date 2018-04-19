package main

import (
	"fmt"

	"github.com/privacybydesign/irmago"
)

func main() {
	fmt.Println("Example app using IRMA (irmago)")
	//conf, err := irma.NewConfiguration("irma_configuration", "/Users/timen/Code/schememanagers")
	conf, err := irma.NewConfiguration("irma_configuration", "")
	if err != nil {
		fmt.Println(err)
	}

	// XML / Github
	// schemeId := irma.NewSchemeManagerIdentifier("pbdf")
	// err = conf.ParseFolder()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Ethereum
	schemeId := irma.NewSchemeManagerIdentifier("awesome-scheme")
	err = conf.InitAndParseEthereumContract("0xdce5f31dc8f22bfa6dc320a6a577808abb8c5bb5")
	if err != nil {
		fmt.Println(err)
	}

	// Print parsed info
	fmt.Printf("%+v\n", conf.SchemeManagers)
	fmt.Printf("%+v\n", conf.SchemeManagers[schemeId])

	// for issuerId, issuerData := range conf.Issuers {
	// 	fmt.Printf("Issuer id: %s\n", issuerId)
	// 	fmt.Printf("%+v\n", issuerData)
	// 	//fmt.Println("Key:", key, "Value:", value)
	// }
}
