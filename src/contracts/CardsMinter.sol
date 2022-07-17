// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

///@title CryptoWizards
///@notice the purpose of CardsMinter is obviously to mint new cards
contract CardsMinter is ERC721URIStorage, Ownable {
    struct Collection {
        string c1;
        string c2;
        string c3;
        string c4;
    }

    uint256 public collectionId;
    /// @notice each 4 cards collection has its own baseURI
    mapping(uint256 => string) public baseURIs;

    constructor() ERC721("Crypto-wizards card", "CWC") {}

    function mintCards(string memory ) external onlyOwner {
        collectionId++;
        _mint(address(this), collectionId);
        _setCollectionURIs()
    }

    function _setCollectionURIs(
        string memory _c1,
        string memory _c2,
        string memory _c3,
        string memory _c4
    ) internal {
        
    }

    /// https://bafybeiey4xcvfimu442tx2qlg2kv3cv27iwul752smgdogjgb562loeibm.ipfs.nftstorage.link/
    function setBaseURI(uint256 id, string memory uri) external onlyOwner {
        baseURIs[id] = uri;
    }
}