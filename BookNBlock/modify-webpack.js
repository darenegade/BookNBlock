var fs = require('fs')
var webpackFile = 'node_modules/@angular/cli/models/webpack-configs/browser.js';
fs.readFile(webpackFile, 'utf8', function (err,data) {
  if (err) {
    return console.log(err);
  } else {
    var result = data.replace("crypto: 'empty'", 'crypto: true');

    try {
        fs.copyFileSync(webpackFile, webpackFile + '.bak', fs.constants.COPYFILE_EXCL);
    } catch (err) {}

    fs.writeFile(webpackFile, result, 'utf8', function (err) {
        if (err) return console.log(err);
    });
  }
});