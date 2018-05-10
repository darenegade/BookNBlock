var WebSocketServer = require('rpc-websockets').Server

var forge = require('node-forge');
var crypto = require('crypto');

var server = new WebSocketServer({
  port: 9999,
  host: 'localhost'
})

// modify this in order to get a modular function call
var tools = require('../query');

server.register('openDoor', function() {
var arr = Object.values(arguments[0]);

/***********************************************************************
************************************************************************
************************************************************************

See Frontend-API for further information

*/

tools.query(arguments)

/***********************************************************************
************************************************************************
************************************************************************

*/

})
