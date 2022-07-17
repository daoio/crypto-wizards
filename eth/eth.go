// Interact with ethereum blockchain
package eth

import (
	"bufio"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
)

var Accounts map[int]Account
var counter int

type Account struct {
	// 20 byte eth address
	Address common.Address
	// ERC721 deck
	Deck
}

// check: DeckMinter.sol
type Deck struct {
	deckId int
	deck   string
	card1  string
	card2  string
	card3  string
	card4  string
}

func init() {

}

func ImportAccount() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// convert pasted private key string to ECDSA PrivateKey type
	pk, err := crypto.HexToECDSA(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(pk.PublicKey)
	counter++
	b, deck := GetDeck()
	if b {
		Accounts[counter] = Account{
			Address: addr,
			Deck:    deck,
		}
	} else {

	}
}

// returns true + Deck if user has it, or false + default Deck if has not
func GetDeck() (bool, Deck) {

}
