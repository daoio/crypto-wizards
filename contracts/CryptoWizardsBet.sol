//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "./node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "./node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol";

/// @title CryptoWizardsBet
/// @notice contains functionality for NFT betting
contract CryptoWizardsBet is ReentrancyGuard {
    uint256 private gameId;
    /// @notice address that is allowed to start and end player's games
    address public wizard;
    address public cryptoWizardsCards;

    /// @notice track parties actions
    struct Player {
        address addr;
        bool betMade;
        // bet card's id
        uint256 cardId;
    }

    struct Game {
        uint256 gameId;
        Player player1;
        Player player2;
        address winner;
        bool gameEnded;
    }

    /// @notice track each game by its id
    mapping(uint256 => Game) public games;

    event NewGame(uint256 gameId);
    event BetMade(uint256 gameId, address indexed player, uint256 cardId);
    event EndGame(uint256 gameId, address indexed winner);

    constructor() {
        /// @notice for simplicity it's contract deployer
        wizard = msg.sender;
    }

    /*
     ***************
     ** MODIFIERS **
     ***************
     */
    modifier onlyWizard {
        require(
            msg.sender == wizard,
            "CryptoWizards{onlyStarter}: access denied"
        );
        _;
    }

    modifier betReview(uint256 _select, uint256 _gameId) {
        if (_select == 1) {
            // player 1 bet
            require(
                IERC721(cryptoWizardsCards).ownerOf(games[_gameId].player1.cardId) == msg.sender &&
                msg.sender == games[_gameId].player1.addr &&
                games[_gameId].player1.betMade == false &&
                games[_gameId].gameEnded == false,
                "CryptoWizardsBet{betReview}: invalid conditions"
            );
        } else {
            // player 2 bet
            require(
                IERC721(cryptoWizardsCards).ownerOf(games[_gameId].player2.cardId) == msg.sender &&
                msg.sender == games[_gameId].player2.addr &&
                games[_gameId].player2.betMade == false &&
                games[_gameId].gameEnded == false,
                "CryptoWizardsBet{betReview}: invalid conditions"
            );
        }
        _;
    }

    modifier betRemovalReview(uint256 _select, uint256 _gameId) {
        if (_select == 1) {
            // player 1 remove bet
            require(
                msg.sender == games[_gameId].player1.addr &&
                    games[_gameId].player1.betMade == true &&
                    games[_gameId].player2.betMade == false &&
                    games[_gameId].gameEnded == false,
                "CryptoWizardsBet{removeBetReview}: invalid conditions"
            );
        } else {
            // player 2 remove bet
            require(
                msg.sender == games[_gameId].player2.addr &&
                    games[_gameId].player2.betMade &&
                    games[_gameId].player1.betMade == false &&
                    games[_gameId].gameEnded == false,
                "CryptoWizardsBet{removeBetReview}: invalid conditions"
            );
        }
        _;
    }

    modifier endGameReview(uint256 _gameId) {
        require(
            games[_gameId].player1.betMade &&
            games[_gameId].player2.betMade &&
            games[_gameId].gameEnded == false
        );
        _;
    }

    /*
     ***********
     ** START **
     ***********
     */
    function newGame(
        address p1,
        address p2,
        uint256 p1CardId,
        uint256 p2CardId
    ) external onlyWizard {
        require(
            IERC721(cryptoWizardsCards).ownerOf(p1CardId) == p1 &&
            IERC721(cryptoWizardsCards).ownerOf(p2CardId) == p2
        );
        Player memory player1 = Player(p1, false, p1CardId);
        Player memory player2 = Player(p2, false, p2CardId);
        gameId++;
        Game memory game = Game(gameId, player1, player2, address(0), false);
        games[gameId] = game;

        emit NewGame(gameId);
    }

    /*
     *************
     ** BETTING **
     *************
     */
    /// @notice player1 bet
    function bet1(uint256 _gameId) external betReview(1, _gameId) nonReentrant {
        /// approve first
        /// @notice transfer card that was bet in the newGame
        IERC721(cryptoWizardsCards).transferFrom(
            msg.sender,
            address(this),
            games[_gameId].player1.cardId
        );
        games[_gameId].player1.betMade = true;

        emit BetMade(_gameId, msg.sender, games[_gameId].player1.cardId);
    }

    /// @notice player2 bet
    function bet2(uint256 _gameId) external betReview(2, _gameId) nonReentrant {
        /// approve first
        /// @notice transfer card that was bet in the newGame
        IERC721(cryptoWizardsCards).transferFrom(
            msg.sender,
            address(this),
            games[_gameId].player2.cardId
        );
        games[_gameId].player1.betMade = true;

        emit BetMade(_gameId, msg.sender, games[_gameId].player2.cardId);
    }

    /**
    *   @notice bet removal allowed only if one of the players hasn't made bet yet
    *   removeBet() removes certain player's bet, transfers locked cards
    *   to respective players and ends game
    */
    function removeBet1(uint256 _gameId) external betRemovalReview(1, _gameId) nonReentrant {
        IERC721(cryptoWizardsCards).transferFrom(
            address(this),
            msg.sender,
            games[_gameId].player1.cardId
        );
        IERC721(cryptoWizardsCards).transferFrom(
            address(this),
            games[_gameId].player2.addr,
            games[_gameId].player2.cardId
        );

        games[_gameId].gameEnded = true;
        games[_gameId].player1.betMade = false;
        games[_gameId].player2.betMade = false;
    }

    function removeBet2(uint256 _gameId) external betRemovalReview(2, _gameId) nonReentrant {
         IERC721(cryptoWizardsCards).transferFrom(
            address(this),
            msg.sender,
            games[_gameId].player1.cardId
        );
        IERC721(cryptoWizardsCards).transferFrom(
            address(this),
            games[_gameId].player2.addr,
            games[_gameId].player2.cardId
        );

        games[_gameId].gameEnded = true;
        games[_gameId].player1.betMade = false;
        games[_gameId].player2.betMade = false;
    }

    function endGame(address _winner, uint256 _gameId)
        public
        onlyWizard
        endGameReview(_gameId)
    {
        uint256[] memory cardIds = new uint256[](2);

        cardIds[0] = games[_gameId].player1.cardId;
        cardIds[1] = games[_gameId].player2.cardId;

        /// @notice transfer both locked cards to winner
        _transferCards(cardIds, address(this), _winner);

        games[_gameId].gameEnded = true;
        games[_gameId].winner = _winner;

        emit EndGame(_gameId, _winner);
    }

    /*
     ***************
     ** INTERNALS **
     ***************
     */
    function _transferCards(
        uint256[] memory cardIds,
        address _from,
        address _to
    ) internal {
        for (uint256 i = 0; i < cardIds.length; i++) {
            IERC721(cryptoWizardsCards).transferFrom(_from, _to, cardIds[i]);
        }
    }

    function _setCWC(address _cwc) external onlyWizard {
        cryptoWizardsCards = _cwc;
    }
}
