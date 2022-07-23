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

// CryptoWizardsBetPlayer is an auto generated low-level Go binding around an user-defined struct.
type CryptoWizardsBetPlayer struct {
	Addr    common.Address
	BetMade bool
	CardId  *big.Int
}

// CWBMetaData contains all meta data concerning the CWB contract.
var CWBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"BetMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"EndGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cwc\",\"type\":\"address\"}],\"name\":\"_setCWC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"bet1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"bet2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cryptoWizardsCards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"endGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"betMade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoWizardsBet.Player\",\"name\":\"player1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"betMade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoWizardsBet.Player\",\"name\":\"player2\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"gameEnded\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"p1CardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p2CardId\",\"type\":\"uint256\"}],\"name\":\"newGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"removeBet1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"removeBet2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wizard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600160008190555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506128dd806100696000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806384a762131161006657806384a7621314610147578063aabbbd6514610163578063ab747b6414610181578063b1d89e991461019d578063b63ce384146101b95761009e565b806301c93690146100a3578063117a5b90146100bf5780631cd93b73146100f35780634e4988be1461010f5780635ca934dc1461012b575b600080fd5b6100bd60048036038101906100b891906122c9565b6101d7565b005b6100d960048036038101906100d49190612330565b61069f565b6040516100ea959493929190612405565b60405180910390f35b61010d60048036038101906101089190612330565b61080c565b005b61012960048036038101906101249190612330565b610db2565b005b6101456004803603810190610140919061245a565b6112c2565b005b610161600480360381019061015c9190612330565b611396565b005b61016b6118a6565b6040516101789190612487565b60405180910390f35b61019b600480360381019061019691906124a2565b6118cc565b005b6101b760048036038101906101b29190612330565b611b94565b005b6101c161213a565b6040516101ce9190612487565b60405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610267576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025e90612565565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b81526004016102d99190612585565b602060405180830381865afa1580156102f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031a91906125b5565b73ffffffffffffffffffffffffffffffffffffffff1614801561040357508273ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b81526004016103aa9190612585565b602060405180830381865afa1580156103c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103eb91906125b5565b73ffffffffffffffffffffffffffffffffffffffff16145b61040c57600080fd5b600060405180606001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001848152509050600060405180606001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020018481525090506001600081548092919061048f90612611565b919050555060006040518060a001604052806001548152602001848152602001838152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152509050806004600060015481526020019081526020016000206000820151816000015560208201518160010160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155505060408201518160030160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155505060608201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160050160146101000a81548160ff0219169083151502179055509050507fa0b7f6f22bb4f69adc52dd16ffe2a964ab8b577be274247fa9a1547b03ce3e5560015460405161068e9190612585565b60405180910390a150505050505050565b6004602052806000526040600020600091509050806000015490806001016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1615151515815260200160018201548152505090806003016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff16151515158152602001600182015481525050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160149054906101000a900460ff16905085565b60028160018203610a17573373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e60046000858152602001908152602001600020600101600101546040518263ffffffff1660e01b81526004016108a29190612585565b602060405180830381865afa1580156108bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e391906125b5565b73ffffffffffffffffffffffffffffffffffffffff1614801561096a57506004600082815260200190815260200160002060010160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80156109a05750600015156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b80156109d35750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b610a12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a09906126cb565b60405180910390fd5b610c13565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e60046000858152602001908152602001600020600301600101546040518263ffffffff1660e01b8152600401610aa29190612585565b602060405180830381865afa158015610abf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae391906125b5565b73ffffffffffffffffffffffffffffffffffffffff16148015610b6a57506004600082815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8015610ba05750600015156004600083815260200190815260200160002060030160000160149054906101000a900460ff161515145b8015610bd35750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b610c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c09906126cb565b60405180910390fd5b5b600260005403610c58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4f90612737565b60405180910390fd5b6002600081905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333060046000888152602001908152602001600020600301600101546040518463ffffffff1660e01b8152600401610cd893929190612757565b600060405180830381600087803b158015610cf257600080fd5b505af1158015610d06573d6000803e3d6000fd5b5050505060016004600085815260200190815260200160002060010160000160146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f942ce7d81758c4cf68fa7638d68d695ae1c785e10b2413ccca93098d14f4cf7a846004600087815260200190815260200160002060030160010154604051610d9d92919061278e565b60405180910390a26001600081905550505050565b60018160018203610f08576004600082815260200190815260200160002060010160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148015610e5b5750600115156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b8015610e915750600015156004600083815260200190815260200160002060030160000160149054906101000a900460ff161515145b8015610ec45750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b610f03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610efa90612829565b60405180910390fd5b611048565b6004600082815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148015610f9f57506004600082815260200190815260200160002060030160000160149054906101000a900460ff165b8015610fd55750600015156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b80156110085750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b611047576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103e90612829565b60405180910390fd5b5b60026000540361108d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108490612737565b60405180910390fd5b6002600081905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd303360046000888152602001908152602001600020600101600101546040518463ffffffff1660e01b815260040161110d93929190612757565b600060405180830381600087803b15801561112757600080fd5b505af115801561113b573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd306004600087815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660046000888152602001908152602001600020600301600101546040518463ffffffff1660e01b81526004016111f093929190612757565b600060405180830381600087803b15801561120a57600080fd5b505af115801561121e573d6000803e3d6000fd5b5050505060016004600085815260200190815260200160002060050160146101000a81548160ff02191690831515021790555060006004600085815260200190815260200160002060010160000160146101000a81548160ff02191690831515021790555060006004600085815260200190815260200160002060030160000160146101000a81548160ff0219169083151502179055506001600081905550505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161134990612565565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600281600182036114ec576004600082815260200190815260200160002060010160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614801561143f5750600115156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b80156114755750600015156004600083815260200190815260200160002060030160000160149054906101000a900460ff161515145b80156114a85750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b6114e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114de90612829565b60405180910390fd5b61162c565b6004600082815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614801561158357506004600082815260200190815260200160002060030160000160149054906101000a900460ff165b80156115b95750600015156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b80156115ec5750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b61162b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162290612829565b60405180910390fd5b5b600260005403611671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161166890612737565b60405180910390fd5b6002600081905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd303360046000888152602001908152602001600020600101600101546040518463ffffffff1660e01b81526004016116f193929190612757565b600060405180830381600087803b15801561170b57600080fd5b505af115801561171f573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd306004600087815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660046000888152602001908152602001600020600301600101546040518463ffffffff1660e01b81526004016117d493929190612757565b600060405180830381600087803b1580156117ee57600080fd5b505af1158015611802573d6000803e3d6000fd5b5050505060016004600085815260200190815260200160002060050160146101000a81548160ff02191690831515021790555060006004600085815260200190815260200160002060010160000160146101000a81548160ff02191690831515021790555060006004600085815260200190815260200160002060030160000160146101000a81548160ff0219169083151502179055506001600081905550505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461195c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195390612565565b60405180910390fd5b806004600082815260200190815260200160002060010160000160149054906101000a900460ff1680156119b357506004600082815260200190815260200160002060030160000160149054906101000a900460ff165b80156119e65750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b6119ef57600080fd5b6000600267ffffffffffffffff811115611a0c57611a0b612849565b5b604051908082528060200260200182016040528015611a3a5781602001602082028036833780820191505090505b509050600460008481526020019081526020016000206001016001015481600081518110611a6b57611a6a612878565b5b602002602001018181525050600460008481526020019081526020016000206003016001015481600181518110611aa557611aa4612878565b5b602002602001018181525050611abc813086612160565b60016004600085815260200190815260200160002060050160146101000a81548160ff021916908315150217905550836004600085815260200190815260200160002060050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff167f84f8e1147ed860fdcb13c73d1ec1bf4aa9cded9250811ad1067ea2e81b61affd84604051611b869190612585565b60405180910390a250505050565b60018160018203611d9f573373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e60046000858152602001908152602001600020600101600101546040518263ffffffff1660e01b8152600401611c2a9190612585565b602060405180830381865afa158015611c47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6b91906125b5565b73ffffffffffffffffffffffffffffffffffffffff16148015611cf257506004600082815260200190815260200160002060010160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8015611d285750600015156004600083815260200190815260200160002060010160000160149054906101000a900460ff161515145b8015611d5b5750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b611d9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d91906126cb565b60405180910390fd5b611f9b565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e60046000858152602001908152602001600020600301600101546040518263ffffffff1660e01b8152600401611e2a9190612585565b602060405180830381865afa158015611e47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e6b91906125b5565b73ffffffffffffffffffffffffffffffffffffffff16148015611ef257506004600082815260200190815260200160002060030160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8015611f285750600015156004600083815260200190815260200160002060030160000160149054906101000a900460ff161515145b8015611f5b5750600015156004600083815260200190815260200160002060050160149054906101000a900460ff161515145b611f9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f91906126cb565b60405180910390fd5b5b600260005403611fe0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fd790612737565b60405180910390fd5b6002600081905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333060046000888152602001908152602001600020600101600101546040518463ffffffff1660e01b815260040161206093929190612757565b600060405180830381600087803b15801561207a57600080fd5b505af115801561208e573d6000803e3d6000fd5b5050505060016004600085815260200190815260200160002060010160000160146101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f942ce7d81758c4cf68fa7638d68d695ae1c785e10b2413ccca93098d14f4cf7a84600460008781526020019081526020016000206001016001015460405161212592919061278e565b60405180910390a26001600081905550505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60005b835181101561222a57600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd84848785815181106121bf576121be612878565b5b60200260200101516040518463ffffffff1660e01b81526004016121e593929190612757565b600060405180830381600087803b1580156121ff57600080fd5b505af1158015612213573d6000803e3d6000fd5b50505050808061222290612611565b915050612163565b50505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061226082612235565b9050919050565b61227081612255565b811461227b57600080fd5b50565b60008135905061228d81612267565b92915050565b6000819050919050565b6122a681612293565b81146122b157600080fd5b50565b6000813590506122c38161229d565b92915050565b600080600080608085870312156122e3576122e2612230565b5b60006122f18782880161227e565b94505060206123028782880161227e565b9350506040612313878288016122b4565b9250506060612324878288016122b4565b91505092959194509250565b60006020828403121561234657612345612230565b5b6000612354848285016122b4565b91505092915050565b61236681612293565b82525050565b61237581612255565b82525050565b60008115159050919050565b6123908161237b565b82525050565b61239f81612293565b82525050565b6060820160008201516123bb600085018261236c565b5060208201516123ce6020850182612387565b5060408201516123e16040850182612396565b50505050565b6123f081612255565b82525050565b6123ff8161237b565b82525050565b60006101208201905061241b600083018861235d565b61242860208301876123a5565b61243560808301866123a5565b61244260e08301856123e7565b6124506101008301846123f6565b9695505050505050565b6000602082840312156124705761246f612230565b5b600061247e8482850161227e565b91505092915050565b600060208201905061249c60008301846123e7565b92915050565b600080604083850312156124b9576124b8612230565b5b60006124c78582860161227e565b92505060206124d8858286016122b4565b9150509250929050565b600082825260208201905092915050565b7f43727970746f57697a617264737b6f6e6c79537461727465727d3a206163636560008201527f73732064656e6965640000000000000000000000000000000000000000000000602082015250565b600061254f6029836124e2565b915061255a826124f3565b604082019050919050565b6000602082019050818103600083015261257e81612542565b9050919050565b600060208201905061259a600083018461235d565b92915050565b6000815190506125af81612267565b92915050565b6000602082840312156125cb576125ca612230565b5b60006125d9848285016125a0565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061261c82612293565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361264e5761264d6125e2565b5b600182019050919050565b7f43727970746f57697a617264734265747b6265745265766965777d3a20696e7660008201527f616c696420636f6e646974696f6e730000000000000000000000000000000000602082015250565b60006126b5602f836124e2565b91506126c082612659565b604082019050919050565b600060208201905081810360008301526126e4816126a8565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000612721601f836124e2565b915061272c826126eb565b602082019050919050565b6000602082019050818103600083015261275081612714565b9050919050565b600060608201905061276c60008301866123e7565b61277960208301856123e7565b612786604083018461235d565b949350505050565b60006040820190506127a3600083018561235d565b6127b0602083018461235d565b9392505050565b7f43727970746f57697a617264734265747b72656d6f766542657452657669657760008201527f7d3a20696e76616c696420636f6e646974696f6e730000000000000000000000602082015250565b60006128136035836124e2565b915061281e826127b7565b604082019050919050565b6000602082019050818103600083015261284281612806565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220a11c689b2a9162f3d04282950ecc00e6422582973cdcc51937c974ce979e16b264736f6c634300080f0033",
}

// CWBABI is the input ABI used to generate the binding from.
// Deprecated: Use CWBMetaData.ABI instead.
var CWBABI = CWBMetaData.ABI

// CWBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CWBMetaData.Bin instead.
var CWBBin = CWBMetaData.Bin

// DeployCWB deploys a new Ethereum contract, binding an instance of CWB to it.
func DeployCWB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CWB, error) {
	parsed, err := CWBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CWBBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CWB{CWBCaller: CWBCaller{contract: contract}, CWBTransactor: CWBTransactor{contract: contract}, CWBFilterer: CWBFilterer{contract: contract}}, nil
}

// CWB is an auto generated Go binding around an Ethereum contract.
type CWB struct {
	CWBCaller     // Read-only binding to the contract
	CWBTransactor // Write-only binding to the contract
	CWBFilterer   // Log filterer for contract events
}

// CWBCaller is an auto generated read-only Go binding around an Ethereum contract.
type CWBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CWBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CWBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CWBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CWBSession struct {
	Contract     *CWB              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CWBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CWBCallerSession struct {
	Contract *CWBCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CWBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CWBTransactorSession struct {
	Contract     *CWBTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CWBRaw is an auto generated low-level Go binding around an Ethereum contract.
type CWBRaw struct {
	Contract *CWB // Generic contract binding to access the raw methods on
}

// CWBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CWBCallerRaw struct {
	Contract *CWBCaller // Generic read-only contract binding to access the raw methods on
}

// CWBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CWBTransactorRaw struct {
	Contract *CWBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCWB creates a new instance of CWB, bound to a specific deployed contract.
func NewCWB(address common.Address, backend bind.ContractBackend) (*CWB, error) {
	contract, err := bindCWB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CWB{CWBCaller: CWBCaller{contract: contract}, CWBTransactor: CWBTransactor{contract: contract}, CWBFilterer: CWBFilterer{contract: contract}}, nil
}

// NewCWBCaller creates a new read-only instance of CWB, bound to a specific deployed contract.
func NewCWBCaller(address common.Address, caller bind.ContractCaller) (*CWBCaller, error) {
	contract, err := bindCWB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CWBCaller{contract: contract}, nil
}

// NewCWBTransactor creates a new write-only instance of CWB, bound to a specific deployed contract.
func NewCWBTransactor(address common.Address, transactor bind.ContractTransactor) (*CWBTransactor, error) {
	contract, err := bindCWB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CWBTransactor{contract: contract}, nil
}

// NewCWBFilterer creates a new log filterer instance of CWB, bound to a specific deployed contract.
func NewCWBFilterer(address common.Address, filterer bind.ContractFilterer) (*CWBFilterer, error) {
	contract, err := bindCWB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CWBFilterer{contract: contract}, nil
}

// bindCWB binds a generic wrapper to an already deployed contract.
func bindCWB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CWBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CWB *CWBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CWB.Contract.CWBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CWB *CWBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CWB.Contract.CWBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CWB *CWBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CWB.Contract.CWBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CWB *CWBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CWB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CWB *CWBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CWB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CWB *CWBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CWB.Contract.contract.Transact(opts, method, params...)
}

// CryptoWizardsCards is a free data retrieval call binding the contract method 0xb63ce384.
//
// Solidity: function cryptoWizardsCards() view returns(address)
func (_CWB *CWBCaller) CryptoWizardsCards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CWB.contract.Call(opts, &out, "cryptoWizardsCards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CryptoWizardsCards is a free data retrieval call binding the contract method 0xb63ce384.
//
// Solidity: function cryptoWizardsCards() view returns(address)
func (_CWB *CWBSession) CryptoWizardsCards() (common.Address, error) {
	return _CWB.Contract.CryptoWizardsCards(&_CWB.CallOpts)
}

// CryptoWizardsCards is a free data retrieval call binding the contract method 0xb63ce384.
//
// Solidity: function cryptoWizardsCards() view returns(address)
func (_CWB *CWBCallerSession) CryptoWizardsCards() (common.Address, error) {
	return _CWB.Contract.CryptoWizardsCards(&_CWB.CallOpts)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 gameId, (address,bool,uint256) player1, (address,bool,uint256) player2, address winner, bool gameEnded)
func (_CWB *CWBCaller) Games(opts *bind.CallOpts, arg0 *big.Int) (struct {
	GameId    *big.Int
	Player1   CryptoWizardsBetPlayer
	Player2   CryptoWizardsBetPlayer
	Winner    common.Address
	GameEnded bool
}, error) {
	var out []interface{}
	err := _CWB.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		GameId    *big.Int
		Player1   CryptoWizardsBetPlayer
		Player2   CryptoWizardsBetPlayer
		Winner    common.Address
		GameEnded bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Player1 = *abi.ConvertType(out[1], new(CryptoWizardsBetPlayer)).(*CryptoWizardsBetPlayer)
	outstruct.Player2 = *abi.ConvertType(out[2], new(CryptoWizardsBetPlayer)).(*CryptoWizardsBetPlayer)
	outstruct.Winner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.GameEnded = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 gameId, (address,bool,uint256) player1, (address,bool,uint256) player2, address winner, bool gameEnded)
func (_CWB *CWBSession) Games(arg0 *big.Int) (struct {
	GameId    *big.Int
	Player1   CryptoWizardsBetPlayer
	Player2   CryptoWizardsBetPlayer
	Winner    common.Address
	GameEnded bool
}, error) {
	return _CWB.Contract.Games(&_CWB.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(uint256 gameId, (address,bool,uint256) player1, (address,bool,uint256) player2, address winner, bool gameEnded)
func (_CWB *CWBCallerSession) Games(arg0 *big.Int) (struct {
	GameId    *big.Int
	Player1   CryptoWizardsBetPlayer
	Player2   CryptoWizardsBetPlayer
	Winner    common.Address
	GameEnded bool
}, error) {
	return _CWB.Contract.Games(&_CWB.CallOpts, arg0)
}

// Wizard is a free data retrieval call binding the contract method 0xaabbbd65.
//
// Solidity: function wizard() view returns(address)
func (_CWB *CWBCaller) Wizard(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CWB.contract.Call(opts, &out, "wizard")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wizard is a free data retrieval call binding the contract method 0xaabbbd65.
//
// Solidity: function wizard() view returns(address)
func (_CWB *CWBSession) Wizard() (common.Address, error) {
	return _CWB.Contract.Wizard(&_CWB.CallOpts)
}

// Wizard is a free data retrieval call binding the contract method 0xaabbbd65.
//
// Solidity: function wizard() view returns(address)
func (_CWB *CWBCallerSession) Wizard() (common.Address, error) {
	return _CWB.Contract.Wizard(&_CWB.CallOpts)
}

// SetCWC is a paid mutator transaction binding the contract method 0x5ca934dc.
//
// Solidity: function _setCWC(address _cwc) returns()
func (_CWB *CWBTransactor) SetCWC(opts *bind.TransactOpts, _cwc common.Address) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "_setCWC", _cwc)
}

// SetCWC is a paid mutator transaction binding the contract method 0x5ca934dc.
//
// Solidity: function _setCWC(address _cwc) returns()
func (_CWB *CWBSession) SetCWC(_cwc common.Address) (*types.Transaction, error) {
	return _CWB.Contract.SetCWC(&_CWB.TransactOpts, _cwc)
}

// SetCWC is a paid mutator transaction binding the contract method 0x5ca934dc.
//
// Solidity: function _setCWC(address _cwc) returns()
func (_CWB *CWBTransactorSession) SetCWC(_cwc common.Address) (*types.Transaction, error) {
	return _CWB.Contract.SetCWC(&_CWB.TransactOpts, _cwc)
}

// Bet1 is a paid mutator transaction binding the contract method 0xb1d89e99.
//
// Solidity: function bet1(uint256 _gameId) returns()
func (_CWB *CWBTransactor) Bet1(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "bet1", _gameId)
}

// Bet1 is a paid mutator transaction binding the contract method 0xb1d89e99.
//
// Solidity: function bet1(uint256 _gameId) returns()
func (_CWB *CWBSession) Bet1(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.Bet1(&_CWB.TransactOpts, _gameId)
}

// Bet1 is a paid mutator transaction binding the contract method 0xb1d89e99.
//
// Solidity: function bet1(uint256 _gameId) returns()
func (_CWB *CWBTransactorSession) Bet1(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.Bet1(&_CWB.TransactOpts, _gameId)
}

// Bet2 is a paid mutator transaction binding the contract method 0x1cd93b73.
//
// Solidity: function bet2(uint256 _gameId) returns()
func (_CWB *CWBTransactor) Bet2(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "bet2", _gameId)
}

// Bet2 is a paid mutator transaction binding the contract method 0x1cd93b73.
//
// Solidity: function bet2(uint256 _gameId) returns()
func (_CWB *CWBSession) Bet2(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.Bet2(&_CWB.TransactOpts, _gameId)
}

// Bet2 is a paid mutator transaction binding the contract method 0x1cd93b73.
//
// Solidity: function bet2(uint256 _gameId) returns()
func (_CWB *CWBTransactorSession) Bet2(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.Bet2(&_CWB.TransactOpts, _gameId)
}

// EndGame is a paid mutator transaction binding the contract method 0xab747b64.
//
// Solidity: function endGame(address _winner, uint256 _gameId) returns()
func (_CWB *CWBTransactor) EndGame(opts *bind.TransactOpts, _winner common.Address, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "endGame", _winner, _gameId)
}

// EndGame is a paid mutator transaction binding the contract method 0xab747b64.
//
// Solidity: function endGame(address _winner, uint256 _gameId) returns()
func (_CWB *CWBSession) EndGame(_winner common.Address, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.EndGame(&_CWB.TransactOpts, _winner, _gameId)
}

// EndGame is a paid mutator transaction binding the contract method 0xab747b64.
//
// Solidity: function endGame(address _winner, uint256 _gameId) returns()
func (_CWB *CWBTransactorSession) EndGame(_winner common.Address, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.EndGame(&_CWB.TransactOpts, _winner, _gameId)
}

// NewGame is a paid mutator transaction binding the contract method 0x01c93690.
//
// Solidity: function newGame(address p1, address p2, uint256 p1CardId, uint256 p2CardId) returns()
func (_CWB *CWBTransactor) NewGame(opts *bind.TransactOpts, p1 common.Address, p2 common.Address, p1CardId *big.Int, p2CardId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "newGame", p1, p2, p1CardId, p2CardId)
}

// NewGame is a paid mutator transaction binding the contract method 0x01c93690.
//
// Solidity: function newGame(address p1, address p2, uint256 p1CardId, uint256 p2CardId) returns()
func (_CWB *CWBSession) NewGame(p1 common.Address, p2 common.Address, p1CardId *big.Int, p2CardId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.NewGame(&_CWB.TransactOpts, p1, p2, p1CardId, p2CardId)
}

// NewGame is a paid mutator transaction binding the contract method 0x01c93690.
//
// Solidity: function newGame(address p1, address p2, uint256 p1CardId, uint256 p2CardId) returns()
func (_CWB *CWBTransactorSession) NewGame(p1 common.Address, p2 common.Address, p1CardId *big.Int, p2CardId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.NewGame(&_CWB.TransactOpts, p1, p2, p1CardId, p2CardId)
}

// RemoveBet1 is a paid mutator transaction binding the contract method 0x4e4988be.
//
// Solidity: function removeBet1(uint256 _gameId) returns()
func (_CWB *CWBTransactor) RemoveBet1(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "removeBet1", _gameId)
}

// RemoveBet1 is a paid mutator transaction binding the contract method 0x4e4988be.
//
// Solidity: function removeBet1(uint256 _gameId) returns()
func (_CWB *CWBSession) RemoveBet1(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.RemoveBet1(&_CWB.TransactOpts, _gameId)
}

// RemoveBet1 is a paid mutator transaction binding the contract method 0x4e4988be.
//
// Solidity: function removeBet1(uint256 _gameId) returns()
func (_CWB *CWBTransactorSession) RemoveBet1(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.RemoveBet1(&_CWB.TransactOpts, _gameId)
}

// RemoveBet2 is a paid mutator transaction binding the contract method 0x84a76213.
//
// Solidity: function removeBet2(uint256 _gameId) returns()
func (_CWB *CWBTransactor) RemoveBet2(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _CWB.contract.Transact(opts, "removeBet2", _gameId)
}

// RemoveBet2 is a paid mutator transaction binding the contract method 0x84a76213.
//
// Solidity: function removeBet2(uint256 _gameId) returns()
func (_CWB *CWBSession) RemoveBet2(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.RemoveBet2(&_CWB.TransactOpts, _gameId)
}

// RemoveBet2 is a paid mutator transaction binding the contract method 0x84a76213.
//
// Solidity: function removeBet2(uint256 _gameId) returns()
func (_CWB *CWBTransactorSession) RemoveBet2(_gameId *big.Int) (*types.Transaction, error) {
	return _CWB.Contract.RemoveBet2(&_CWB.TransactOpts, _gameId)
}

// CWBBetMadeIterator is returned from FilterBetMade and is used to iterate over the raw logs and unpacked data for BetMade events raised by the CWB contract.
type CWBBetMadeIterator struct {
	Event *CWBBetMade // Event containing the contract specifics and raw log

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
func (it *CWBBetMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CWBBetMade)
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
		it.Event = new(CWBBetMade)
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
func (it *CWBBetMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CWBBetMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CWBBetMade represents a BetMade event raised by the CWB contract.
type CWBBetMade struct {
	GameId *big.Int
	Player common.Address
	CardId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetMade is a free log retrieval operation binding the contract event 0x942ce7d81758c4cf68fa7638d68d695ae1c785e10b2413ccca93098d14f4cf7a.
//
// Solidity: event BetMade(uint256 gameId, address indexed player, uint256 cardId)
func (_CWB *CWBFilterer) FilterBetMade(opts *bind.FilterOpts, player []common.Address) (*CWBBetMadeIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _CWB.contract.FilterLogs(opts, "BetMade", playerRule)
	if err != nil {
		return nil, err
	}
	return &CWBBetMadeIterator{contract: _CWB.contract, event: "BetMade", logs: logs, sub: sub}, nil
}

// WatchBetMade is a free log subscription operation binding the contract event 0x942ce7d81758c4cf68fa7638d68d695ae1c785e10b2413ccca93098d14f4cf7a.
//
// Solidity: event BetMade(uint256 gameId, address indexed player, uint256 cardId)
func (_CWB *CWBFilterer) WatchBetMade(opts *bind.WatchOpts, sink chan<- *CWBBetMade, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _CWB.contract.WatchLogs(opts, "BetMade", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CWBBetMade)
				if err := _CWB.contract.UnpackLog(event, "BetMade", log); err != nil {
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

// ParseBetMade is a log parse operation binding the contract event 0x942ce7d81758c4cf68fa7638d68d695ae1c785e10b2413ccca93098d14f4cf7a.
//
// Solidity: event BetMade(uint256 gameId, address indexed player, uint256 cardId)
func (_CWB *CWBFilterer) ParseBetMade(log types.Log) (*CWBBetMade, error) {
	event := new(CWBBetMade)
	if err := _CWB.contract.UnpackLog(event, "BetMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CWBEndGameIterator is returned from FilterEndGame and is used to iterate over the raw logs and unpacked data for EndGame events raised by the CWB contract.
type CWBEndGameIterator struct {
	Event *CWBEndGame // Event containing the contract specifics and raw log

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
func (it *CWBEndGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CWBEndGame)
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
		it.Event = new(CWBEndGame)
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
func (it *CWBEndGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CWBEndGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CWBEndGame represents a EndGame event raised by the CWB contract.
type CWBEndGame struct {
	GameId *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEndGame is a free log retrieval operation binding the contract event 0x84f8e1147ed860fdcb13c73d1ec1bf4aa9cded9250811ad1067ea2e81b61affd.
//
// Solidity: event EndGame(uint256 gameId, address indexed winner)
func (_CWB *CWBFilterer) FilterEndGame(opts *bind.FilterOpts, winner []common.Address) (*CWBEndGameIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CWB.contract.FilterLogs(opts, "EndGame", winnerRule)
	if err != nil {
		return nil, err
	}
	return &CWBEndGameIterator{contract: _CWB.contract, event: "EndGame", logs: logs, sub: sub}, nil
}

// WatchEndGame is a free log subscription operation binding the contract event 0x84f8e1147ed860fdcb13c73d1ec1bf4aa9cded9250811ad1067ea2e81b61affd.
//
// Solidity: event EndGame(uint256 gameId, address indexed winner)
func (_CWB *CWBFilterer) WatchEndGame(opts *bind.WatchOpts, sink chan<- *CWBEndGame, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CWB.contract.WatchLogs(opts, "EndGame", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CWBEndGame)
				if err := _CWB.contract.UnpackLog(event, "EndGame", log); err != nil {
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

// ParseEndGame is a log parse operation binding the contract event 0x84f8e1147ed860fdcb13c73d1ec1bf4aa9cded9250811ad1067ea2e81b61affd.
//
// Solidity: event EndGame(uint256 gameId, address indexed winner)
func (_CWB *CWBFilterer) ParseEndGame(log types.Log) (*CWBEndGame, error) {
	event := new(CWBEndGame)
	if err := _CWB.contract.UnpackLog(event, "EndGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CWBNewGameIterator is returned from FilterNewGame and is used to iterate over the raw logs and unpacked data for NewGame events raised by the CWB contract.
type CWBNewGameIterator struct {
	Event *CWBNewGame // Event containing the contract specifics and raw log

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
func (it *CWBNewGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CWBNewGame)
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
		it.Event = new(CWBNewGame)
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
func (it *CWBNewGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CWBNewGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CWBNewGame represents a NewGame event raised by the CWB contract.
type CWBNewGame struct {
	GameId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewGame is a free log retrieval operation binding the contract event 0xa0b7f6f22bb4f69adc52dd16ffe2a964ab8b577be274247fa9a1547b03ce3e55.
//
// Solidity: event NewGame(uint256 gameId)
func (_CWB *CWBFilterer) FilterNewGame(opts *bind.FilterOpts) (*CWBNewGameIterator, error) {

	logs, sub, err := _CWB.contract.FilterLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return &CWBNewGameIterator{contract: _CWB.contract, event: "NewGame", logs: logs, sub: sub}, nil
}

// WatchNewGame is a free log subscription operation binding the contract event 0xa0b7f6f22bb4f69adc52dd16ffe2a964ab8b577be274247fa9a1547b03ce3e55.
//
// Solidity: event NewGame(uint256 gameId)
func (_CWB *CWBFilterer) WatchNewGame(opts *bind.WatchOpts, sink chan<- *CWBNewGame) (event.Subscription, error) {

	logs, sub, err := _CWB.contract.WatchLogs(opts, "NewGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CWBNewGame)
				if err := _CWB.contract.UnpackLog(event, "NewGame", log); err != nil {
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

// ParseNewGame is a log parse operation binding the contract event 0xa0b7f6f22bb4f69adc52dd16ffe2a964ab8b577be274247fa9a1547b03ce3e55.
//
// Solidity: event NewGame(uint256 gameId)
func (_CWB *CWBFilterer) ParseNewGame(log types.Log) (*CWBNewGame, error) {
	event := new(CWBNewGame)
	if err := _CWB.contract.UnpackLog(event, "NewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
