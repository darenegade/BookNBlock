package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// This is the structure of an offer from BookAndBlock.
// It is based on the official documentation, see: https://github.com/darenegade/BookNBlock/blob/master/Aufgabenbeschreibung.pdf
type offer struct {
	OfferId    int64   `json:"offerId"`
	Price      float64 `json:"price"`
	CheckIn    string  `json:"checkIn"`
	CheckOut   string  `json:"checkOut"`
	ObjectName string  `json:"objectName"`
	OwnerName  string  `json:"ownerName"`
	TenantPk   string  `json:"tenantPk"`
	LandlordPk string  `json:"landlordPk"`
	CheckInM   string  `json:"checkInM"`
	CheckOutM  string  `json:"checkOutM"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// The Invoke function is always the first call of an invocation of the Chaincode, it redirects the request to the correct
// function.
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// All available functions of the Chaincode are called from this Invoke function.
	if function == "insertOffer" {
		return t.insertOffer(stub, args)
	} else if function == "transferOffer" {
		return t.transferOffer(stub, args)
	} else if function == "delete" {
		return t.delete(stub, args)
	} else if function == "getOffer" {
		return t.getOffer(stub, args)
	} else if function == "queryOffersByPk" {
		return t.queryOffersByPk(stub, args)
	} else if function == "queryOffers" {
		return t.queryOffers(stub, args)
	} else if function == "getHistoryForOffer" {
		return t.getHistoryForOffer(stub, args)
	} else if function == "rentAnOffer" {
		return t.rentAnOffer(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)
	return shim.Error("Received unknown function invocation")
}

// This function rents an available offer. Handling of availability should be done by the frontend.
// Parameters in the correct order are:
// "12324354" <- the id of the offer (should be unique)
// "PKM" <- public Key of the tenant
// "CheckInM" <- checkIn of the tenant
// "CheckOutM" checkOut of the tenant
func (t *SimpleChaincode) rentAnOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	offerId := args[0]
	tenantPk := args[1]
	checkIn := args[2]
	checkOut := args[3]

	fmt.Println("- start rentOffer ", offerId, checkIn, tenantPk)

	offerAsBytes, err := stub.GetState(offerId)
	if err != nil {
		return shim.Error("Failed to get offer:" + err.Error())
	} else if offerAsBytes == nil {
		return shim.Error("Offer does not exist")
	}

	offerToRent := offer{}
	err = json.Unmarshal(offerAsBytes, &offerToRent)
	if err != nil {
		return shim.Error(err.Error())
	}

	if !isNotOverlapping(stub, args[0], checkIn, checkOut) {
		return shim.Error("Overlapping got triggered")
	}

	offerToRent.CheckInM = checkIn
	offerToRent.CheckOutM = checkOut
	offerToRent.TenantPk = tenantPk

	offerJSONasBytes, _ := json.Marshal(offerToRent)
	err = stub.PutState(offerId, offerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferOffer (success)")
	return shim.Success(nil)
}

// This function inserts an offer with various parameters as the code below shows.
// Parameters in the correct order are:
// "12324354" <- the id of the offer (should be unique)
// "20" <- the price of the offer
// "1528053733" <- the start date of the booking (as unixtime format)
// "1528053734" <- the end date of the booking (as unixtime format)
// "LuxusHotel" <- the name of the offer
// "hans" <- the name of the owner
// "PKM" <- the public key of the tenant (pay attention for the correct format, for example base64)
// "PKV" <- the public key of the landlord (pay attention for the correct format, for example base64)
func (t *SimpleChaincode) insertOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	fmt.Println("- start init offer")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th argument must be a non-empty string")
	}

	offerId, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed by parsing offerId: %s", err))
	}

	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed by parsing price: %s", err))
	}
	checkIn := args[2]
	checkOut := args[3]
	objectName := args[4]
	ownerName := args[5]
	tenantPk := args[6]
	landlordPk := args[7]

	offerAsBytes, err := stub.GetState(strconv.FormatInt(offerId, 10))
	if err != nil {
		return shim.Error("Failed to get offer: " + err.Error())
	} else if offerAsBytes != nil {
		fmt.Println("This offer already exists: " + strconv.FormatInt(offerId, 10))
		return shim.Error("This offer already exists: " + strconv.FormatInt(offerId, 10))
	}

	offer := &offer{offerId, price, checkIn, checkOut, objectName, ownerName, tenantPk, landlordPk, "", ""}
	offerJSONasBytes, err := json.Marshal(offer)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.FormatInt(offerId, 10), offerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	indexName := "ownerName~objectName"
	ownerNameObjectNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{offer.OwnerName, offer.ObjectName})
	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	stub.PutState(ownerNameObjectNameIndexKey, value)

	fmt.Println("- end insert offer")
	return shim.Success(nil)
}

// This function may fail due to the character of a Blockchain. Especially something like a race condition may be seen when using
// this function. In order to minimize issues, adjust the transaction batch timeout and the number of transactions collected in the configuration file configtx.yaml
// Check the parameters MaxMessageCount and BatchTimeout especially.
func isNotOverlapping(stub shim.ChaincodeStubInterface, offerId string, checkIn string, checkOut string) bool {
	resultsIterator, err := stub.GetHistoryForKey(offerId)
	if err != nil {
		fmt.Println("- Fail ")
		return false
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()

		offerInBlockchain := offer{}
		err = json.Unmarshal(response.Value, &offerInBlockchain)
		if err != nil {
			fmt.Println("- Fail ")
			return false
		}

		in, err := strconv.ParseInt(checkIn, 10, 64)
		out, err := strconv.ParseInt(checkOut, 10, 64)
		inBlockchain, err := strconv.ParseInt(offerInBlockchain.CheckIn, 10, 64)
		outBlockchain, err := strconv.ParseInt(offerInBlockchain.CheckOut, 10, 64)
		if in > out || in < inBlockchain || out > outBlockchain {
			return false
		}

		inBlockchainM, err := strconv.ParseInt(offerInBlockchain.CheckInM, 10, 64)
		outBlockchainM, err := strconv.ParseInt(offerInBlockchain.CheckOutM, 10, 64)
		if (inBlockchainM != 0 && outBlockchainM != 0) && !(out < inBlockchainM || outBlockchainM < in) {
			return false
		}
	}
	return true

}

// This function returns an offer for a given offer id. An example parameter could be "12324354" as the function above this one describes.
func (t *SimpleChaincode) getOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp, offerIdAsString string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the offer to query")
	}

	offerIdAsString = args[0]

	valAsbytes, err := stub.GetState(offerIdAsString)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + offerIdAsString + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Offer does not exist: " + offerIdAsString + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// This function "deletes" an offer for a given id, which could be "12324354".
// The deletion call in the Blockchain provided by Hyperledger is an internal mechanism.
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var offerJSON offer
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	offerIdAsString := args[0]
	valAsbytes, err := stub.GetState(offerIdAsString)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + offerIdAsString + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Offer does not exist: " + offerIdAsString + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &offerJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + offerIdAsString + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(offerIdAsString)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	indexKey := "ownerName~objectName"
	ownerNameObjectNameIndexKey, err := stub.CreateCompositeKey(indexKey, []string{offerJSON.OwnerName, offerJSON.ObjectName})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(ownerNameObjectNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// This function transfers an offer.
// Parameters in the correct order are:
// "12324354" <- the id of the offer (should be unique)
// "martin" <- the new name of the owner
// "PKV" <- the public key of the new landlord (pay attention for the correct format, for example base64)
func (t *SimpleChaincode) transferOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	offerId := args[0]
	newOwnerName := args[1]
	newLandlordPK := args[2]

	fmt.Println("- start transferOffer ", offerId, newOwnerName, newLandlordPK)

	offerAsBytes, err := stub.GetState(offerId)
	if err != nil {
		return shim.Error("Failed to get offer:" + err.Error())
	} else if offerAsBytes == nil {
		return shim.Error("Offer does not exist")
	}

	offerToTransfer := offer{}
	err = json.Unmarshal(offerAsBytes, &offerToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}

	offerToTransfer.OwnerName = newOwnerName
	offerToTransfer.LandlordPk = newLandlordPK

	offerJSONasBytes, _ := json.Marshal(offerToTransfer)
	err = stub.PutState(offerId, offerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferOffer (success)")
	return shim.Success(nil)
}

// This function gets all offers specified by the public key of a landlord. (pay attention for the correct format, for example base64)
// "PKV" <- the public key of the new landlord (pay attention for the correct format, for example base64)
func (t *SimpleChaincode) queryOffersByPk(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	publicKey := args[0]

	queryString := fmt.Sprintf(
		"{\"selector\":{"+
			"\"landlordPk\":\"%s\"}}", publicKey)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// This function gives the possibility to create a rich query for filtering values in the Blockchain.
// Parameters in the correct order are:
// "queryString" <- the query string, see for example: https://hyperledger-fabric.readthedocs.io/en/release-1.1/couchdb_as_state_database.html
// This allows very extensive queries against the state db (currently CouchDB) of the Blockchain and shouldn't be underestimated for complex scenarios.
func (t *SimpleChaincode) queryOffers(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// This function handles the logic of the rich query. See the function queryOffers above.
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// This function gets the history for an offer of a given id, which could be "12324354".
// The values should be in a chronological order and illustrate the functionality of a Blockchain.
func (t *SimpleChaincode) getHistoryForOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	offerId := args[0]

	fmt.Printf("- start getHistoryForOffer: %s\n", offerId)

	resultsIterator, err := stub.GetHistoryForKey(offerId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()

		offerToRent := offer{}
		err = json.Unmarshal(response.Value, &offerToRent)
		if err != nil {
			return shim.Error(err.Error())
		}

		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForOffer returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
