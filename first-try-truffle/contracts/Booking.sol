pragma solidity ^0.4.0;

contract Booking {

    // ==== these are all state variables ===

    // the price of the flat for one night
    uint public price;
    // the ethereum address of the landlord
    address public landlord;
    // the ethereum address of the renter
    address public renter;
    // current state of a flat
    enum State {Booked, Free}
    State public state;

    //  ==== here comes the functions that modify the state varaiables ====

    // Ensure that `msg.value` is an even number.
    // Division will truncate if it is an odd number.
    // Check via multiplication that it wasn't an odd number.
    function Purchase() public payable {
        landlord = msg.sender;
        price = msg.value / 2;
        require((2 * price) == msg.value);
    }

    modifier condition(bool _condition) {
        require(_condition);
        _;
    }

    modifier onlyRenter() {
        require(msg.sender == renter);
        _;
    }

    modifier onlyLandlord() {
        require(msg.sender == landlord);
        _;
    }

    modifier inState(State _state) {
        require(state == _state);
        _;
    }

    event Aborted();
    event BookingConfirmed();
    event PayBooking();

    /// Abort the purchase and reclaim the ether.
    /// Can only be called by the seller before
    /// the contract is locked.
    function abort()
    public
    onlyLandlord
    inState(State.Booked)
    {
        emit Aborted();
        state = State.Free;
        landlord.transfer(this.balance);
    }

    /// Confirm the booking as landlord.
    /// Transaction has to include `2 * value` ether.
    /// The ether will be locked until confirmReceived
    /// is called.
    function confirmBooking()
    public
    inState(State.Free)
    condition(msg.value == (2 * value))
    payable
    {
        emit BookingConfirmed();
        landlord = msg.sender;
        state = State.Booked;
    }

    /// Pay the booking
    /// This will release the locked ether.
    function payBooking()
    public
    onlyRenter
    inState(State.Booked)
    {
        emit PayBooking();

        // NOTE: This actually allows both the landlord and the renter to
        // block the refund - the withdraw pattern should be used.

        renter.transfer(value);
        landlord.transfer(this.balance);
    }
}
