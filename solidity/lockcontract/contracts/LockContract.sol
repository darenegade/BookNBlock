pragma solidity ^0.4.23;

contract LockContract {

    struct Booking {
        uint offerID;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        string price;   //Fixed Numbers currently not supported: https://github.com/ethereum/solidity/issues/409
        string objectName;
        string ownerName;
        address owner;
    }

    Offer[] public offers;
    Booking[] public bookings;

    function insertOffer(string price, string objectName, string ownerName) public {
        
        Offer memory c;
        c.price = price;
        c.objectName = objectName;
        c.ownerName = ownerName;
        c.owner = msg.sender;
        offers.push(c);
    }

    function deleteOffer(uint offerID) public {

        require(offers.length > offerID);

        Offer storage offer = offers[offerID];
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

        for(uint i = 0; i < bookings.length; i++) {
            Booking storage b = bookings[i];
            if(b.offerID == offerID){
                require(b.checkIn > checkOut || b.checkOut < checkIn);
            }
        }

        bookings.push(Booking(offerID, checkIn, checkOut, msg.sender));
    }

}