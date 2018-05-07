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
        .deploy({ data: bytecode, arguments: ['Room locked!']})
        .send({ from: accounts[0], gas: '1000000' })

    lockContract.setProvider(provider);
});

describe('lockContract', () => {
    it('deploys a contract', () => {
        assert.ok(lockContract.options.address);
    });

    it('has default message Room locked', async () => {
        const message = await lockContract.methods.message().call();
        assert.equal(message, 'Room locked!')
    });
});