package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RequestDetails represents the structure of the delivery order request details
type RequestDetails struct {
	Id            string        `json:"id"` // nomor request do
	Requestor     Requestor     `json:"requestor"`
	ShippingLine  ShippingLine  `json:"shippingLine"`
	Payment       Payment       `json:"payment"`
	Document      Document      `json:"document"`
	Parties       Parties       `json:"parties"`
	CargoDetails  CargoDetails  `json:"cargoDetails"`
	Location      Location      `json:"location"`
	PaymentDetail PaymentDetail `json:"paymentDetail"`
	SupportingDoc SupportingDoc `json:"supportingDocument"`
	Status        string        `json:"status"`
}

type Requestor struct {
	RequestorType    string `json:"requestorType"`
	UrlFile          string `json:"urlFile"`
	NPWP             string `json:"npwp"`
	NIB              string `json:"nib"`
	RequestorName    string `json:"requestorName"`
	RequestorAddress string `json:"requestorAddress"`
}

type ShippingLine struct {
	ShippingType string `json:"shippingType"`
	DoExpired    string `json:"doExpired"`
	VesselName   string `json:"vesselName"`
	VoyageNumber string `json:"voyageNumber"`
}

type Payment struct {
	TermOfPayment string `json:"termOfPayment"`
}

type Document struct {
	LadingBillNumber string `json:"ladingBillNumber"`
	LadingBillDate   string `json:"ladingBillDate"`
	LadingBillType   string `json:"ladingBillType"`
	UrlFile          string `json:"urlFile"`
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
	IsContainer  bool         `json:"isContainer"`
	Container    Container    `json:"container"`
	NonContainer NonContainer `json:"nonContainer"`
}

type Container struct {
	ContainerSeq int64       `json:"containerSeq"`
	ContainerNo  string      `json:"containerNo"`
	SealNo       string      `json:"sealNo"`
	SizeType     SizeType    `json:"sizeType"`
	GrossWeight  GrossWeight `json:"grossWeight"`
	Ownership    string      `json:"ownership"`
}

type NonContainer struct {
	NonContainerSeq   int64             `json:"nonContainerSeq"`
	GoodsDescription  string            `json:"goodsDescription"`
	PackageQuantity   PackageQuantity   `json:"packageQuantity"`
	GrossWeight       GrossWeight       `json:"grossWeight"`
	MeasurementVolume MeasurementVolume `json:"measurementVolume"`
}

type PackageQuantity struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type MeasurementVolume struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type SizeType struct {
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type GrossWeight struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type Location struct {
	LocationType LocationType `json:"locationType"`
}

type LocationType struct {
	Location    string `json:"location"`
	CountryCode string `json:"countryCode"`
	PortCode    string `json:"portCode"`
}

type PaymentDetail struct {
	Invoice Invoice `json:"invoice"`
}

type Invoice struct {
	InvoiceNo   string  `json:"invoiceNo"`
	InvoiceDate string  `json:"invoiceDate"`
	TotalAmount float64 `json:"totalAmount"`
	BankID      string  `json:"bankId"`
	AccountNo   string  `json:"accountNo"`
	UrlFile     string  `json:"urlFile"`
}

type SupportingDoc struct {
	DocumentType DocumentType `json:"documentType"`
}

type DocumentType struct {
	TypeDocument string `json:"document"`
	DocumentNo   string `json:"documentNo"`
	DocumentDate string `json:"documentDate"`
	UrlFile      string `json:"urlFile"`
}

type DeliveryOrderContract struct {
	contractapi.Contract
}

func NewDeliveryOrderContract() contractapi.ContractInterface {
	return &DeliveryOrderContract{
		Contract: contractapi.Contract{},
	}
}
