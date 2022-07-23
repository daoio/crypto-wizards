// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICWD {
    function updateDeck(address _owner, uint256 _select, uint256 _cardId) external;
}