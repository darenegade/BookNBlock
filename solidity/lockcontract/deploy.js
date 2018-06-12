const HDWalletProvider = require('truffle-hdwallet-provider');
const Web3 = require('web3');
const { interface, bytecode } = require('./compile');

const provider = new HDWalletProvider(
    'diet asthma equip loan jealous twist divorce cloth gym ramp stomach noise',
    ' https://rinkeby.infura.io/VhXic4UDRfv5w86p2hq7'
    );
const web3 = new Web3(provider);

const deploy = async () => {
    const accounts = await web3.eth.getAccounts();

    console.log('Attempting to deploy from account', accounts[0]);

    const result = await new web3.eth.Contract(JSON.parse(interface))
    .deploy({ data: '0x' + bytecode })
    .send({ gas: '3000000', from: accounts[0] });

    //you can check the deployed contract on rinkeby.etherscan.io
    //contract can be interacted with via remix IDE
    // 0xEe86D8d8163844517676C918556CDf42310c1671
    console.log('Contract deployed to', result.options.address);
};
deploy();