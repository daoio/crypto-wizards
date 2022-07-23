// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICWC {
    function mintCollection(uint256 _collectionId, address _origin) external returns(uint256[4] memory ids);
}