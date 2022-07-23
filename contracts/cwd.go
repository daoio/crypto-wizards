// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CryptoWizardsDeckDeck is an auto generated low-level Go binding around an user-defined struct.
type CryptoWizardsDeckDeck struct {
	CollectionId *big.Int
	DeckURI      string
	Card1        *big.Int
	Card2        *big.Int
	Card3        *big.Int
	Card4        *big.Int
}

// CWDMetaData contains all meta data concerning the CWD contract.
var CWDMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cwc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"}],\"name\":\"createDeck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cryptoWizardsCard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deckURIs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"decks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"deckURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"card1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card4\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserDeck\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"deckURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"card1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"card4\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoWizardsDeck.Deck\",\"name\":\"deck\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasDeck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"deckURI\",\"type\":\"string\"}],\"name\":\"setDeckURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_select\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"updateDeck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001e5e38038062001e5e8339818101604052810190620000379190620001d5565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000207565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b620001af8162000190565b8114620001bb57600080fd5b50565b600081519050620001cf81620001a4565b92915050565b600060208284031215620001ee57620001ed6200016b565b5b6000620001fe84828501620001be565b91505092915050565b611c4780620002176000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063822b6eeb11610071578063822b6eeb146101435780638da5cb5b14610173578063b6781b7814610191578063d4f2df3e146101ad578063da583940146101dd578063f2fde38b1461020d576100a9565b80631c86871d146100ae5780632c9fad06146100ca5780632ed21357146100ff5780634be5154d1461011d578063715018a614610139575b600080fd5b6100c860048036038101906100c391906110ed565b610229565b005b6100e460048036038101906100df9190611140565b6102e0565b6040516100f696959493929190611215565b60405180910390f35b6101076103a4565b604051610114919061128c565b60405180910390f35b610137600480360381019061013291906113dc565b6103ca565b005b6101416103f7565b005b61015d60048036038101906101589190611140565b61040b565b60405161016a919061151a565b60405180910390f35b61017b610528565b604051610188919061128c565b60405180910390f35b6101ab60048036038101906101a6919061153c565b610551565b005b6101c760048036038101906101c2919061153c565b610896565b6040516101d49190611569565b60405180910390f35b6101f760048036038101906101f29190611140565b610936565b60405161020491906115a6565b60405180910390f35b61022760048036038101906102229190611140565b610956565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b090611633565b60405180910390fd5b600082036102d0576102cb83826109d9565b6102db565b6102da8382610c4f565b5b505050565b600260205280600052604060002060009150905080600001549080600101805461030990611682565b80601f016020809104026020016040519081016040528092919081815260200182805461033590611682565b80156103825780601f1061035757610100808354040283529160200191610382565b820191906000526020600020905b81548152906001019060200180831161036557829003601f168201915b5050505050908060020154908060030154908060040154908060050154905086565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103d2610ec5565b806003600084815260200190815260200160002090816103f2919061185f565b505050565b6103ff610ec5565b6104096000610f43565b565b61041361100f565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c00160405290816000820154815260200160018201805461047790611682565b80601f01602080910402602001604051908101604052809291908181526020018280546104a390611682565b80156104f05780601f106104c5576101008083540402835291602001916104f0565b820191906000526020600020905b8154815290600101906020018083116104d357829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815250509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156105de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d5906119a3565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ea95180d83336040518363ffffffff1660e01b815260040161063d9291906119c3565b6080604051808303816000875af115801561065c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106809190611ab7565b905060006040518060c001604052808481526020016003600086815260200190815260200160002080546106b390611682565b80601f01602080910402602001604051908101604052809291908181526020018280546106df90611682565b801561072c5780601f106107015761010080835404028352916020019161072c565b820191906000526020600020905b81548152906001019060200180831161070f57829003601f168201915b505050505081526020018360006004811061074a57610749611ae4565b5b602002015181526020018360016004811061076857610767611ae4565b5b602002015181526020018360026004811061078657610785611ae4565b5b60200201518152602001836003600481106107a4576107a3611ae4565b5b6020020151815250905080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101908161080d919061185f565b5060408201518160020155606082015181600301556080820151816004015560a082015181600501559050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505050565b600360205280600052604060002060009150905080546108b590611682565b80601f01602080910402602001604051908101604052809291908181526020018280546108e190611682565b801561092e5780601f106109035761010080835404028352916020019161092e565b820191906000526020600020905b81548152906001019060200180831161091157829003601f168201915b505050505081565b60046020528060005260406000206000915054906101000a900460ff1681565b61095e610ec5565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c490611b85565b60405180910390fd5b6109d681610f43565b50565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182018054610a3f90611682565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6b90611682565b8015610ab85780601f10610a8d57610100808354040283529160200191610ab8565b820191906000526020600020905b815481529060010190602001808311610a9b57829003601f168201915b5050505050815260200160028201548152602001600382015481526020016004820154815260200160058201548152505090506000816040015103610b435781600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550610c4a565b6000816060015103610b9b5781600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550610c49565b6000816080015103610bf35781600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040181905550610c48565b60008160a0015103610c475781600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600501819055505b5b5b5b505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182018054610cb590611682565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce190611682565b8015610d2e5780601f10610d0357610100808354040283529160200191610d2e565b820191906000526020600020905b815481529060010190602001808311610d1157829003601f168201915b50505050508152602001600282015481526020016003820154815260200160048201548152602001600582015481525050905081816040015103610db9576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550610ec0565b81816060015103610e11576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550610ebf565b81816080015103610e69576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040181905550610ebe565b818160a0015103610ebd576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600501819055505b5b5b5b505050565b610ecd611007565b73ffffffffffffffffffffffffffffffffffffffff16610eeb610528565b73ffffffffffffffffffffffffffffffffffffffff1614610f41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3890611bf1565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060c001604052806000815260200160608152602001600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061108482611059565b9050919050565b61109481611079565b811461109f57600080fd5b50565b6000813590506110b18161108b565b92915050565b6000819050919050565b6110ca816110b7565b81146110d557600080fd5b50565b6000813590506110e7816110c1565b92915050565b6000806000606084860312156111065761110561104f565b5b6000611114868287016110a2565b9350506020611125868287016110d8565b9250506040611136868287016110d8565b9150509250925092565b6000602082840312156111565761115561104f565b5b6000611164848285016110a2565b91505092915050565b611176816110b7565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111b657808201518184015260208101905061119b565b838111156111c5576000848401525b50505050565b6000601f19601f8301169050919050565b60006111e78261117c565b6111f18185611187565b9350611201818560208601611198565b61120a816111cb565b840191505092915050565b600060c08201905061122a600083018961116d565b818103602083015261123c81886111dc565b905061124b604083018761116d565b611258606083018661116d565b611265608083018561116d565b61127260a083018461116d565b979650505050505050565b61128681611079565b82525050565b60006020820190506112a1600083018461127d565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6112e9826111cb565b810181811067ffffffffffffffff82111715611308576113076112b1565b5b80604052505050565b600061131b611045565b905061132782826112e0565b919050565b600067ffffffffffffffff821115611347576113466112b1565b5b611350826111cb565b9050602081019050919050565b82818337600083830152505050565b600061137f61137a8461132c565b611311565b90508281526020810184848401111561139b5761139a6112ac565b5b6113a684828561135d565b509392505050565b600082601f8301126113c3576113c26112a7565b5b81356113d384826020860161136c565b91505092915050565b600080604083850312156113f3576113f261104f565b5b6000611401858286016110d8565b925050602083013567ffffffffffffffff81111561142257611421611054565b5b61142e858286016113ae565b9150509250929050565b611441816110b7565b82525050565b600082825260208201905092915050565b60006114638261117c565b61146d8185611447565b935061147d818560208601611198565b611486816111cb565b840191505092915050565b600060c0830160008301516114a96000860182611438565b50602083015184820360208601526114c18282611458565b91505060408301516114d66040860182611438565b5060608301516114e96060860182611438565b5060808301516114fc6080860182611438565b5060a083015161150f60a0860182611438565b508091505092915050565b600060208201905081810360008301526115348184611491565b905092915050565b6000602082840312156115525761155161104f565b5b6000611560848285016110d8565b91505092915050565b6000602082019050818103600083015261158381846111dc565b905092915050565b60008115159050919050565b6115a08161158b565b82525050565b60006020820190506115bb6000830184611597565b92915050565b7f43727970746f57697a617264734465636b7b6f6e6c794357437d3a206163636560008201527f73732064656e6965640000000000000000000000000000000000000000000000602082015250565b600061161d602983611187565b9150611628826115c1565b604082019050919050565b6000602082019050818103600083015261164c81611610565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061169a57607f821691505b6020821081036116ad576116ac611653565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026117157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116d8565b61171f86836116d8565b95508019841693508086168417925050509392505050565b6000819050919050565b600061175c611757611752846110b7565b611737565b6110b7565b9050919050565b6000819050919050565b61177683611741565b61178a61178282611763565b8484546116e5565b825550505050565b600090565b61179f611792565b6117aa81848461176d565b505050565b5b818110156117ce576117c3600082611797565b6001810190506117b0565b5050565b601f821115611813576117e4816116b3565b6117ed846116c8565b810160208510156117fc578190505b611810611808856116c8565b8301826117af565b50505b505050565b600082821c905092915050565b600061183660001984600802611818565b1980831691505092915050565b600061184f8383611825565b9150826002028217905092915050565b6118688261117c565b67ffffffffffffffff811115611881576118806112b1565b5b61188b8254611682565b6118968282856117d2565b600060209050601f8311600181146118c957600084156118b7578287015190505b6118c18582611843565b865550611929565b601f1984166118d7866116b3565b60005b828110156118ff578489015182556001820191506020850194506020810190506118da565b8683101561191c5784890151611918601f891682611825565b8355505b6001600288020188555050505b505050505050565b7f43727970746f57697a617264734465636b7b6372656174654465636b7d3a206d60008201527f73672e73656e646572206861732061206465636b000000000000000000000000602082015250565b600061198d603483611187565b915061199882611931565b604082019050919050565b600060208201905081810360008301526119bc81611980565b9050919050565b60006040820190506119d8600083018561116d565b6119e5602083018461127d565b9392505050565b600067ffffffffffffffff821115611a0757611a066112b1565b5b602082029050919050565b600080fd5b600081519050611a26816110c1565b92915050565b6000611a3f611a3a846119ec565b611311565b90508060208402830185811115611a5957611a58611a12565b5b835b81811015611a825780611a6e8882611a17565b845260208401935050602081019050611a5b565b5050509392505050565b600082601f830112611aa157611aa06112a7565b5b6004611aae848285611a2c565b91505092915050565b600060808284031215611acd57611acc61104f565b5b6000611adb84828501611a8c565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611b6f602683611187565b9150611b7a82611b13565b604082019050919050565b60006020820190508181036000830152611b9e81611b62565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611bdb602083611187565b9150611be682611ba5565b602082019050919050565b60006020820190508181036000830152611c0a81611bce565b905091905056fea2646970667358221220ef69ea0accfd2eccab047f9d1a9ab042f69df4949c4970eb528703fe2a861b6264736f6c634300080f0033",
}

// CWDABI is the input ABI used to generate the binding from.
// Deprecated: Use CWDMetaData.ABI instead.
var CWDABI = CWDMetaData.ABI

// CWDBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CWDMetaData.Bin instead.
var CWDBin = CWDMetaData.Bin

// DeployCWD deploys a new Ethereum contract, binding an instance of CWD to it.
func DeployCWD(auth *bind.TransactOpts, backend bind.ContractBackend, _cwc common.Address) (common.Address, *types.Transaction, *CWD, error) {
	parsed, err := CWDMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CWDBin), backend, _cwc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CWD{CWDCaller: CWDCaller{contract: contract}, CWDTransactor: CWDTransactor{contract: contract}, CWDFilterer: CWDFilterer{contract: contract}}, nil
}

// CWD is an auto generated Go binding around an Ethereum contract.
type CWD struct {
	CWDCaller     // Read-only binding to the contract
	CWDTransactor // Write-only binding to the contract
	CWDFilterer   // Log filterer for contract events
}

// CWDCaller is an auto generated read-only Go binding around an Ethereum contract.
type CWDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CWDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CWDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CWDSession struct {
	Contract     *CWD              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CWDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CWDCallerSession struct {
	Contract *CWDCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CWDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CWDTransactorSession struct {
	Contract     *CWDTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CWDRaw is an auto generated low-level Go binding around an Ethereum contract.
type CWDRaw struct {
	Contract *CWD // Generic contract binding to access the raw methods on
}

// CWDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CWDCallerRaw struct {
	Contract *CWDCaller // Generic read-only contract binding to access the raw methods on
}

// CWDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CWDTransactorRaw struct {
	Contract *CWDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCWD creates a new instance of CWD, bound to a specific deployed contract.
func NewCWD(address common.Address, backend bind.ContractBackend) (*CWD, error) {
	contract, err := bindCWD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CWD{CWDCaller: CWDCaller{contract: contract}, CWDTransactor: CWDTransactor{contract: contract}, CWDFilterer: CWDFilterer{contract: contract}}, nil
}

// NewCWDCaller creates a new read-only instance of CWD, bound to a specific deployed contract.
func NewCWDCaller(address common.Address, caller bind.ContractCaller) (*CWDCaller, error) {
	contract, err := bindCWD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CWDCaller{contract: contract}, nil
}

// NewCWDTransactor creates a new write-only instance of CWD, bound to a specific deployed contract.
func NewCWDTransactor(address common.Address, transactor bind.ContractTransactor) (*CWDTransactor, error) {
	contract, err := bindCWD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CWDTransactor{contract: contract}, nil
}

// NewCWDFilterer creates a new log filterer instance of CWD, bound to a specific deployed contract.
func NewCWDFilterer(address common.Address, filterer bind.ContractFilterer) (*CWDFilterer, error) {
	contract, err := bindCWD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CWDFilterer{contract: contract}, nil
}

// bindCWD binds a generic wrapper to an already deployed contract.
func bindCWD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CWDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CWD *CWDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CWD.Contract.CWDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CWD *CWDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CWD.Contract.CWDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CWD *CWDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CWD.Contract.CWDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CWD *CWDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CWD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CWD *CWDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CWD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CWD *CWDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CWD.Contract.contract.Transact(opts, method, params...)
}

// CryptoWizardsCard is a free data retrieval call binding the contract method 0x2ed21357.
//
// Solidity: function cryptoWizardsCard() view returns(address)
func (_CWD *CWDCaller) CryptoWizardsCard(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "cryptoWizardsCard")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CryptoWizardsCard is a free data retrieval call binding the contract method 0x2ed21357.
//
// Solidity: function cryptoWizardsCard() view returns(address)
func (_CWD *CWDSession) CryptoWizardsCard() (common.Address, error) {
	return _CWD.Contract.CryptoWizardsCard(&_CWD.CallOpts)
}

// CryptoWizardsCard is a free data retrieval call binding the contract method 0x2ed21357.
//
// Solidity: function cryptoWizardsCard() view returns(address)
func (_CWD *CWDCallerSession) CryptoWizardsCard() (common.Address, error) {
	return _CWD.Contract.CryptoWizardsCard(&_CWD.CallOpts)
}

// DeckURIs is a free data retrieval call binding the contract method 0xd4f2df3e.
//
// Solidity: function deckURIs(uint256 ) view returns(string)
func (_CWD *CWDCaller) DeckURIs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "deckURIs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DeckURIs is a free data retrieval call binding the contract method 0xd4f2df3e.
//
// Solidity: function deckURIs(uint256 ) view returns(string)
func (_CWD *CWDSession) DeckURIs(arg0 *big.Int) (string, error) {
	return _CWD.Contract.DeckURIs(&_CWD.CallOpts, arg0)
}

// DeckURIs is a free data retrieval call binding the contract method 0xd4f2df3e.
//
// Solidity: function deckURIs(uint256 ) view returns(string)
func (_CWD *CWDCallerSession) DeckURIs(arg0 *big.Int) (string, error) {
	return _CWD.Contract.DeckURIs(&_CWD.CallOpts, arg0)
}

// Decks is a free data retrieval call binding the contract method 0x2c9fad06.
//
// Solidity: function decks(address ) view returns(uint256 collectionId, string deckURI, uint256 card1, uint256 card2, uint256 card3, uint256 card4)
func (_CWD *CWDCaller) Decks(opts *bind.CallOpts, arg0 common.Address) (struct {
	CollectionId *big.Int
	DeckURI      string
	Card1        *big.Int
	Card2        *big.Int
	Card3        *big.Int
	Card4        *big.Int
}, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "decks", arg0)

	outstruct := new(struct {
		CollectionId *big.Int
		DeckURI      string
		Card1        *big.Int
		Card2        *big.Int
		Card3        *big.Int
		Card4        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollectionId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DeckURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Card1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Card2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Card3 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Card4 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Decks is a free data retrieval call binding the contract method 0x2c9fad06.
//
// Solidity: function decks(address ) view returns(uint256 collectionId, string deckURI, uint256 card1, uint256 card2, uint256 card3, uint256 card4)
func (_CWD *CWDSession) Decks(arg0 common.Address) (struct {
	CollectionId *big.Int
	DeckURI      string
	Card1        *big.Int
	Card2        *big.Int
	Card3        *big.Int
	Card4        *big.Int
}, error) {
	return _CWD.Contract.Decks(&_CWD.CallOpts, arg0)
}

// Decks is a free data retrieval call binding the contract method 0x2c9fad06.
//
// Solidity: function decks(address ) view returns(uint256 collectionId, string deckURI, uint256 card1, uint256 card2, uint256 card3, uint256 card4)
func (_CWD *CWDCallerSession) Decks(arg0 common.Address) (struct {
	CollectionId *big.Int
	DeckURI      string
	Card1        *big.Int
	Card2        *big.Int
	Card3        *big.Int
	Card4        *big.Int
}, error) {
	return _CWD.Contract.Decks(&_CWD.CallOpts, arg0)
}

// GetUserDeck is a free data retrieval call binding the contract method 0x822b6eeb.
//
// Solidity: function getUserDeck(address _user) view returns((uint256,string,uint256,uint256,uint256,uint256) deck)
func (_CWD *CWDCaller) GetUserDeck(opts *bind.CallOpts, _user common.Address) (CryptoWizardsDeckDeck, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "getUserDeck", _user)

	if err != nil {
		return *new(CryptoWizardsDeckDeck), err
	}

	out0 := *abi.ConvertType(out[0], new(CryptoWizardsDeckDeck)).(*CryptoWizardsDeckDeck)

	return out0, err

}

// GetUserDeck is a free data retrieval call binding the contract method 0x822b6eeb.
//
// Solidity: function getUserDeck(address _user) view returns((uint256,string,uint256,uint256,uint256,uint256) deck)
func (_CWD *CWDSession) GetUserDeck(_user common.Address) (CryptoWizardsDeckDeck, error) {
	return _CWD.Contract.GetUserDeck(&_CWD.CallOpts, _user)
}

// GetUserDeck is a free data retrieval call binding the contract method 0x822b6eeb.
//
// Solidity: function getUserDeck(address _user) view returns((uint256,string,uint256,uint256,uint256,uint256) deck)
func (_CWD *CWDCallerSession) GetUserDeck(_user common.Address) (CryptoWizardsDeckDeck, error) {
	return _CWD.Contract.GetUserDeck(&_CWD.CallOpts, _user)
}

// HasDeck is a free data retrieval call binding the contract method 0xda583940.
//
// Solidity: function hasDeck(address ) view returns(bool)
func (_CWD *CWDCaller) HasDeck(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "hasDeck", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeck is a free data retrieval call binding the contract method 0xda583940.
//
// Solidity: function hasDeck(address ) view returns(bool)
func (_CWD *CWDSession) HasDeck(arg0 common.Address) (bool, error) {
	return _CWD.Contract.HasDeck(&_CWD.CallOpts, arg0)
}

// HasDeck is a free data retrieval call binding the contract method 0xda583940.
//
// Solidity: function hasDeck(address ) view returns(bool)
func (_CWD *CWDCallerSession) HasDeck(arg0 common.Address) (bool, error) {
	return _CWD.Contract.HasDeck(&_CWD.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CWD *CWDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CWD.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CWD *CWDSession) Owner() (common.Address, error) {
	return _CWD.Contract.Owner(&_CWD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CWD *CWDCallerSession) Owner() (common.Address, error) {
	return _CWD.Contract.Owner(&_CWD.CallOpts)
}

// CreateDeck is a paid mutator transaction binding the contract method 0xb6781b78.
//
// Solidity: function createDeck(uint256 _collectionId) returns()
func (_CWD *CWDTransactor) CreateDeck(opts *bind.TransactOpts, _collectionId *big.Int) (*types.Transaction, error) {
	return _CWD.contract.Transact(opts, "createDeck", _collectionId)
}

// CreateDeck is a paid mutator transaction binding the contract method 0xb6781b78.
//
// Solidity: function createDeck(uint256 _collectionId) returns()
func (_CWD *CWDSession) CreateDeck(_collectionId *big.Int) (*types.Transaction, error) {
	return _CWD.Contract.CreateDeck(&_CWD.TransactOpts, _collectionId)
}

// CreateDeck is a paid mutator transaction binding the contract method 0xb6781b78.
//
// Solidity: function createDeck(uint256 _collectionId) returns()
func (_CWD *CWDTransactorSession) CreateDeck(_collectionId *big.Int) (*types.Transaction, error) {
	return _CWD.Contract.CreateDeck(&_CWD.TransactOpts, _collectionId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CWD *CWDTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CWD.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CWD *CWDSession) RenounceOwnership() (*types.Transaction, error) {
	return _CWD.Contract.RenounceOwnership(&_CWD.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CWD *CWDTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CWD.Contract.RenounceOwnership(&_CWD.TransactOpts)
}

// SetDeckURI is a paid mutator transaction binding the contract method 0x4be5154d.
//
// Solidity: function setDeckURI(uint256 collectionId, string deckURI) returns()
func (_CWD *CWDTransactor) SetDeckURI(opts *bind.TransactOpts, collectionId *big.Int, deckURI string) (*types.Transaction, error) {
	return _CWD.contract.Transact(opts, "setDeckURI", collectionId, deckURI)
}

// SetDeckURI is a paid mutator transaction binding the contract method 0x4be5154d.
//
// Solidity: function setDeckURI(uint256 collectionId, string deckURI) returns()
func (_CWD *CWDSession) SetDeckURI(collectionId *big.Int, deckURI string) (*types.Transaction, error) {
	return _CWD.Contract.SetDeckURI(&_CWD.TransactOpts, collectionId, deckURI)
}

// SetDeckURI is a paid mutator transaction binding the contract method 0x4be5154d.
//
// Solidity: function setDeckURI(uint256 collectionId, string deckURI) returns()
func (_CWD *CWDTransactorSession) SetDeckURI(collectionId *big.Int, deckURI string) (*types.Transaction, error) {
	return _CWD.Contract.SetDeckURI(&_CWD.TransactOpts, collectionId, deckURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CWD *CWDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CWD.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CWD *CWDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CWD.Contract.TransferOwnership(&_CWD.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CWD *CWDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CWD.Contract.TransferOwnership(&_CWD.TransactOpts, newOwner)
}

// UpdateDeck is a paid mutator transaction binding the contract method 0x1c86871d.
//
// Solidity: function updateDeck(address _owner, uint256 _select, uint256 _cardId) returns()
func (_CWD *CWDTransactor) UpdateDeck(opts *bind.TransactOpts, _owner common.Address, _select *big.Int, _cardId *big.Int) (*types.Transaction, error) {
	return _CWD.contract.Transact(opts, "updateDeck", _owner, _select, _cardId)
}

// UpdateDeck is a paid mutator transaction binding the contract method 0x1c86871d.
//
// Solidity: function updateDeck(address _owner, uint256 _select, uint256 _cardId) returns()
func (_CWD *CWDSession) UpdateDeck(_owner common.Address, _select *big.Int, _cardId *big.Int) (*types.Transaction, error) {
	return _CWD.Contract.UpdateDeck(&_CWD.TransactOpts, _owner, _select, _cardId)
}

// UpdateDeck is a paid mutator transaction binding the contract method 0x1c86871d.
//
// Solidity: function updateDeck(address _owner, uint256 _select, uint256 _cardId) returns()
func (_CWD *CWDTransactorSession) UpdateDeck(_owner common.Address, _select *big.Int, _cardId *big.Int) (*types.Transaction, error) {
	return _CWD.Contract.UpdateDeck(&_CWD.TransactOpts, _owner, _select, _cardId)
}

// CWDOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CWD contract.
type CWDOwnershipTransferredIterator struct {
	Event *CWDOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CWDOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CWDOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CWDOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CWDOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CWDOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CWDOwnershipTransferred represents a OwnershipTransferred event raised by the CWD contract.
type CWDOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CWD *CWDFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CWDOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CWD.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CWDOwnershipTransferredIterator{contract: _CWD.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CWD *CWDFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CWDOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CWD.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CWDOwnershipTransferred)
				if err := _CWD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CWD *CWDFilterer) ParseOwnershipTransferred(log types.Log) (*CWDOwnershipTransferred, error) {
	event := new(CWDOwnershipTransferred)
	if err := _CWD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
