package main

import (
	"fmt"

	bl "github.com/ViktorOshkalo/GoLearnLib/blockchain"
)

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := bl.CreateBlockchain(2)

	// record transactions on the blockchain for John, Mark, and Richard
	blockchain.AddBlock("John", "Mark", 10)
	blockchain.AddBlock("Mark", "Richard", 15)
	blockchain.AddBlock("John", "Richard", 12)
	blockchain.AddBlock("Richard", "John", 8)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.IsValid())
}
