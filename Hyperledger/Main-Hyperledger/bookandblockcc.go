/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

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

type offer struct {
	OfferId    int64     `json:"offerId"`
	Free       bool      `json:"free"`
	Price      float64   `json:"price"`
	CheckIn    time.Time `json:"checkIn"`
	CheckOut   time.Time `json:"checkOut"`
	ObjectName string    `json:"objectName"`
	OwnerName  string    `json:"ownerName"`
	TenantPk   string    `json:"tenantPk"`
	LandlordPk string    `json:"landlordPk"`
}

// Maybe change this to unixtime , see time.Unix
const datelayout = "2006-01-02T15:04:05.000Z"

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

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

func (t *SimpleChaincode) rentAnOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	offerId := args[0]
	tenantPk := args[1]
	checkIn := time.Now()
	free := false

	fmt.Println("- start rentOffer ", offerId, checkIn, free, tenantPk)

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

	offerToRent.Free = free
	offerToRent.CheckIn = checkIn
	offerToRent.TenantPk = tenantPk

	offerJSONasBytes, _ := json.Marshal(offerToRent)
	err = stub.PutState(offerId, offerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferOffer (success)")
	return shim.Success(nil)
}

func (t *SimpleChaincode) insertOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
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
	if len(args[8]) <= 0 {
		return shim.Error("9th argument must be a non-empty string")
	}

	offerId, err := strconv.ParseInt(args[0], 10, 64)
	if (err != nil) {
		return shim.Error(fmt.Sprintf("failed by parsing offerId: %s", err))
	}
	free, err := strconv.ParseBool(args[1])
	if (err != nil) {
		return shim.Error(fmt.Sprintf("failed by parsing free: %s", err))
	}
	price, err := strconv.ParseFloat(args[2], 64)
	if (err != nil) {
		return shim.Error(fmt.Sprintf("failed by parsing price: %s", err))
	}
	checkIn, err := time.Parse(datelayout, args[3])
	if (err != nil) {
		return shim.Error(fmt.Sprintf("failed by parsing checkIn: %s", err))
	}
	checkOut, err := time.Parse(datelayout, args[4])
	if (err != nil) {
		return shim.Error(fmt.Sprintf("failed by parsing checkOut: %s", err))
	}
	objectName := args[5]
	ownerName := args[6]
	tenantPk := args[7]
	landlordPk := args[8]

	offerAsBytes, err := stub.GetState(strconv.FormatInt(offerId, 10))
	if err != nil {
		return shim.Error("Failed to get offer: " + err.Error())
	} else if offerAsBytes != nil {
		fmt.Println("This offer already exists: " + strconv.FormatInt(offerId, 10))
		return shim.Error("This offer already exists: " + strconv.FormatInt(offerId, 10))
	}

	offer := &offer{offerId, free, price, checkIn, checkOut, objectName, ownerName, tenantPk, landlordPk}
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

