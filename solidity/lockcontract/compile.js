const path = require('path');
const fs = require('fs');
const solc = require('solc');

const lockContractPath = path.resolve(__dirname, 'contracts', 'LockContract.sol');
const source = fs.readFileSync(lockContractPath, 'utf8');

module.exports = solc.compile(source, 1).contracts[':LockContract'];

