package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/NebulousLabs/Sia/types"
)

// verifyAddress returns an address generated from a seed at the index specified
// by `index`.
func verifyAddress(uh types.UnlockHash, pks []types.SiaPublicKey, height, n uint64) bool {
	var uc types.UnlockConditions
	if height != 0 {
		uc = types.UnlockConditions{
			PublicKeys:         pks,
			SignaturesRequired: n,
			Timelock:	    types.BlockHeight(height),
		}
	} else {
		uc = types.UnlockConditions{
			PublicKeys:         pks,
			SignaturesRequired: n,
		}
	}
	return uh == uc.UnlockHash()
}

func main() {
	addressPtr := flag.String("address", "", "address to be validated")
	timelock := flag.Int("timelock", 0, "timelock block height for the addresses")
	n := flag.Int("n", 1, "signatures required")
	flag.Parse()
	keystrs := flag.Args()
	if len(keystrs) == 0 {
		log.Fatal("Must provide at least one pubkey")
		return
	}
	if *addressPtr == "" {
		log.Fatal("Must provide an address")
		return
	}

	var err error
	var uh types.UnlockHash
	err = uh.LoadString(*addressPtr)
	if err != nil {
		log.Fatal(err)
		return
	}
	var pubkeys []types.SiaPublicKey
	for _, keystr := range(keystrs) {
		var pk types.SiaPublicKey
		pk.LoadString(keystr)
		pubkeys = append(pubkeys, pk)
	}
	if verifyAddress(uh, pubkeys, uint64(*timelock), uint64(*n)) {
		fmt.Println("This address validates")
	} else {
		fmt.Println("This address does not validate")
	}
}
