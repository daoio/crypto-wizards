// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "./node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "./node_modules/@openzeppelin/contracts/utils/Strings.sol";
import "./interfaces/ICWD.sol";

/// @title CryptoWizardsCards
/// @notice contract for minting of cards collections
contract CryptoWizardsCard is ERC721URIStorage, Ownable {
    uint256 private collectionId;
    uint256 private tokenId;

    address public cryptoWizardsDeck;

    /// @notice Collection is a set of 4 unique cards
    struct Collection {
        uint256 collectionId;
        string baseURI;
    }

    /// @notice collection id => Collection
    mapping(uint256 => Collection) public collections;
    mapping(address => uint256) public _balances;
    mapping(uint256 => address) public _owners;
    mapping(uint256 => string) public _tokenURIs;
    mapping(uint256 => address) public _tokenApprovals;
    mapping(address => mapping(address => bool)) public _operatorApprovals;

    event CollectionMinted(address indexed owner, uint256 collectionId);

    constructor() ERC721("Crypto-wizards card", "CWC") {}

    /*
     ***************
     ** MODIFIERS **
     ***************
     */
    modifier onlyCWD() {
        require(
            msg.sender == cryptoWizardsDeck,
            "CryptoWizardsCard{onlyCWD}: access denied"
        );
        _;
    }

    /*
        *************
        ** GETTERS **
        *************
    */
    function _exists(uint256 _tokenId) internal view override returns (bool) {
        return _owners[_tokenId] != address(0);
    }

    function ownerOf(uint256 _tokenId) public view override returns (address) {
        address owner = _owners[_tokenId];
        require(owner != address(0), "ERC721: invalid token ID");
        return owner;
    }

    /*
     *************************
     ** COLLECTION CREATION **
     *************************
     */
    // https://nftstorage.link/ipfs/bafybeigr3plauwv4nz3wn2eliorfuzgwj2xeap7j36vbn2wdrmyvveun2i/1 +.json
    function setCollectionBaseURI(string memory _baseURI) external onlyOwner {
        collectionId++;
        Collection memory collection = Collection(collectionId, _baseURI);
        collections[collectionId] = collection;
    }

    /// @notice function should be called via deck contract
    function mintCollection(uint256 _collectionId, address _origin)
        external
        onlyCWD
        returns (uint256[4] memory ids)
    {
        // mint 4 cards
        for (uint256 i = 0; i < 4; i++) {
            tokenId++;
            _mint(_origin, tokenId);
            _setTokenURI(
                tokenId,
                _getTokenURI(
                    collections[_collectionId].baseURI, 
                    (i+1)
                )
            );
            ids[i] = tokenId;
        }

        /// @notice in the end `_origin` will have 4 cards in ownership, i.e. collection
        emit CollectionMinted(_origin, collectionId);

        return ids;
    }

    function _setTokenURI(uint256 _tokenId, string memory _tokenURI) internal override {
        require(_exists(_tokenId), "ERC721URIStorage: URI set of nonexistent token");
        _tokenURIs[_tokenId] = _tokenURI;
    }

    function approve(address _to, uint256 _tokenId) public override {
        address owner = ownerOf(_tokenId);
        require(_to != owner, "ERC721: approval to current owner");

        require(
            msg.sender == owner || isApprovedForAll(owner, msg.sender),
            "ERC721: approve caller is not token owner nor approved for all"
        );

        _approve(_to, _tokenId);
    }

     function transferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) public virtual override {
        require(_isApprovedOrOwner(msg.sender, _tokenId), "ERC721: caller is not token owner nor approved");

        _transfer(_from, _to, _tokenId);
    }

    /*
     ***************
     ** INTERNALS **
     ***************
     */
    function _getTokenURI(string memory _baseURI, uint256 _i) internal pure returns(string memory uri) {
        uri = string.concat(_baseURI, Strings.toString(_i), ".json");
    }

    function _transfer(
        address _from,
        address _to,
        uint256 _tokenId
    ) internal override {
        require(
            ownerOf(_tokenId) == _from,
            "ERC721: transfer from incorrect owner"
        );
        require(_to != address(0), "ERC721: transfer to the zero address");

        // Clear approvals from the previous owner
        _approve(address(0), _tokenId);

        _balances[_from] -= 1;
        _balances[_to] += 1;
        _owners[_tokenId] = _to;

        /// @notice remove card from `_from` deck
        ICWD(cryptoWizardsDeck).updateDeck(_from, 1, _tokenId);
        /// @notice add card to `_to` deck
        ICWD(cryptoWizardsDeck).updateDeck(_to, 0, _tokenId);

        emit Transfer(_from, _to, _tokenId);
    }

    function _setCWD(address _cwd) external onlyOwner {
        cryptoWizardsDeck = _cwd;
    }

    function _mint(address _to, uint256 _tokenId) internal override {
        require(_to != address(0), "ERC721: mint to the zero address");
        require(!_exists(_tokenId), "ERC721: token already minted");

        _balances[_to] += 1;
        _owners[_tokenId] = _to;

        emit Transfer(address(0), _to, _tokenId);
    }

    function _approve(address _to, uint256 _tokenId) internal override {
        _tokenApprovals[_tokenId] = _to;
        emit Approval(ownerOf(_tokenId), _to, _tokenId);
    }

    function _isApprovedOrOwner(address _spender, uint256 _tokenId) internal view override returns (bool) {
        address owner = ownerOf(_tokenId);
        return (_spender == owner || _tokenApprovals[_tokenId] == _spender);
    }
}
