pragma solidity ^0.4.23;

contract LockContract {

    struct Booking {
        uint offerID;
        uint256 checkIn;
        uint256 checkOut;
        address tenant;
    }

    struct Offer{
        uint index;
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

    mapping(uint => Offer) public offers;
    uint[] public offerIDs;
    uint public nextID;
    Booking[] public bookings;

    modifier offerAvailable(uint offerID) {
        require(
            offerID >= 0 && offerIDs[offers[offerID].index] == offerID,
            "Offer not found"
        );
        _;
    }

    modifier bookingAvailable(uint bookingID) {
        require(
            bookingID >= 0 && bookingID < bookings.length,
            "Booking not found"
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

    event OfferSaved(uint offerID);
    event OfferDeleted(uint offerID);
    event BookingAccepted(uint bookingID);

    constructor() public {

    }

    function insertOffer(
        uint priceInWei, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
        ) public {
        
        require(validFrom < validUntil);

        uint newOfferID = getNextID();

        Offer memory c;
        c.index = offerIDs.length;
        c.priceInWei = priceInWei;
        c.objectName = objectName;
        c.objectAddress = objectAddress;
        c.ownerName = ownerName;
        c.description = description;
        c.owner = msg.sender;
        c.door = door;
        c.validFrom = validFrom;
        c.validUntil = validUntil;
        offers[newOfferID] = c;

        offerIDs.push(newOfferID);

        emit OfferSaved(newOfferID);
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

        emit OfferSaved(offerID);
    }

    function deleteOffer(uint offerID) 
        public 
        offerAvailable(offerID)
        onlyOwner(offerID) {

        uint rowToDelete = offers[offerID].index;
        uint keyToMove = offerIDs[offerIDs.length-1];
        offerIDs[rowToDelete] = keyToMove;
        offers[keyToMove].index = rowToDelete; 
        offerIDs.length--;

        emit OfferDeleted(offerID);
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

        emit BookingAccepted(bookings.length - 1);
    }

    function getOffer(uint offerID) public view offerAvailable(offerID) 
        returns (
            uint priceInWei, string objectName, string objectAddress, string ownerName, string description, address door, uint256 validFrom, uint256 validUntil
            ){
        Offer storage offer = offers[offerID];
        return (
            offer.priceInWei,
            offer.objectName,
            offer.objectAddress,
            offer.ownerName,
            offer.description,
            offer.door,
            offer.validFrom,
            offer.validUntil
            );
    }

    function getBooking(uint bookingID) public view bookingAvailable(bookingID) 
        returns (uint offerID, uint256 checkIn, uint256 checkOut) {
        Booking storage booking = bookings[bookingID];
        return (
            booking.offerID,
            booking.checkIn,
            booking.checkOut
            );
    }

    function getOfferIDs() public view returns(uint[] offerIDs) {
        return offerIDs;
    }

    function getBookingIDsForOffer(uint offerID) public view offerAvailable(offerID) returns(uint[] offerIDs) {
        Offer storage offer = offers[offerID];
        return offer.bookingIndexes;
    }

    function getOffersLength() public view returns(uint length) {
        return offerIDs.length;
    }

    function getNextID() private returns(uint nextID) {
        return nextID++;
    }
}