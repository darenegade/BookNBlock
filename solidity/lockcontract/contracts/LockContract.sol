pragma solidity ^0.4.17;

contract LockContract {
    string public message;

    function LockContract(string initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string newMessage) public {
        message = newMessage;
    }
}