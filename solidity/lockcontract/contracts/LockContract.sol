pragma solidity ^0.4.23;

contract LockContract {

    struct Booking {
        uint offerID;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        uint priceInWei;
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
        require(
            offers.length > offerID,
            "Offer not found"
        );
        _;
    }

    modifier onlyOwner(uint offerID) {
        require(
            offers[offerID].owner == msg.sender,
            "Only Owner of Offer is alowed"
        );
        _;
    }

    modifier costs(uint offerID) {
        uint priceInWei = offers[offerID].priceInWei;
        require(
            msg.value >= priceInWei,
            "Not enough Wei provided."
        );
        _;
        if (msg.value > priceInWei)
            msg.sender.transfer(msg.value - priceInWei);
    }

    event OfferSaved();
    event OfferDeleted();
    event BookingAccepted();

    function insertOffer(
        uint priceInWei, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
        ) public {
        
        require(validFrom < validUntil);

        Offer memory c;
        c.priceInWei = priceInWei;
        c.objectName = objectName;
        c.objectAddress = objectAddress;
        c.ownerName = ownerName;
        c.description = description;
        c.owner = msg.sender;
        c.door = door;
        c.validFrom = validFrom;
        c.validUntil = validUntil;
        offers.push(c);

        emit OfferSaved();
    }

    function updateOffer(
        uint offerID, uint priceInWei, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
        )
         public 
         offerAvailable(offerID)
         onlyOwner(offerID){
        
        require(validFrom < validUntil);

        Offer storage offer = offers[offerID];
        offer.priceInWei = priceInWei;
        offer.objectName = objectName;
        offer.objectAddress = objectAddress;
        offer.ownerName = ownerName;
        offer.description = description;
        offer.door = door;
        offer.validFrom = validFrom;
        offer.validUntil = validUntil;

        emit OfferSaved();
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

        emit OfferDeleted();
    }

    function rentAnOffer(uint offerID,  uint256 checkIn, uint256 checkOut) 
        public
        payable 
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
        offer.owner.transfer(offer.priceInWei);

        emit BookingAccepted();
    }

}