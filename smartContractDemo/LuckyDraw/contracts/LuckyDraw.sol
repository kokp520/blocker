// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract LuckyDraw {
    address public owner;
    uint256 public ticketPrice = 1 ether;
    uint256 public totalTicketsBought;
    uint256 public totalPrizeAmount;
    address public winner;
    bool public isDrawing;

    mapping(address => uint256) public ticketsBought;
    mapping(address => uint256) public tempParticipantMap;
    address[] public participants; // Array to store participant addresses just for temp

    event TicketPurchased(address indexed buyer, uint256 amount);
    event DrawCompleted(address indexed winner, uint256 prizeAmount);
    event MSG(address[] participants);

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "auth faild, only contract owner can call this");
        _;
    }

    modifier notDrawing() {
        require(!isDrawing, "drawing is in progress");
        _;
    }

    function buyTickets(uint256 _numberOfTickets) external payable notDrawing {
        require(msg.value >= ticketPrice * _numberOfTickets, "Insufficient funds");
    
        uint256 totalPrice = ticketPrice * _numberOfTickets;
        
        // payable(owner).transfer(totalPrice); // fee?
        
        ticketsBought[msg.sender] += _numberOfTickets;
        totalTicketsBought += _numberOfTickets;
        totalPrizeAmount += totalPrice;

        tempParticipantMap[msg.sender] = 1;

        participants.push(msg.sender);

        emit TicketPurchased(msg.sender, _numberOfTickets);
    }

    function startDrawing() external onlyOwner notDrawing {
        require(totalTicketsBought > 0, "No tickets bought yet");

        isDrawing = true;
        uint256 randomTicketNumber = uint256(keccak256(abi.encodePacked(block.timestamp, blockhash(block.number - 1)))) % totalTicketsBought;

        winner = getWinnerParticipant(randomTicketNumber);
        payable(winner).transfer(totalPrizeAmount); // Transfer the prize to the winner

        // Reset for next round
        totalTicketsBought = 0;
        totalPrizeAmount = 0;
        isDrawing = false;
        clearParticipants();

        emit DrawCompleted(winner, totalPrizeAmount);
    }

    function getWinnerParticipant(uint256 index) private view returns (address) {
        require(index < totalTicketsBought, "Over than participants");

        // If no participants found, return the first participant's address from the array
        if (participants.length == 0) {
            revert("No participants found");
        }

        return participants[index % participants.length];
    }

    function withdrawFunds(uint256 _amount) external onlyOwner {
        require(_amount <= address(this).balance, "Insufficient contract balance");
        payable(owner).transfer(_amount);
    }

    function clearParticipants() private {
        participants = new address[](0);
    }

    function echoParticipants() public {
        emit MSG(participants);
    }

    receive() external payable {}
}
