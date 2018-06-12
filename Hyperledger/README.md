![alt text](https://www.hyperledger.org/wp-content/uploads/2016/09/logo_hl_new.png "BookAndBlock powered by Hyperledger https://hyperledger.com")
# Hyperledger Fabric
Dieses Projekt von der Hochschule München wird in der Informatik Masterveranstaltung Hauptseminar umgesetzt. Es wird eine Blockchain mittels Hyperledger Fabric für BookAndBlock entwickelt.    

## Installation

### Vorbereitung

* Es sollte ein [Ubuntu 16.04](releases.ubuntu.com/16.04/) verwendet werden.
* Installieren Sie [cURL](https://curl.haxx.se/)
* Installieren Sie [Go](https://golang.org/) in der Version >= 1.9
* Installieren Sie die [Communitiy Edition](https://docs.docker.com/install/linux/docker-ce/ubuntu/) von [Docker](https://www.docker.com/) und Docker-Compose
* Installieren Sie [Node.js](https://nodejs.org) in der Version 9.11.1 und npm in der Version 6.0.0

### Troubleshoot
Im Nachfolgenden soll ein Troubelshoot bei eventuell auftretenden Fehlern helfen.

#### Allgemein

---
Bei **Permission denied** Fehlern die Rechte zurücksetzten.

---

Letzte version von Node.js verwenden.

---

#### Docker
---
Bei dem Fehler

```
rpc error: code = 13 desc = invalid header field value "oci runtime error: exec failed: container_linux.go:1153: sending signal 0 to pid 10628 caused \"permission denied\ "\n"
```
fürhen Sie folgendes aus
```
sudo usermod -a -G docker $USER
```
und anschließend den **Rechner neu starten**.

---
**Fehlt** beim Bauen der Anwendung der **Fabric-CA-Container**. Prüfbar durch den Befehl
```
docker ps
```
sollte in dem Ordner **Main-Hyperledger/base/docker-compose-base.yaml** die Einstellung von **FABRIC_CA_SERVER_CA_KEYFILE** wie folgt geändert werden
```
 FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/STRING_TO_REPLACE
```
Anschließend die Anwendung **löschen** und **neu erstellen**.
```
./bookandblockhyperledgerstartscript.sh destroy
[...]
./bookandblockhyperledgerstartscript.sh build
```

---


### Erstellen, starten, stopen und löschen der Anwendung

Das Skript wie folgt beschrieben in dem Ordner Main-Hyperledger ausführen

| Aktion       | Befehl       |
| -------------|:-------------:
| erstellen    |build         |
| starten      |start         |
| stopen       |stop          |
| löschen      |destroy       |


```
./bookandblockhyperledgerstartscript.sh [Befehl]
```


## Schnittstellen zur Erprobung

Die zentrale Verwaltung des Netzwerks erfolgt über die CLI. Um in die CLI zu gelangen, folgendes Befehl ausführen.
```
docker exec -it cli bash
```
Weitere Schnittstellen sind in der Tabelle aufgelistet.

| Schnittstelle      | Port                        | Path                        |Info                        |
| ------------------|:--------------|:--------------|:--------------|
| Rest                                |3000                        |/explorer                |Gilt als Wrapper für Node.js (könnte als generische API ausgebaut werden)|
|CouchDB                        |5984, 6984, 7984|/_utils                |Anzeigen der Blockchain Inhalte, User: bookandblock Passwort: bookandblock|
|FabricCA                        |7054                        |/api/v1                |Zertifikatverwaltung, siehe [swagger-fabric-ca.json](http://github.com/hyperledger/fabric-ca/blob/master/swagger/swagger-fabric-ca.json) Für direkt ansicht hier einfügen [Swagger](http://editor2.swagger.io/#!/)
|Peer0                                |7051                        |                                |Zum debuggen und/oder für die Entwicklung
|Peer1                                |8051                        |                                |Zum debuggen und/oder für die Entwicklung
|Peer2                                |9051                        |                                |Zum debuggen und/oder für die Entwicklung
|Orderer                        |7050                        |                                |Zum debuggen und/oder für die Entwicklung


## Nötige Anpassungen
Folgende anpassungen sollten für die Türe und das Front-End durchgeführt werden.

### Tür
Alle Dateien für die Tür sind im Ordner Door-API.

#### Client.js
* Eine Nachricht wird mit [rpc-websockets](https://github.com/qaap/rpc-websockets) an den Server gesendet.

#### server.js
* Server ließt den Public Key aus der BlockChain und vergleicht es mit der Signierten-Nachricht. Anschließend wird entschieden ob sich die Türe öffnet oder nicht. In einer zukünftigen Iteration des Projekts könnten weitere Sicherheitsprüfungen durchgeführt werden (Nonce, TTL der Nachricht etc.).

### Front-End
Alle Dateien für das Front-End sind im Ordner Frontend-API.

#### enrollAdmin.js
* Den richtigen AdminUser setzen (bookandblock)
#### registerUser.js
* Den richtigen User setzen (bookandblockgenericuser)
#### invoke.js
* Den richtigen User setzen (bookandblockgenericuser)
* Die Parementer für den Request
        * Funktionsname
        * Argumente für die Funktion (siehe bookandblockcc.go)
#### query.js
* Den richtigen User setzen (bookandblockgenericuser)
* Die Parementer für den Request
        * Funktionsname
        * Argumente für die Funktion (siehe bookandblockcc.go)

## ReST zugriff auf die Blockchain (wip)
Alternativ kann die Restschnittstelle genutzt werden, um mit der Blockchain zu kommunizieren. Hierfür die Befehle im Skript unter der funktion getRestService() erstmal manuell ausführen.

## Aktuelle Features (wip)

|Bezeichnung                                                | Chaincode funktion                |
| ----------------------------------|:--------------------------|
|Angebot anlegen                                        |insertOffer                                |
|Angebot löschen                                        |delete                                                |
|Angebot über seine Id finden                |getOffer                                        |
|Angebote über Public Key suchen        |queryOffersByPk                        |
|Anzeigen der Historie                                |getHistoryForOffer                        |


## Geplante Features (wip)

* Angebot übertragen
* Angebot mieten
* Optimierung
* ABAC (Attribute-Based-Access-Control)
* Kafka / ZooKeeper


## Quellen

* [Hyperledger Fabric](http://hyperledger-fabric.readthedocs.io) bietet eine Reihe von Tutorials für den Einstieg.
        * [fabcar](https://github.com/hyperledger/fabric-samples/tree/release-1.1/fabcar)
        * [marble](https://github.com/hyperledger/fabric-samples/tree/release-1.1/chaincode/marbles02)
        * Hilfreich für das Verständis könnten auch die anderen Beispiele in [fabric-samples](https://github.com/hyperledger/fabric-samples) sein.
* Hilfreiche Videos zu Hyperledger von Иван Ванков 
        * [Hyperledger Fabric - build first network](https://www.youtube.com/watch?v=MPNkUqOKhVE&list=PLjsqymUqgpSTGC4L6ULHCB_Mqmy43OcIh)
        * [Hyperledger - Blockchain Technologies for Business](https://www.youtube.com/watch?v=7EpPrSJtqZU&list=PLjsqymUqgpSRXC9ywNIVUUoGXelQa4olO)

## Betreuer
* [**Michael Schäfer**](https://www.cs.hm.edu/die_fakultaet/ansprechpartner/lehrbeauftragte/lba_liste_iframe_zpa.de.html) 


## Autoren

* [**Frank Christian Geyer**](https://github.com/frankchrisg)
* [**Deniz Mardin**](https://github.com/dmardin)
