#!/bin/sh

#  bookandblockhyperledgerstartscript.sh
#
#  Created on 28.04.18.
#  


#
#peer chaincode invoke -o orderer.bookandblock.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem -C bookchannel -n bookandblockcc -c '{"Args":["insertOffer", "12324354","20", "1","10","LH","hans","PKM","PKV"]}'
#
#


# Todo:
#   persist data on the local storage in order to keep the ledgers
#   some nice-to-know information has been found in the README file of the high-throughput example
#   use https://github.com/hyperledger/fabric-sdk-rest#sample-configuration or https://github.com/Altoros/fabric-rest
#   use a correct configured fabric-ca instance in production environments
#   use a valid policy for further understanding e.g. -P "OR ('Org1MSP.member','Org2MSP.member')"
#   experiment with user enrollment and some ABAC (Attribute-Based Access Control)

# Various notes:
#{
#   "proposal" : {
#        "chaincodeId" : "bookandblockcc",
#        "fcn": "insertOffer",
#        "args": ["12345678910", "0", "20", "2014-11-12T11:45:26.371Z", "2017-11-12T11:45:26.371Z", #"LuxusHotel","Hans","MieterPublicKey","VermieterPublicKey"]
#        }
#}
#{ "status": "SUCCESS",
#  "transactionID": "5d2c1b1b3386ae92757bb9df9e4b52e87c575c43993f55e4802eb770d59939e1"

#docker-compose -f docker-compose-cli.yaml logs --tail=0 --follow

#peer chaincode invoke -o orderer.bookandblock.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem -C bookchannel -n bookandblockcc -c '{"Args":["insertOffer", "12324354","0","20","2014-11-12T11:45:26.371Z","2014-11-12T11:45:26.371Z","LH","hans","PKM","PKV"]}'

#peer chaincode query -C bookchannel -n bookandblockcc -c '{"Args":["getOffer","12324354"]}'

#. scripts/utils.sh might be useful to use

# apt-get update && apt-get upgrade && apt-get install golang-go nodejs (pay attention, latest version) curl docker git docker-compose
# docker installation may differ

# sudo usermod -a -G docker $USER

#attention: delete the chaincode from the docker images to be sure that they will be regenerated. Otherwise the chaincode will not compile with the new version.

#cat ../hfc-key-store/user-key-pub | base64 multiline


#Exit on first error
set -ex #v see https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html

export GOPATH=/usr/share/go #$HOME/go
export PATH=${GOPATH}/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
export CHANNEL_NAME=bookchannel
SDIR=$(dirname "$0")
#source $SDIR/scripts/env.sh might be useful to use with fabric-ca
OS_ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
export COMPOSE_FILE=docker-compose-cli.yaml
export COMPOSE_FILE_COUCH=docker-compose-couch.yaml
LANGUAGE=golang
CC_NAME=bookandblockcc
CORE_PEER_TLS_ENABLED=true
starttime=$(date +%s)

printHelp() {
echo "This script offers various functions in order to start a simple Hyperledger network with 3 peers, one orderer (which is a single point of failure and one organisation)"
echo "For further information see the source"
}

createEnv() {
if [ ! -f .env ]; then
echo "COMPOSE_PROJECT_NAME=net" > .env
echo "IMAGE_TAG=latest" >> .env
fi
}
                        
getAllImagesAndFiles() {
if [ ! -d "fabric-samples" ]; then
curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0
#sudo GOPATH=${GOPATH} go get -u github.com/hyperledger/fabric-ca/cmd/...
#sudo apt install libtool libltdl-dev
#sudo /usr/lib/go-1.10/bin/go get -u github.com/hyperledger/fabric-ca/cmd/...
fi
}

listImages() {
echo "Available Hyperledger and Chaincode docker images:"
docker images | grep -E "hyperledger*|dev-*"
}

generateCerts() {
command -v cryptogen
if [ "$?" -ne 0 ]; then
echo "cryptogen tool not found"
exit 1
fi
if [ -d "crypto-config" ]; then
rm -rf crypto-config
fi
cryptogen generate --config=./crypto-config.yaml
res=$?
if [ ${res} -ne 0 ]; then
echo "Failed to generate certificates"
exit 1
fi
}

generateChannelArtifacts() {
command -v configtxgen
if [ "$?" -ne 0 ]; then
echo "configtxgen tool not found"
exit 1
fi
if [ -d "channel-artifacts" ]; then
rm -rf channel-artifacts
fi
mkdir channel-artifacts
configtxgen -profile OneSingleOrgOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
res=$?
if [ ${res} -ne 0 ]; then
echo "Failed to generate orderer genesis block"
exit 1
fi
configtxgen -profile OneSingleOrgChannel -outputCreateChannelTx ./channel-artifacts/${CHANNEL_NAME}.tx -channelID ${CHANNEL_NAME}
res=$?
if [ ${res} -ne 0 ]; then
echo "Failed to generate channel configuration transaction"
exit 1
fi
configtxgen -profile OneSingleOrgChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgBookAndBlockanchors.tx -channelID ${CHANNEL_NAME} -asOrg OrgBookAndBlock
res=$?
if [ ${res} -ne 0 ]; then
echo "Failed to generate anchor peer update for OrgBookAndBlock"
exit 1
fi
echo
}

setRequiredPeerInfos() {
#$1 = peernumber $2 = bookandblock.com as domain $3 = port
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.$2/users/Admin@org1.$2/msp
echo $CORE_PEER_MSPCONFIGPATH
export CORE_PEER_ADDRESS=peer$1.org1.$2:$3
echo $CORE_PEER_ADDRESS
export CORE_PEER_LOCALMSPID="OrgBookAndBlock"
echo $CORE_PEER_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.$2/peers/peer$1.org1.$2/tls/ca.crt
echo $CORE_PEER_TLS_ROOTCERT_FILE
}

createAndJoinTheChannelForAllPeers() {
#remember to set these values upon execution
setRequiredPeerInfos 0 bookandblock.com 7051
docker exec cli peer channel create -o orderer.bookandblock.com:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/${CHANNEL_NAME}.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem
#going to sleep in order to finish channel setup
sleep 30s
i=0
portNumber=7051
for VARIABLE in peer0.org1.bookandblock.com peer1.org1.bookandblock.com peer2.org1.bookandblock.com
do
setRequiredPeerInfos ${i} bookandblock.com ${portNumber}

echo "CORE_PEER_ADDRESS=peer${i}.org1.bookandblock.com:7051 peer channel join -b bookchannel.block" | docker exec --interactive cli /bin/bash -

docker exec cli peer channel list
i=$((i+1))
done

docker exec cli peer channel update -o orderer.bookandblock.com:7050 -c ${CHANNEL_NAME} -f ./channel-artifacts/OrgBookAndBlockanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem
}

askProceed() {
read -p "Continue? [Y/n] " answer
case "$answer" in
y|Y|"" )
echo "proceeding ..."
;;
n|N )
echo "No selected"
exit 1
;;
* )
echo "Invalid response"
askProceed
;;
esac
}

recreateHfcKeyStore() {
echo "Recreate key-store?"
askProceed
rm -rf .hfc-key-store
mkdir .hfc-key-store
}

replaceKeyFile() {
sed -i "s/$1/$2/g" base/docker-compose-base.yaml
}

startNetwork() {
string=$(ls ./crypto-config/peerOrganizations/org1.bookandblock.com/ca/ | grep _sk)
replaceKeyFile "STRING_TO_REPLACE" $string
docker-compose -f ${COMPOSE_FILE} -f ${COMPOSE_FILE_COUCH} up -d
}

stopNetwork() {
docker-compose -f ${COMPOSE_FILE} -f ${COMPOSE_FILE_COUCH} down
#docker kill might be empty if all targeted images have been shutdown properly
#docker kill $(docker ps -q)
}

stopNetworkAndDestroyAllData() {
stopNetwork
string=$(ls ./crypto-config/peerOrganizations/org1.bookandblock.com/ca/ | grep _sk)
replaceKeyFile $string "STRING_TO_REPLACE"
#docker rm -f $(docker ps -aq)
docker network prune
askProceed
#docker rmi $(docker images -q)
docker container prune
docker volume prune
}

createAndJoinTheChannelForVariousPeers() {
#TODO if necessary
exit 7090
}

processChaincode() {
#docker exec cli peer chaincode $1 -l ${LANGUAGE} -o orderer.bookandblock.com:7050 --tls ${CORE_PEER_TLS_ENABLED} --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem -C ${CHANNEL_NAME} -n ${CC_NAME} -v 0.1 -c '{"Args": []}'

#peer chaincode instantiate -n mycc -v 0 -c '{"Args":[""]}' -C myc
#${CORE_PEER_TLS_ENABLED} 
echo "peer chaincode $1 -o orderer.bookandblock.com:7050 --tls ${CORE_PEER_TLS_ENABLED} --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookandblock.com/orderers/orderer.bookandblock.com/msp/tlscacerts/tlsca.bookandblock.com-cert.pem -C ${CHANNEL_NAME} -n ${CC_NAME} -v 0.1 -l ${LANGUAGE} -c '{\"Args\": [\"\"]}'" | docker exec --interactive cli /bin/bash -
}

installAndInstantiateChaincode() {
docker exec cli mkdir -p /opt/gopath/src/github.com/hyperledger/fabric/bookandblock/chaincode/go
docker cp bookandblockcc.go cli:/opt/gopath/src/github.com/hyperledger/fabric/bookandblock/chaincode/go
#install the chaincode on every peer
i=0
portNumber=7051
for VARIABLE in peer0.org1.bookandblock.com peer1.org1.bookandblock.com peer2.org1.bookandblock.com
do
setRequiredPeerInfos ${i} bookandblock.com ${portNumber}
echo "CORE_PEER_ADDRESS=peer${i}.org1.bookandblock.com:7051 peer chaincode install -l ${LANGUAGE} -n ${CC_NAME} -v 0.1 -p github.com/hyperledger/fabric/bookandblock/chaincode/go/" | docker exec --interactive cli /bin/bash -
i=$((i+1))
done
#pay attention for the correct instantiation arguments
sleep 10s
processChaincode instantiate
}

upgradeChaincode() {
processChaincode upgrade
}

redeployChaincode()  {
processChaincode instantiate
}

hyperledgerFabricCaHint() {
#TODO start Hyperledger Fabric from an external source if it DOESN'T START
exit 7090
}

getRestService() {
# clone dev branch for testing purposes
git clone -b dev https://github.com/hyperledger/fabric-sdk-rest.git

cd fabric-sdk-rest/
#--unsafe-perm=true --allow-root
#sudo apt-get install libcairo2-dev libjpeg-dev libgif-dev
#maybe put the ReST service into a docker container
cd packages/loopback-connector-fabric && sudo npm link && cd ..
cd packages/fabric-rest && npm link loopback-connector-fabric && cd ..
cd packages/fabric-rest && npm install && cd ..
cd ..

# start the rest service here and don't forget to modify datasource.json
# ./fabric-rest-server -l '{"debug":"console"}' -d

#https://hub.docker.com/r/maxxx1313/fabric-rest/
}

main() {
if [ "$#" -lt 1 ]; then
printHelp
exit 1
fi
if [ "$#" -eq 3 ]; then
setRequiredPeerInfos $1 $2 $3
exit 1
fi
if [ "$1" = "rhks" ]; then
recreateHfcKeyStore
exit 1
fi
if ([ "$1" = "stop" ] || [ "$1" = "down" ]); then
stopNetwork
exit 1
fi
if ([ "$1" = "start" ] || [ "$1" = "up" ]); then
startNetwork
exit 1
fi
if [ "$1" = "destroy" ]; then
stopNetworkAndDestroyAllData
exit 1
fi
if [ "$1" = "build" ]; then
createEnv
getAllImagesAndFiles
listImages
generateCerts
generateChannelArtifacts
startNetwork
createAndJoinTheChannelForAllPeers
installAndInstantiateChaincode
printf "\nTotal setup execution time : $(($(date +%s) - starttime)) secs\n\n\n"
exit 1
fi
}

main "$@"


