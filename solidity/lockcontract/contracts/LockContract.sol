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
        string objectAddress;
        string ownerName;
        string description;
        uint256 validFrom;
        uint256 validUntil;
        address owner;
        address door;
        uint[] bookingIndexes;
    }

    Offer[] public offers;
    Booking[] public bookings;

    function insertOffer(
        string price, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
        ) public {
        
        require(validFrom < validUntil);

        Offer memory c;
        c.price = price;
        c.objectName = objectName;
        c.objectAddress = objectAddress;
        c.ownerName = ownerName;
        c.description = description;
        c.owner = msg.sender;
        c.door = door;
        c.validFrom = validFrom;
        c.validUntil = validUntil;
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

        Offer storage offer = offers[offerID];
        require(checkIn >= offer.validFrom && checkOut <= offer.validUntil);

        uint[] storage bookingIndexes = offer.bookingIndexes;

        for(uint i = 0; i < bookingIndexes.length; i++) {
            Booking storage b = bookings[bookingIndexes[i]];
            if(b.offerID == offerID){
                require(b.checkIn > checkOut || b.checkOut < checkIn);
            }
        }

        bookingIndexes.push(bookings.length);
        bookings.push(Booking(offerID, checkIn, checkOut, msg.sender));
    }

}