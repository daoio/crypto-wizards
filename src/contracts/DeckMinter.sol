// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

///@title CryptoWizards
contract DeckMinter is ERC721URIStorage, Ownable {
    /// @notice Deck is a collection of 4 cards.
    struct Deck {
        uint256 deckId;
        string uri;

        /// @notice cards returns either card's NFT or address(0)
        IERC721 card1;
        IERC721 card2;
        IERC721 card3;
        IERC721 card4;
    }
    // deckId => Deck
    /// @notice represenets unique decks
    mapping(uint256 => Deck) public decks;
    /// @notice represents ERC721 deck tokens minted by users
    mapping(uint256 => Deck) public deckTokens;

    ///@notice deckId represents unique deck's id
    uint256 public deckId;
    ///@notice tokenId represents ids of ERC721 deck instances
    uint256 public tokenId;
    address public cardsMinter;

    modifier onlyCardsMinter {
        require(msg.sender == cardsMinter, "DeckMinter: Access denied");
        _;
    }

    constructor(address _cardsMinter) ERC721("Crypto-Wizards deck", "CWD") {
        cardsMinter = _cardsMinter;
    }

    /** 
        *************
        ** GETTERS **
        *************
    */
    ///@notice returns unique 4 cards collection's URI
    function getDeckURI(uint256 id) public view returns(string memory) {
        return decks[id].uri;
    }

    /** 
        ***********************
        ** DECK INTERACTIONS **
        ***********************
    */
    ///@notice mints new deck as an ERC721 token
    function mintDeck() external {
        deckId++;
        _mint(msg.sender, deckId);
        //_setDeck(deckId, _uri, _c1, _c2, _c3, _c4);
        decks[tokenId] = getDeckURI(deckId);
    }

    // deck : https://bafkreifmylhymbvnq72cb4imuhofykj6g75hfxll4743kunl4wnasjrva4.ipfs.nftstorage.link/
    function _addDeck(
        uint256 _id, 
        string memory _uri, 
        IERC721 _c1,
        IERC721 _c2, 
        IERC721 _c3, 
        IERC721 _c4
    ) external onlyCardsMinter returns(Deck memory) {
        Deck memory deck = Deck
        (
            _id, 
            _uri, 
            _c1, 
            _c2,
            _c3,
            _c4
        );
        decks[deck.deckId] = deck;

        return deck;
    }

    /// @notice Deck with 4 cards is added when that cards is minted
    function addCards(string memory uri) external onlyCardsMinter {
        deckId++;
        decks[deckId] = uri;
    }
}