pragma solidity ^0.4.17;

contract LockContract {

    struct Booking {
        uint bookingId;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        uint offerId;
        string price;   //Fixed Numbers currently not supported: https://github.com/ethereum/solidity/issues/409
        string objectName;
        string ownerName;
        address owner;
        Booking[] bookings;
    }

    Offer[] offers;

    constructor() public {
    }

    

}