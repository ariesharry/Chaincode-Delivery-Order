package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RequestDetails represents the structure of the delivery order request details
type RequestDetails struct {
	id            string        `json:"id"` // nomor request do
	requestor     Requestor     `json:"requestor"`
	shippingLine  ShippingLine  `json:"shippingLine"`
	payment       Payment       `json:"payment"`
	document      Document      `json:"document"`
	parties       Parties       `json:"parties"`
	cargoDetails  CargoDetails  `json:"cargoDetails"`
	location      Location      `json:"location"`
	paymentDetail PaymentDetail `json:"paymentDetail"`
	supportingDoc SupportingDoc `json:"supportingDocument"`
	status        string        `json:"status"`
}

type Requestor struct {
	requestorType    string `json:"requestorType"`
	urlFile          string `json:"urlFile"`
	npwp             string `json:"npwp"`
	nib              string `json:"nib"`
	requestorName    string `json:"requestorName"`
	requestorAddress string `json:"requestorAddress"`
}

type ShippingLine struct {
	shippingType string `json:"shippingType"`
	doExpired    string `json:"doExpired"`
	vesselName   string `json:"vesselName"`
	voyageNumber string `json:"voyageNumber"`
}

type Payment struct {
	termOfPayment string `json:"termOfPayment"`
}

type Document struct {
	ladingBillNumber string `json:"ladingBillNumber"`
	ladingBillDate   string `json:"ladingBillDate"`
	ladingBillType   string `json:"ladingBillType"`
	urlFile          string `json:"urlFile"`
}

type Parties struct {
	shipper     Party `json:"shipper"`
	consignee   Party `json:"consignee"`
	notifyParty Party `json:"notifyParty"`
}

type Party struct {
	name string `json:"name"`
	npwp string `json:"npwp"`
}

type CargoDetails struct {
	isContainer  bool         `json:"isContainer"`
	container    Container    `json:"container"`
	nonContainer NonContainer `json:"nonContainer"`
}

type Container struct {
	containerSeq int64       `json:"containerSeq"`
	containerNo  string      `json:"containerNo"`
	sealNo       string      `json:"sealNo"`
	sizeType     SizeType    `json:"sizeType"`
	grossWeight  GrossWeight `json:"grossWeight"`
	ownership    string      `json:"ownership"`
}

type NonContainer struct {
	nonContainerSeq   int64             `json:"nonContainerSeq"`
	goodsDescription  string            `json:"goodsDescription"`
	packageQuantity   PackageQuantity   `json:"packageQuantity"`
	grossWeight       GrossWeight       `json:"grossWeight"`
	measurementVolume MeasurementVolume `json:"measurementVolume"`
}

type PackageQuantity struct {
	amount float64 `json:"amount"`
	unit   string  `json:"unit"`
}

type MeasurementVolume struct {
	amount float64 `json:"amount"`
	unit   string  `json:"unit"`
}

type SizeType struct {
	size int64  `json:"size"`
	tipe string `json:"type"`
}

type GrossWeight struct {
	amount float64 `json:"amount"`
	unit   string  `json:"unit"`
}

type Location struct {
	locationType LocationType `json:"locationType"`
}

type LocationType struct {
	location    string `json:"location"`
	countryCode string `json:"countryCode"`
	portCode    string `json:"portCode"`
}

type PaymentDetail struct {
	invoice Invoice `json:"invoice"`
}

type Invoice struct {
	invoiceNo   string  `json:"invoiceNo"`
	invoiceDate string  `json:"invoiceDate"`
	totalAmount float64 `json:"totalAmount"`
	bankID      string  `json:"bankId"`
	accountNo   string  `json:"accountNo"`
	urlFile     string  `json:"urlFile"`
}

type SupportingDoc struct {
	documentType DocumentType `json:"documentType"`
}

type DocumentType struct {
	typeDocument string `json:"document"`
	documentNo   string `json:"documentNo"`
	documentDate string `json:"documentDate"`
	urlFile      string `json:"urlFile"`
}

type DeliveryOrderContract struct {
	contractapi.Contract
}

func NewDeliveryOrderContract() contractapi.ContractInterface {
	return &DeliveryOrderContract{
		Contract: contractapi.Contract{},
	}
}
