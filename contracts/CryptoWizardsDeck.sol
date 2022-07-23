// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "./node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/ICWC.sol";

/// @title CryptoWizardsDeck
/// @notice deck is an abstraction on top of CWC contract
/// to provide easy-to-use storage of ERC721 cards
contract CryptoWizardsDeck is Ownable {
    struct Deck {
        uint256 collectionId;
        string deckURI;
        /*
            tokenIds of cards
            if == 0 => card isn't in deck 
        */
        uint256 card1;
        uint256 card2;
        uint256 card3;
        uint256 card4;
    }
    address public cryptoWizardsCard;

    mapping(address => Deck) public decks;
    /// @notice collection id => deck URI
    mapping(uint256 => string) public deckURIs;
    /// @notice notify that user has deck
    mapping(address => bool) public hasDeck;

    modifier onlyCWC {
        require(msg.sender == cryptoWizardsCard, "CryptoWizardsDeck{onlyCWC}: access denied");
        _;
    }
    
    constructor(address _cwc) {
        cryptoWizardsCard = _cwc;
    }

    /*
        **********************
        ** GETTER FUNCTIONS **
        **********************
    */
    function getUserDeck(address _user) external view returns(Deck memory deck) {
        deck = decks[_user];
    }

    /*
        *******************
        ** DECK CREATION **
        *******************
    */
    function setDeckURI(uint256 collectionId, string memory deckURI) external onlyOwner {
        deckURIs[collectionId] = deckURI;
    }

    /// @notice deck creation from `_collectionId`
    function createDeck(uint256 _collectionId) external {
        require(!hasDeck[msg.sender], "CryptoWizardsDeck{createDeck}: msg.sender has a deck");
        /// @notice create 4 ERC721 card tokens
        (uint256[4] memory ids) = ICWC(cryptoWizardsCard).mintCollection(_collectionId, msg.sender);
        Deck memory deck = Deck(
            _collectionId,
            deckURIs[_collectionId],
            ids[0],
            ids[1],
            ids[2],
            ids[3]
        );
        decks[msg.sender] = deck;
        hasDeck[msg.sender] = true;
    }

    /// @param _select == 0 ? add card to deck : remove card from deck
    function updateDeck(address _owner, uint256 _select, uint256 _cardId) external onlyCWC {
        if (_select == 0) {
            _addCard(_owner, _cardId);
        } else {
            _removeCard(_owner, _cardId);
        }
    }

    /*
        ***************
        ** INTERNALS **
        ***************
    */
    function _addCard(address _owner, uint256 _newCardId) internal {
        Deck memory deck = decks[_owner];

        /// @notice add new card to the first empty card
        if (deck.card1 == 0) {
            decks[_owner].card1 = _newCardId;
        } else if (deck.card2 == 0) {
            decks[_owner].card2 = _newCardId;
        } else if (deck.card3 == 0) {
            decks[_owner].card3 = _newCardId;
        } else if (deck.card4 == 0) {
            decks[_owner].card4 = _newCardId;
        }
    }

    /// @param _cardId - card id to remove
    function _removeCard(address _owner, uint256 _cardId) internal {
        Deck memory deck = decks[_owner];

        if (deck.card1 == _cardId) {
            decks[_owner].card1 = 0;
        } else if (deck.card2 == _cardId) {
            decks[_owner].card2 = 0;
        } else if (deck.card3 == _cardId) {
            decks[_owner].card3 = 0;
        } else if (deck.card4 == _cardId) {
            decks[_owner].card4 = 0;
        }
    }
}