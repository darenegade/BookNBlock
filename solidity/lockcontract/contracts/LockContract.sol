pragma solidity ^0.4.17;

contract LockContract {

    struct Booking {
        uint bookingID;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        uint offerID;
        string price;   //Fixed Numbers currently not supported: https://github.com/ethereum/solidity/issues/409
        string objectName;
        string ownerName;
        address owner;
        Booking[] bookings;
    }

    Offer[] offers;

    constructor() public {
        offers = new Offer[](10);
    }

    function insertOffer(string price, string objectName, string ownerName) public {
        Booking[] memory bookingInit = new Booking[](10);
        offers.push(Offer(offers.length + 1, price, objectName, ownerName, msg.sender, bookingInit));
    }

    function deleteOffer(uint offerID) public {

        require(offers.length > offerID);

        Offer offer = offers[offerID];
        require(offer.owner == msg.sender);

        for (uint i = offerID; i<offers.length-1; i++) {
            offers[i] = offers[i+1];
        }
        delete offers[offers.length-1];
        offers.length--;
    }

    function rentAnOffer(uint offerID,  uint256 checkIn, uint256 checkOut) public {

        require(checkIn < checkOut);

        require(offers.length > offerID);
        Offer offer = offers[offerID];        

        for(uint i = 0; i<offer.bookings.length; i++) {
            require(offer.bookings[i].checkIn > checkOut || offer.bookings[i].checkOut < checkIn);
        }

        offer.bookings.push(Booking(offer.bookings.length, checkIn, checkOut, msg.sender));
    }

}