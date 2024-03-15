// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

contract Voting {
    mapping(string => uint256) public votes;

    function voteForColor(string memory color) public {
        votes["red"]=4;
    }

    function getVoteCount(string memory color) public view returns (uint256) {
        return votes[color];
    }
}