const assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require('web3');

const provider = ganache.provider();
const web3 = new Web3(provider);

const { interface, bytecode } = require('../compile');
const GAS = '3000000'

require('events').EventEmitter.prototype._maxListeners = 100;

let accounts;
let lockContract;

beforeEach(async () => {
    // Get a list of all accounts
    accounts = await web3.eth.getAccounts();

    //Use one of those accounts to deploy the contract
    lockContract = await new web3.eth.Contract(JSON.parse(interface))
        .deploy({ data: bytecode})
        .send({ from: accounts[0], gas: GAS})

    lockContract.setProvider(provider);
});

describe('lockContract', () => {
    it('deploys a contract', () => {
        assert.ok(lockContract.options.address);
    });

    it('offer can be created', async () => {

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["OfferSaved"], undefined);
                assert.equal(tx.events["OfferSaved"].returnValues.offerID, 0);
              });

            let offer = await lockContract.methods.getOffer(0).call();
            assert.equal( offer[0], 1)
    })

    it('offer can be deleted', async () => {

        let id = -1; 

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                 id = tx.events["OfferSaved"].returnValues.offerID
            });

            assert.equal(id, 0)

            await lockContract.methods
            .deleteOffer(id)
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["OfferDeleted"], undefined);
                assert.equal(tx.events["OfferDeleted"].returnValues.offerID, id);
              })
    })

    it('offer can be rented', async () => {

        let id = -1; 

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                 id = tx.events["OfferSaved"].returnValues.offerID
            });

            assert.equal(id, 0)

            await lockContract.methods
            .rentAnOffer(
                id,
                1520035200,
                1520121600
            )
            .send({value: 1, from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["BookingAccepted"], undefined);
                assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 0);
              })

              await lockContract.methods
              .rentAnOffer(
                  id,
                  1520121601,
                  1520121700
              )
              .send({value: 1, from: accounts[0], gas: GAS})
              .then(function (tx) {
                  assert.notEqual(tx.events["BookingAccepted"], undefined);
                  assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 1);
                })
    })

    it('free offer can be seen', async () => {
        
        let id = -1; 

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.equal(tx.events["OfferSaved"].returnValues.offerID, 0);
                id = tx.events["OfferSaved"].returnValues.offerID
            });

        let offers = await lockContract.methods
            .getOfferIDs()
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(offers, [0]);

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer 2',
                'Teststraße 12, München', 
                'Hans', 
                'Not so Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.equal(tx.events["OfferSaved"].returnValues.offerID, 1);
            });

        await lockContract.methods
            .rentAnOffer(
                0,
                1520035200,
                1520121600
            )
            .send({value: 1, from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["BookingAccepted"], undefined);
                assert.equal(tx.events["BookingAccepted"].returnValues.offerID, 0);
                assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 0);
                })

        let freeOffers = await lockContract.methods
            .getFreeOfferIDs(
                1520121601,
                1520121700
            )
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(freeOffers, [0,1]);

        freeOffers = await lockContract.methods
            .getFreeOfferIDs(
                1520035100,
                1520035201
            )
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(freeOffers, [1]);
            
    })

    it('booking can allow tenant', async () => {

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.equal(tx.events["OfferSaved"].returnValues.offerID, 0);
            });

        await lockContract.methods
            .rentAnOffer(
                0,
                1520035200,
                1520121600
            )
            .send({value: 1, from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["BookingAccepted"], undefined);
                assert.equal(tx.events["BookingAccepted"].returnValues.offerID, 0);
                assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 0);
                })

        let allowed = await lockContract.methods
            .isAllowedAt(
                0,
                accounts[0],
                1520035205
            )
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(allowed, true);

        allowed = await lockContract.methods
            .isAllowedAt(
                0,
                accounts[0],
                1520121605
            )
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(allowed, false);
            
    })

    it('can see own bookings', async () => {

        await lockContract.methods
            . insertOffer(
                1, 
                'Cool Flat Offer',
                'Teststraße 1, München', 
                'Hans', 
                'Very Cool Flat', 
                accounts[0], 
                1514764800, 
                1546214400
                )
            .send({ from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.equal(tx.events["OfferSaved"].returnValues.offerID, 0);
            });

        await lockContract.methods
            .rentAnOffer(
                0,
                1520035200,
                1520121600
            )
            .send({value: 1, from: accounts[0], gas: GAS})
            .then(function (tx) {
                assert.notEqual(tx.events["BookingAccepted"], undefined);
                assert.equal(tx.events["BookingAccepted"].returnValues.offerID, 0);
                assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 0);
                })

                await lockContract.methods
                .rentAnOffer(
                    0,
                    1520121700,
                    1520122600
                )
                .send({value: 1, from: accounts[1], gas: GAS})
                .then(function (tx) {
                    assert.notEqual(tx.events["BookingAccepted"], undefined);
                    assert.equal(tx.events["BookingAccepted"].returnValues.offerID, 0);
                    assert.equal(tx.events["BookingAccepted"].returnValues.bookingID, 1);
                    })
                
        let ownBookings = await lockContract.methods
            .getOwnBookingIDs( )
            .call({from: accounts[0], gas: GAS})

        assert.deepEqual(ownBookings, [0]);
            
    })

});