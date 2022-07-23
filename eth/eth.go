// Interact with ethereum blockchain
package eth

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"strconv"

	"github.com/common-nighthawk/go-figure"
	"github.com/daoio/crypto-wizards/cards"
	"github.com/daoio/crypto-wizards/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

// instances of deployed contracts
var CWD *contracts.CWD
var CWB *contracts.CWB
var CWC *contracts.CWC

var Accounts map[int]Account
var counter int

type Account struct {
	// 20 byte eth address
	address common.Address
	// ERC721 deck
	deck contracts.CryptoWizardsDeckDeck
}

func init() {
	initAccount()
}

func initAccount() {
	greeting := figure.NewFigure("Crypto Wizards", "big", true)
	greeting.Print()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Import your private key:")
	scanner.Scan()

	// convert pasted private key string to ECDSA PrivateKey type
	pk, err := crypto.HexToECDSA(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(pk.PublicKey)
	counter++
	fmt.Println("Imported Ethereum address:", addr)

	// connect to Rinkeby testnet
	client, err := ethclient.Dial("INSERT_URL")
	if err != nil {
		log.Fatal(err)
	}
	// create new signer from imported private key
	initContracts(client)
	t, d := checkDeck(addr)
	if t {
		fmt.Printf("\nYou have deck with the folowing cardIds: %v, %v, %v, %v\n", d.Card1, d.Card2, d.Card3, d.Card4)
		cards.InitDeck(addr, d.DeckURI)
	} else {
		fmt.Println("It's required to have a deck to play CryptoWizards, do you want to mint one? [y/n]")
		scanner.Scan()
		switch scanner.Text() {
		case "y":
			auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(4))
			if err != nil {
				log.Fatal(err)
			}
			// atm that's default deck with id 1
			fmt.Println("Minting new deck...")
			CWD.CreateDeck(auth, big.NewInt(1))
			_, d = checkDeck(addr)
		case "n":
			fmt.Println("Application closed...")
		}
	}
	cards.InitDeck(addr, d.DeckURI)
}

func NewGame(deck contracts.CryptoWizardsDeckDeck) {
	/*


		TODO:::::::::::::TODO:::::::::::::::::::TODO::::::::::::::::TODO

	*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type in id of a card you want to bet: ")
	scanner.Scan()
	id := scanner.Text()

	// card ids
	c1 := strconv.Itoa(int(deck.Card1.Int64()))
	c2 := strconv.Itoa(int(deck.Card2.Int64()))
	c3 := strconv.Itoa(int(deck.Card3.Int64()))
	c4 := strconv.Itoa(int(deck.Card4.Int64()))

	switch id {
	case c1:
		fmt.Println("You've made a bet with your 1 card")
	case c2:
		fmt.Println("You've made a bet with your 2 card")
	case c3:
		fmt.Println("You've made a bet with your 3 card")
	case c4:
		fmt.Println("You've made a bet with your 4 card")
	}
}

func checkDeck(addr common.Address) (bool, contracts.CryptoWizardsDeckDeck) {
	deck, err := CWD.GetUserDeck(nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	v := reflect.ValueOf(deck)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	// skip collection id
	for i := 2; i < len(values); i++ {
		card := fmt.Sprintf("%v", values[i])
		n := new(big.Int)
		n, ok := n.SetString(card, 0)
		if !ok {
			fmt.Println("SetString: error")
		}
		c := n.Int64()
		if c != 0 {
			return true, deck
		}
	}

	return false, deck
}

func initContracts(client *ethclient.Client) {
	// already deployed addresses
	cwdAddr := common.HexToAddress("0xDEE1E025002F8AFfE084F57EC8F80771b1F60b2E")
	cwcAddr := common.HexToAddress("0x7714e72678d19b7895910fcA99e0CB9dBC9a1A03")
	cwbAddr := common.HexToAddress("0xf7Ca99115b93b85B52ddfB706aBe91cA73eC4705")
	CWD, _ = contracts.NewCWD(cwdAddr, client)
	CWC, _ = contracts.NewCWC(cwcAddr, client)
	CWB, _ = contracts.NewCWB(cwbAddr, client)
}
