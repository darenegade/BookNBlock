pragma solidity ^0.4.23;

contract LockContract {

    struct Booking {
        uint offerID;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        uint price;   //Price in Cent - Fixed Numbers currently not supported: https://github.com/ethereum/solidity/issues/409
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

    modifier offerAvailable(uint offerID) {
        require(offers.length > offerID);
        _;
    }

    modifier onlyOwner(uint offerID) {
        require(offers[offerID].owner == msg.sender);
        _;
    }

    function insertOffer(
        uint price, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
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

    function updateOffer(
        uint offerID, uint price, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
        )
         public 
         offerAvailable(offerID)
         onlyOwner(offerID){
        
        require(validFrom < validUntil);

        Offer storage offer = offers[offerID];
        offer.price = price;
        offer.objectName = objectName;
        offer.objectAddress = objectAddress;
        offer.ownerName = ownerName;
        offer.description = description;
        offer.door = door;
        offer.validFrom = validFrom;
        offer.validUntil = validUntil;
    }

    function deleteOffer(uint offerID) 
        public 
        offerAvailable(offerID)
        onlyOwner(offerID) {

        for (uint i = offerID; i<offers.length-1; i++) {
            offers[i] = offers[i+1];
        }
        delete offers[offers.length-1];
        offers.length--;
    }

    function rentAnOffer(uint offerID,  uint256 checkIn, uint256 checkOut) 
        public 
        offerAvailable(offerID) {

        require(checkIn < checkOut);

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