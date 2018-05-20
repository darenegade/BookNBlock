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
        .send({ from: accounts[0], gas: '1000000' })

    lockContract.setProvider(provider);
});

describe('lockContract', () => {
    it('deploys a contract', () => {
        assert.ok(lockContract.options.address);
    });

    /** 
    it('has default message Room uninitialized', async () => {
        const message = await lockContract.methods.message().call();
        assert.equal(message, 'Room uninitialized')
    });

    it('can change the message', async () => {
        await lockContract.methods.setMessage('Room out of order').send({ from: accounts[0]});
        const message = await lockContract.methods.message().call();
        assert.equal(message, 'Room out of order')
    });

    it('can be booked', async () => {
        await lockContract.methods.setBooked().send({ from: accounts[0]});
        const message = await lockContract.methods.message().call();
        assert.equal(message, 'Booked!')
    })

    it('can be set free', async () => {
        await lockContract.methods.setFree().send({ from: accounts[0]});
        const message = await lockContract.methods.message().call();
        assert.equal(message, 'Free!')
    })
    */

});