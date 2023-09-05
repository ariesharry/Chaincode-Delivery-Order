package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RequestDetails represents the structure of the delivery order request details
type RequestDetails struct {
	ID				string 		`json:"id"`
	Requestor     Requestor     `json:"requestor"`
	ShippingLine  ShippingLine  `json:"shippingLine"`
	Payment       Payment       `json:"payment"`
	Document      Document      `json:"document"`
	Parties       Parties       `json:"parties"`
	CargoDetails  CargoDetails  `json:"cargoDetails"`
	Location      Location      `json:"location"`
	PaymentDetail PaymentDetail `json:"paymentDetail"`
	SupportingDoc SupportingDoc `json:"supportingDocument"`
	Status        string 		`json:"status"`
}

type Requestor struct {
	RequestorType string `json:"requestorType"`
	URLFile       string `json:"urlFile"`
	NPWP          string `json:"npwp"`
	NIB           string `json:"nib"`
	RequestorName string `json:"requestorName"`
	RequestorAddr string `json:"requestorAddress"`
}

type ShippingLine struct {
	ShippingType string `json:"shippingType"`
	DOExpired    string `json:"doExpired"`
	VesselName   string `json:"vesselName"`
	VoyageNumber string `json:"voyageNumber"`
	Payment      string `json:"payment"`
}

type Payment struct {
	TermOfPayment string `json:"termOfPayment"`
}

type Document struct {
	LadingBillNumber string `json:"ladingBillNumber"`
	LadingBillDate   string `json:"ladingBillDate"`
	LadingBillType   string `json:"ladingBillType"`
	URLFile          string `json:"urlFile"`
}

type Parties struct {
	Shipper     Party `json:"shipper"`
	Consignee   Party `json:"consignee"`
	NotifyParty Party `json:"notifyParty"`
}

type Party struct {
	Name string `json:"name"`
	NPWP string `json:"npwp"`
}

type CargoDetails struct {
	Container Container `json:"container"`
}

type Container struct {
	ContainerSeq   string `json:"containerSeq"`
	ContainerNo    string `json:"containerNo"`
	SealNo         string `json:"sealNo"`
	SizeType       string `json:"sizeType"`
	GrossWeight    string `json:"grossWeight"`
	Ownership      string `json:"ownership"`
}

type Location struct {
	LocationType string `json:"locationType"`
	Location     string `json:"location"`
	CountryCode  string `json:"countryCode"`
	PortCode     string `json:"portCode"`
}

type PaymentDetail struct {
	Invoice Invoice `json:"invoice"`
}

type Invoice struct {
	InvoiceNo    string `json:"invoiceNo"`
	InvoiceDate  string `json:"invoiceDate"`
	TotalAmount  string `json:"totalAmount"`
	BankID       string `json:"bankId"`
	AccountNo    string `json:"accountNo"`
	URLFile      string `json:"urlFile"`
}

type SupportingDoc struct {
	DocumentType DocumentType `json:"documentType"`
}

type DocumentType struct {
	DocumentNo   string `json:"documentNo"`
	DocumentDate string `json:"documentDate"`
	URLFile      string `json:"urlFile"`
}

type DeliveryOrderContract struct {
	contractapi.Contract
}

func NewDeliveryOrderContract() contractapi.ContractInterface {
	return &DeliveryOrderContract{
		Contract: contractapi.Contract{},
	}
}
