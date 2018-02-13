package main

import (
	"fmt"
	"github.com/privacybydesign/irmago"
)

func main() {
	fmt.Println("Hello")
	//conf, err := irma.NewConfiguration("irma_configuration", "/Users/timen/Code/schememanagers")
	conf, err := irma.NewConfiguration("irma_configuration", "");
	//if err != nil {
	//	fmt.Println(err)
	//}

	err = conf.ParseFolder()
	if err != nil {
		fmt.Println(err)
	}

	for issuerId, issuerData := range conf.Issuers {
		fmt.Printf("Issuer id: %s\n", issuerId)
		fmt.Printf("%+v\n", issuerData)
		//fmt.Println("Key:", key, "Value:", value)
	}
}
