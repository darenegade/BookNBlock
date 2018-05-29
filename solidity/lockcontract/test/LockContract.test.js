const assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require('web3');

const provider = ganache.provider();
const web3 = new Web3(provider);

const { interface, bytecode } = require('../compile');

let accounts;
let lockContract;

beforeEach(async () => {
    // Get a list of all accounts
    accounts = await web3.eth.getAccounts();

    //Use one of those accounts to deploy the contract
    lockContract = await new web3.eth.Contract(JSON.parse(interface))
        .deploy({ data: bytecode})
        .send({ from: accounts[0], gas: '2000000' })

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
            .send({ from: accounts[0], gas: '2000000'})
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
            .send({ from: accounts[0], gas: '2000000'})
            .then(function (tx) {
                 id = tx.events["OfferSaved"].returnValues.offerID
            });

            assert.equal(id, 0)

            await lockContract.methods
            .deleteOffer(id)
            .send({ from: accounts[0], gas: '2000000'})
            .then(function (tx) {
                assert.notEqual(tx.events["OfferDeleted"], undefined);
                assert.equal(tx.events["OfferDeleted"].returnValues.offerID, id);
              })
    })

});