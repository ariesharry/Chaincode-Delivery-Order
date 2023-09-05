package chaincode

import (
	"encoding/json"
    "fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (doc *DeliveryOrderContract) ReleaseDO(ctx contractapi.TransactionContextInterface, id string) error {
	requestDetailsAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return err
	}

	if requestDetailsAsBytes == nil {
        return fmt.Errorf("No request with ID %s exists", id)
    }

	var requestDetails RequestDetails
	json.Unmarshal(requestDetailsAsBytes, &requestDetails)
	
	requestDetails.Status = "release"
	requestDetailsAsBytes, _ = json.Marshal(requestDetails)
	return ctx.GetStub().PutState(id, requestDetailsAsBytes)
}

func (doc *DeliveryOrderContract) RejectDO(ctx contractapi.TransactionContextInterface, id string) error {
    requestDetailsAsBytes, err := ctx.GetStub().GetState(id)
    if err != nil {
        return err
    }
    if requestDetailsAsBytes == nil {
        return fmt.Errorf("No request with ID %s exists", id)
    }

    var requestDetails RequestDetails
    json.Unmarshal(requestDetailsAsBytes, &requestDetails)

    requestDetails.Status = "rejected"
    requestDetailsAsBytes, _ = json.Marshal(requestDetails)

    return ctx.GetStub().PutState(id, requestDetailsAsBytes)
}

