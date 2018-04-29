# BookNBlock

BookNBlock soll ein Community-Marktplatz für Buchungen und Vermietungen von Unterkünften sein. Die Anwendung basiert auf einer Blockchain-Techologie.

## Installation
* [Node.js](https://nodejs.org/en/ "") muss installiert sein.
* [Angular Cli](https://cli.angular.io/ "") installieren: `npm install -g @angular/cli`
* Projektabhängigkeiten installieren: `npm install`
  * Sollte es hierbei zu Problemen mit node-gyp kommen dann müssen die [windows-build-tools](https://github.com/felixrieseberg/windows-build-tools "") als Administrator installiert werden: `npm install --global --production windows-build-tools`
  * Funktioniert es dann immer noch nicht, muss man seine Python Version überprüfen. node-gyp kommt nur mit Python 2.7 zu Recht: `npm config set python <path/to/python2.7>`
  * Jetzt funktioniert die Installation der Abhängigkeiten hoffentlich

## Starten des Servers
* Entwicklungsserver auf Port 4200 starten: `npm start`


## Smart Contract 
[Smart Contract](./doc/smartContract.md "")

## Links
- [mindmap](https://www.mindmeister.com/1074151893?t=uMmi6XTWvi)
