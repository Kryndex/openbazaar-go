// Code generated by protoc-gen-go.
// source: contracts.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	contracts.proto

It has these top-level messages:
	RicardianContract
	Listing
	Order
	OrderConfirmation
	Rating
	Dispute
	DisputeResolution
	Refund
	Timestamp
	ID
	Signatures
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Listing_Metadata_CategorySub int32

const (
	Listing_Metadata_NA          Listing_Metadata_CategorySub = 0
	Listing_Metadata_FIXED_PRICE Listing_Metadata_CategorySub = 1
	Listing_Metadata_AUCTION     Listing_Metadata_CategorySub = 2
)

var Listing_Metadata_CategorySub_name = map[int32]string{
	0: "NA",
	1: "FIXED_PRICE",
	2: "AUCTION",
}
var Listing_Metadata_CategorySub_value = map[string]int32{
	"NA":          0,
	"FIXED_PRICE": 1,
	"AUCTION":     2,
}

func (x Listing_Metadata_CategorySub) String() string {
	return proto.EnumName(Listing_Metadata_CategorySub_name, int32(x))
}
func (Listing_Metadata_CategorySub) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0, 0}
}

type Listing_Metadata_Category int32

const (
	Listing_Metadata_UNKNOWN       Listing_Metadata_Category = 0
	Listing_Metadata_PHYSICAL_GOOD Listing_Metadata_Category = 1
	Listing_Metadata_DIGITAL_GOOD  Listing_Metadata_Category = 2
	Listing_Metadata_SERVICE       Listing_Metadata_Category = 3
)

var Listing_Metadata_Category_name = map[int32]string{
	0: "UNKNOWN",
	1: "PHYSICAL_GOOD",
	2: "DIGITAL_GOOD",
	3: "SERVICE",
}
var Listing_Metadata_Category_value = map[string]int32{
	"UNKNOWN":       0,
	"PHYSICAL_GOOD": 1,
	"DIGITAL_GOOD":  2,
	"SERVICE":       3,
}

func (x Listing_Metadata_Category) String() string {
	return proto.EnumName(Listing_Metadata_Category_name, int32(x))
}
func (Listing_Metadata_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0, 1}
}

type Order_Payment_Method int32

const (
	Order_Payment_DIRECT    Order_Payment_Method = 0
	Order_Payment_MODERATED Order_Payment_Method = 1
)

var Order_Payment_Method_name = map[int32]string{
	0: "DIRECT",
	1: "MODERATED",
}
var Order_Payment_Method_value = map[string]int32{
	"DIRECT":    0,
	"MODERATED": 1,
}

func (x Order_Payment_Method) String() string {
	return proto.EnumName(Order_Payment_Method_name, int32(x))
}
func (Order_Payment_Method) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 2, 0} }

type Signatures_Section int32

const (
	Signatures_LISTING            Signatures_Section = 0
	Signatures_ORDER              Signatures_Section = 1
	Signatures_ORDER_CONFIRMATION Signatures_Section = 2
	Signatures_RATING             Signatures_Section = 3
	Signatures_DISPUTE            Signatures_Section = 4
	Signatures_DISPUTE_RESOLUTION Signatures_Section = 5
	Signatures_REFUND             Signatures_Section = 6
)

var Signatures_Section_name = map[int32]string{
	0: "LISTING",
	1: "ORDER",
	2: "ORDER_CONFIRMATION",
	3: "RATING",
	4: "DISPUTE",
	5: "DISPUTE_RESOLUTION",
	6: "REFUND",
}
var Signatures_Section_value = map[string]int32{
	"LISTING":            0,
	"ORDER":              1,
	"ORDER_CONFIRMATION": 2,
	"RATING":             3,
	"DISPUTE":            4,
	"DISPUTE_RESOLUTION": 5,
	"REFUND":             6,
}

func (x Signatures_Section) String() string {
	return proto.EnumName(Signatures_Section_name, int32(x))
}
func (Signatures_Section) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{10, 0} }

type RicardianContract struct {
	VendorListing           []*Listing         `protobuf:"bytes,1,rep,name=vendorListing" json:"vendorListing,omitempty"`
	BuyerOrder              *Order             `protobuf:"bytes,2,opt,name=buyerOrder" json:"buyerOrder,omitempty"`
	VendorOrderConfirmation *OrderConfirmation `protobuf:"bytes,3,opt,name=vendorOrderConfirmation" json:"vendorOrderConfirmation,omitempty"`
	BuyerRating             *Rating            `protobuf:"bytes,4,opt,name=buyerRating" json:"buyerRating,omitempty"`
	Dispute                 *Dispute           `protobuf:"bytes,5,opt,name=dispute" json:"dispute,omitempty"`
	DisputeResolution       *DisputeResolution `protobuf:"bytes,6,opt,name=disputeResolution" json:"disputeResolution,omitempty"`
	Refund                  *Refund            `protobuf:"bytes,7,opt,name=refund" json:"refund,omitempty"`
	Signatures              []*Signatures      `protobuf:"bytes,8,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *RicardianContract) Reset()                    { *m = RicardianContract{} }
func (m *RicardianContract) String() string            { return proto.CompactTextString(m) }
func (*RicardianContract) ProtoMessage()               {}
func (*RicardianContract) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RicardianContract) GetVendorListing() []*Listing {
	if m != nil {
		return m.VendorListing
	}
	return nil
}

func (m *RicardianContract) GetBuyerOrder() *Order {
	if m != nil {
		return m.BuyerOrder
	}
	return nil
}

func (m *RicardianContract) GetVendorOrderConfirmation() *OrderConfirmation {
	if m != nil {
		return m.VendorOrderConfirmation
	}
	return nil
}

func (m *RicardianContract) GetBuyerRating() *Rating {
	if m != nil {
		return m.BuyerRating
	}
	return nil
}

func (m *RicardianContract) GetDispute() *Dispute {
	if m != nil {
		return m.Dispute
	}
	return nil
}

func (m *RicardianContract) GetDisputeResolution() *DisputeResolution {
	if m != nil {
		return m.DisputeResolution
	}
	return nil
}

func (m *RicardianContract) GetRefund() *Refund {
	if m != nil {
		return m.Refund
	}
	return nil
}

func (m *RicardianContract) GetSignatures() []*Signatures {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Listing struct {
	ListingName        string            `protobuf:"bytes,1,opt,name=listingName" json:"listingName,omitempty"`
	VendorID           *ID               `protobuf:"bytes,2,opt,name=vendorID" json:"vendorID,omitempty"`
	Metadata           *Listing_Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	Item               *Listing_Item     `protobuf:"bytes,4,opt,name=item" json:"item,omitempty"`
	Shipping           *Listing_Shipping `protobuf:"bytes,5,opt,name=shipping" json:"shipping,omitempty"`
	Moderators         []string          `protobuf:"bytes,6,rep,name=moderators" json:"moderators,omitempty"`
	TermsAndConditions string            `protobuf:"bytes,7,opt,name=termsAndConditions" json:"termsAndConditions,omitempty"`
	RefundPolicy       string            `protobuf:"bytes,8,opt,name=refundPolicy" json:"refundPolicy,omitempty"`
}

func (m *Listing) Reset()                    { *m = Listing{} }
func (m *Listing) String() string            { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()               {}
func (*Listing) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Listing) GetVendorID() *ID {
	if m != nil {
		return m.VendorID
	}
	return nil
}

func (m *Listing) GetMetadata() *Listing_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listing) GetItem() *Listing_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Listing) GetShipping() *Listing_Shipping {
	if m != nil {
		return m.Shipping
	}
	return nil
}

type Listing_Metadata struct {
	Version     uint32                       `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Category    Listing_Metadata_Category    `protobuf:"varint,2,opt,name=category,enum=pb.Listing_Metadata_Category" json:"category,omitempty"`
	CategorySub Listing_Metadata_CategorySub `protobuf:"varint,3,opt,name=categorySub,enum=pb.Listing_Metadata_CategorySub" json:"categorySub,omitempty"`
	Expiry      *Timestamp                   `protobuf:"bytes,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *Listing_Metadata) Reset()                    { *m = Listing_Metadata{} }
func (m *Listing_Metadata) String() string            { return proto.CompactTextString(m) }
func (*Listing_Metadata) ProtoMessage()               {}
func (*Listing_Metadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *Listing_Metadata) GetExpiry() *Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type Listing_Item struct {
	Title          string                  `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description    string                  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	ProcessingTime string                  `protobuf:"bytes,3,opt,name=processingTime" json:"processingTime,omitempty"`
	PricePerUnit   *Listing_Price          `protobuf:"bytes,4,opt,name=pricePerUnit" json:"pricePerUnit,omitempty"`
	Nsfw           bool                    `protobuf:"varint,5,opt,name=nsfw" json:"nsfw,omitempty"`
	Tags           []string                `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	ImageHashes    []string                `protobuf:"bytes,7,rep,name=imageHashes" json:"imageHashes,omitempty"`
	SKU            string                  `protobuf:"bytes,8,opt,name=SKU,json=sKU" json:"SKU,omitempty"`
	Condition      string                  `protobuf:"bytes,9,opt,name=condition" json:"condition,omitempty"`
	Options        []*Listing_Item_Options `protobuf:"bytes,10,rep,name=options" json:"options,omitempty"`
}

func (m *Listing_Item) Reset()                    { *m = Listing_Item{} }
func (m *Listing_Item) String() string            { return proto.CompactTextString(m) }
func (*Listing_Item) ProtoMessage()               {}
func (*Listing_Item) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func (m *Listing_Item) GetPricePerUnit() *Listing_Price {
	if m != nil {
		return m.PricePerUnit
	}
	return nil
}

func (m *Listing_Item) GetOptions() []*Listing_Item_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type Listing_Item_Options struct {
	Name            string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Values          []string       `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	PriceWithOption *Listing_Price `protobuf:"bytes,3,opt,name=priceWithOption" json:"priceWithOption,omitempty"`
}

func (m *Listing_Item_Options) Reset()                    { *m = Listing_Item_Options{} }
func (m *Listing_Item_Options) String() string            { return proto.CompactTextString(m) }
func (*Listing_Item_Options) ProtoMessage()               {}
func (*Listing_Item_Options) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1, 0} }

func (m *Listing_Item_Options) GetPriceWithOption() *Listing_Price {
	if m != nil {
		return m.PriceWithOption
	}
	return nil
}

type Listing_Shipping struct {
	FreeShipping      bool                                `protobuf:"varint,1,opt,name=freeShipping" json:"freeShipping,omitempty"`
	Domestic          *Listing_Price                      `protobuf:"bytes,2,opt,name=domestic" json:"domestic,omitempty"`
	International     *Listing_Price                      `protobuf:"bytes,3,opt,name=international" json:"international,omitempty"`
	ShippingRegions   []CountryCode                       `protobuf:"varint,4,rep,name=shippingRegions,enum=pb.CountryCode" json:"shippingRegions,omitempty"`
	EstimatedDelivery *Listing_Shipping_EstimatedDelivery `protobuf:"bytes,5,opt,name=estimatedDelivery" json:"estimatedDelivery,omitempty"`
	ShippingOrigin    CountryCode                         `protobuf:"varint,6,opt,name=shippingOrigin,enum=pb.CountryCode" json:"shippingOrigin,omitempty"`
}

func (m *Listing_Shipping) Reset()                    { *m = Listing_Shipping{} }
func (m *Listing_Shipping) String() string            { return proto.CompactTextString(m) }
func (*Listing_Shipping) ProtoMessage()               {}
func (*Listing_Shipping) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 2} }

func (m *Listing_Shipping) GetDomestic() *Listing_Price {
	if m != nil {
		return m.Domestic
	}
	return nil
}

func (m *Listing_Shipping) GetInternational() *Listing_Price {
	if m != nil {
		return m.International
	}
	return nil
}

func (m *Listing_Shipping) GetEstimatedDelivery() *Listing_Shipping_EstimatedDelivery {
	if m != nil {
		return m.EstimatedDelivery
	}
	return nil
}

type Listing_Shipping_EstimatedDelivery struct {
	Domestic      string `protobuf:"bytes,1,opt,name=domestic" json:"domestic,omitempty"`
	International string `protobuf:"bytes,2,opt,name=international" json:"international,omitempty"`
}

func (m *Listing_Shipping_EstimatedDelivery) Reset()         { *m = Listing_Shipping_EstimatedDelivery{} }
func (m *Listing_Shipping_EstimatedDelivery) String() string { return proto.CompactTextString(m) }
func (*Listing_Shipping_EstimatedDelivery) ProtoMessage()    {}
func (*Listing_Shipping_EstimatedDelivery) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 2, 0}
}

type Listing_Price struct {
	Bitcoin uint32              `protobuf:"varint,1,opt,name=bitcoin" json:"bitcoin,omitempty"`
	Fiat    *Listing_Price_Fiat `protobuf:"bytes,2,opt,name=fiat" json:"fiat,omitempty"`
}

func (m *Listing_Price) Reset()                    { *m = Listing_Price{} }
func (m *Listing_Price) String() string            { return proto.CompactTextString(m) }
func (*Listing_Price) ProtoMessage()               {}
func (*Listing_Price) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3} }

func (m *Listing_Price) GetFiat() *Listing_Price_Fiat {
	if m != nil {
		return m.Fiat
	}
	return nil
}

type Listing_Price_Fiat struct {
	CurrencyCode string  `protobuf:"bytes,1,opt,name=currencyCode" json:"currencyCode,omitempty"`
	Price        float32 `protobuf:"fixed32,2,opt,name=price" json:"price,omitempty"`
}

func (m *Listing_Price_Fiat) Reset()                    { *m = Listing_Price_Fiat{} }
func (m *Listing_Price_Fiat) String() string            { return proto.CompactTextString(m) }
func (*Listing_Price_Fiat) ProtoMessage()               {}
func (*Listing_Price_Fiat) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3, 0} }

type Order struct {
	RefundAddress string          `protobuf:"bytes,1,opt,name=refundAddress" json:"refundAddress,omitempty"`
	Shipping      *Order_Shipping `protobuf:"bytes,2,opt,name=shipping" json:"shipping,omitempty"`
	BuyerID       *ID             `protobuf:"bytes,3,opt,name=buyerID" json:"buyerID,omitempty"`
	Timestamp     *Timestamp      `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Item          []*Order_Item   `protobuf:"bytes,5,rep,name=item" json:"item,omitempty"`
	Payment       *Order_Payment  `protobuf:"bytes,6,opt,name=payment" json:"payment,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Order) GetShipping() *Order_Shipping {
	if m != nil {
		return m.Shipping
	}
	return nil
}

func (m *Order) GetBuyerID() *ID {
	if m != nil {
		return m.BuyerID
	}
	return nil
}

func (m *Order) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Order) GetItem() []*Order_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Order) GetPayment() *Order_Payment {
	if m != nil {
		return m.Payment
	}
	return nil
}

type Order_Shipping struct {
	ShipTo     string          `protobuf:"bytes,1,opt,name=shipTo" json:"shipTo,omitempty"`
	Address    string          `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	City       string          `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State      string          `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	PostalCode string          `protobuf:"bytes,5,opt,name=postalCode" json:"postalCode,omitempty"`
	Country    CountryCode     `protobuf:"varint,6,opt,name=country,enum=pb.CountryCode" json:"country,omitempty"`
}

func (m *Order_Shipping) Reset()                    { *m = Order_Shipping{} }
func (m *Order_Shipping) String() string            { return proto.CompactTextString(m) }
func (*Order_Shipping) ProtoMessage()               {}
func (*Order_Shipping) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type Order_Item struct {
	ListingHash string               `protobuf:"bytes,1,opt,name=listingHash" json:"listingHash,omitempty"`
	Quantity    uint32               `protobuf:"varint,2,opt,name=quantity" json:"quantity,omitempty"`
	Option      []*Order_Item_Option `protobuf:"bytes,3,rep,name=option" json:"option,omitempty"`
}

func (m *Order_Item) Reset()                    { *m = Order_Item{} }
func (m *Order_Item) String() string            { return proto.CompactTextString(m) }
func (*Order_Item) ProtoMessage()               {}
func (*Order_Item) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 1} }

func (m *Order_Item) GetOption() []*Order_Item_Option {
	if m != nil {
		return m.Option
	}
	return nil
}

type Order_Item_Option struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Order_Item_Option) Reset()                    { *m = Order_Item_Option{} }
func (m *Order_Item_Option) String() string            { return proto.CompactTextString(m) }
func (*Order_Item_Option) ProtoMessage()               {}
func (*Order_Item_Option) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 1, 0} }

type Order_Payment struct {
	Method       Order_Payment_Method `protobuf:"varint,1,opt,name=method,enum=pb.Order_Payment_Method" json:"method,omitempty"`
	Moderator    string               `protobuf:"bytes,2,opt,name=moderator" json:"moderator,omitempty"`
	Amount       uint32               `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Chaincode    string               `protobuf:"bytes,4,opt,name=chaincode" json:"chaincode,omitempty"`
	Address      string               `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	RedeemScript string               `protobuf:"bytes,6,opt,name=redeemScript" json:"redeemScript,omitempty"`
}

func (m *Order_Payment) Reset()                    { *m = Order_Payment{} }
func (m *Order_Payment) String() string            { return proto.CompactTextString(m) }
func (*Order_Payment) ProtoMessage()               {}
func (*Order_Payment) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 2} }

// TODO: complete other messages
type OrderConfirmation struct {
}

func (m *OrderConfirmation) Reset()                    { *m = OrderConfirmation{} }
func (m *OrderConfirmation) String() string            { return proto.CompactTextString(m) }
func (*OrderConfirmation) ProtoMessage()               {}
func (*OrderConfirmation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type Rating struct {
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type Dispute struct {
}

func (m *Dispute) Reset()                    { *m = Dispute{} }
func (m *Dispute) String() string            { return proto.CompactTextString(m) }
func (*Dispute) ProtoMessage()               {}
func (*Dispute) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DisputeResolution struct {
}

func (m *DisputeResolution) Reset()                    { *m = DisputeResolution{} }
func (m *DisputeResolution) String() string            { return proto.CompactTextString(m) }
func (*DisputeResolution) ProtoMessage()               {}
func (*DisputeResolution) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type Refund struct {
}

func (m *Refund) Reset()                    { *m = Refund{} }
func (m *Refund) String() string            { return proto.CompactTextString(m) }
func (*Refund) ProtoMessage()               {}
func (*Refund) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type Timestamp struct {
	SecondsFromUnixEpoch int64 `protobuf:"varint,1,opt,name=secondsFromUnixEpoch" json:"secondsFromUnixEpoch,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type ID struct {
	Guid         string      `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	BlockchainID string      `protobuf:"bytes,2,opt,name=blockchainID" json:"blockchainID,omitempty"`
	Pubkeys      *ID_Pubkeys `protobuf:"bytes,3,opt,name=pubkeys" json:"pubkeys,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ID) GetPubkeys() *ID_Pubkeys {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

type ID_Pubkeys struct {
	Guid    []byte `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Bitcoin []byte `protobuf:"bytes,2,opt,name=bitcoin,proto3" json:"bitcoin,omitempty"`
}

func (m *ID_Pubkeys) Reset()                    { *m = ID_Pubkeys{} }
func (m *ID_Pubkeys) String() string            { return proto.CompactTextString(m) }
func (*ID_Pubkeys) ProtoMessage()               {}
func (*ID_Pubkeys) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9, 0} }

type Signatures struct {
	Section Signatures_Section `protobuf:"varint,1,opt,name=section,enum=pb.Signatures_Section" json:"section,omitempty"`
	Guid    []byte             `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Bitcoin []byte             `protobuf:"bytes,3,opt,name=bitcoin,proto3" json:"bitcoin,omitempty"`
}

func (m *Signatures) Reset()                    { *m = Signatures{} }
func (m *Signatures) String() string            { return proto.CompactTextString(m) }
func (*Signatures) ProtoMessage()               {}
func (*Signatures) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func init() {
	proto.RegisterType((*RicardianContract)(nil), "pb.RicardianContract")
	proto.RegisterType((*Listing)(nil), "pb.Listing")
	proto.RegisterType((*Listing_Metadata)(nil), "pb.Listing.Metadata")
	proto.RegisterType((*Listing_Item)(nil), "pb.Listing.Item")
	proto.RegisterType((*Listing_Item_Options)(nil), "pb.Listing.Item.Options")
	proto.RegisterType((*Listing_Shipping)(nil), "pb.Listing.Shipping")
	proto.RegisterType((*Listing_Shipping_EstimatedDelivery)(nil), "pb.Listing.Shipping.EstimatedDelivery")
	proto.RegisterType((*Listing_Price)(nil), "pb.Listing.Price")
	proto.RegisterType((*Listing_Price_Fiat)(nil), "pb.Listing.Price.Fiat")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*Order_Shipping)(nil), "pb.Order.Shipping")
	proto.RegisterType((*Order_Item)(nil), "pb.Order.Item")
	proto.RegisterType((*Order_Item_Option)(nil), "pb.Order.Item.Option")
	proto.RegisterType((*Order_Payment)(nil), "pb.Order.Payment")
	proto.RegisterType((*OrderConfirmation)(nil), "pb.OrderConfirmation")
	proto.RegisterType((*Rating)(nil), "pb.Rating")
	proto.RegisterType((*Dispute)(nil), "pb.Dispute")
	proto.RegisterType((*DisputeResolution)(nil), "pb.DisputeResolution")
	proto.RegisterType((*Refund)(nil), "pb.Refund")
	proto.RegisterType((*Timestamp)(nil), "pb.Timestamp")
	proto.RegisterType((*ID)(nil), "pb.ID")
	proto.RegisterType((*ID_Pubkeys)(nil), "pb.ID.Pubkeys")
	proto.RegisterType((*Signatures)(nil), "pb.Signatures")
	proto.RegisterEnum("pb.Listing_Metadata_CategorySub", Listing_Metadata_CategorySub_name, Listing_Metadata_CategorySub_value)
	proto.RegisterEnum("pb.Listing_Metadata_Category", Listing_Metadata_Category_name, Listing_Metadata_Category_value)
	proto.RegisterEnum("pb.Order_Payment_Method", Order_Payment_Method_name, Order_Payment_Method_value)
	proto.RegisterEnum("pb.Signatures_Section", Signatures_Section_name, Signatures_Section_value)
}

var fileDescriptor1 = []byte{
	// 1539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x57, 0xcd, 0x92, 0xe2, 0x54,
	0x14, 0x96, 0xbf, 0x00, 0x87, 0x86, 0x6e, 0xae, 0xe3, 0x48, 0x51, 0x6a, 0x75, 0xe1, 0x38, 0x35,
	0x3a, 0x4a, 0x69, 0x4f, 0x59, 0x53, 0x96, 0x0b, 0x6d, 0x81, 0x9e, 0x49, 0x4d, 0x0f, 0x50, 0x17,
	0x70, 0x74, 0xd5, 0x15, 0x92, 0xdb, 0xf4, 0xad, 0x81, 0x24, 0x26, 0xa1, 0x1d, 0x9e, 0xc2, 0x47,
	0x70, 0xa9, 0x4b, 0x5d, 0xb9, 0x74, 0xef, 0xca, 0x87, 0xf0, 0x09, 0x7c, 0x02, 0xcf, 0xfd, 0x09,
	0x49, 0x80, 0x71, 0x97, 0xfb, 0x9d, 0x73, 0xee, 0xcf, 0x77, 0x7e, 0x03, 0xc7, 0xb6, 0xe7, 0x46,
	0x81, 0x65, 0x47, 0x61, 0xd7, 0x0f, 0xbc, 0xc8, 0x23, 0x79, 0x7f, 0xde, 0x26, 0xb6, 0xb7, 0x46,
	0x74, 0x63, 0x7b, 0x0e, 0xd3, 0x78, 0xe7, 0xb7, 0x02, 0x34, 0x29, 0xb7, 0xad, 0xc0, 0xe1, 0x96,
	0xdb, 0xd3, 0x46, 0xe4, 0x33, 0xa8, 0xdf, 0x32, 0xd7, 0xf1, 0x82, 0x4b, 0x1e, 0x46, 0xdc, 0x5d,
	0xb4, 0x72, 0xa7, 0x85, 0x07, 0xb5, 0xb3, 0x5a, 0xd7, 0x9f, 0x77, 0x35, 0x44, 0xb3, 0x1a, 0xe4,
	0x43, 0x80, 0xf9, 0x7a, 0xc3, 0x82, 0x51, 0xe0, 0xb0, 0xa0, 0x95, 0x3f, 0xcd, 0xa1, 0x7e, 0x55,
	0xe8, 0x4b, 0x80, 0xa6, 0x84, 0x64, 0x04, 0x6f, 0x2b, 0x5b, 0xb9, 0xc4, 0x43, 0xaf, 0x79, 0xb0,
	0xb2, 0x22, 0xee, 0xb9, 0xad, 0x82, 0xb4, 0x7b, 0x6b, 0x6b, 0x97, 0x16, 0xd2, 0xd7, 0x59, 0x91,
	0x8f, 0xa1, 0x26, 0xb7, 0xa7, 0x96, 0xbc, 0x6c, 0x51, 0x6e, 0x02, 0x62, 0x13, 0x85, 0xd0, 0xb4,
	0x98, 0x7c, 0x00, 0x65, 0x87, 0x87, 0xfe, 0x3a, 0x62, 0xad, 0x92, 0xd4, 0x94, 0xcf, 0xea, 0x2b,
	0x88, 0xc6, 0x32, 0xd2, 0x83, 0xa6, 0xfe, 0xa4, 0x2c, 0xf4, 0x96, 0x6b, 0x79, 0x3f, 0x23, 0xb9,
	0x5f, 0x7f, 0x57, 0x48, 0xf7, 0xf5, 0x49, 0x07, 0x8c, 0x80, 0x5d, 0xaf, 0x5d, 0xa7, 0x55, 0x4e,
	0x5d, 0x4a, 0x22, 0x54, 0x4b, 0x48, 0x17, 0x20, 0xe4, 0x0b, 0xd7, 0x8a, 0xd6, 0x01, 0x0b, 0x5b,
	0x15, 0xc9, 0x74, 0x43, 0xe8, 0x4d, 0xb6, 0x28, 0x4d, 0x69, 0x74, 0x7e, 0xad, 0x43, 0x39, 0x66,
	0xfd, 0x14, 0x6a, 0x4b, 0xf5, 0x39, 0xb4, 0x56, 0x0c, 0xdd, 0x94, 0x7b, 0x50, 0xa5, 0x69, 0x08,
	0x6f, 0x50, 0x51, 0xb4, 0x99, 0x7d, 0xed, 0x15, 0x43, 0xec, 0x6d, 0xf6, 0xe9, 0x16, 0x27, 0x9f,
	0x42, 0x65, 0xc5, 0x22, 0xcb, 0xb1, 0x22, 0x4b, 0x7b, 0xe0, 0x4e, 0xca, 0xd3, 0xdd, 0xe7, 0x5a,
	0x46, 0xb7, 0x5a, 0xe4, 0x1e, 0x14, 0x79, 0xc4, 0x56, 0x9a, 0xea, 0x93, 0xb4, 0xb6, 0x89, 0x38,
	0x95, 0x52, 0xb1, 0x6f, 0x78, 0xc3, 0x7d, 0x5f, 0x38, 0xa5, 0xb4, 0xbf, 0xef, 0x44, 0xcb, 0xe8,
	0x56, 0x8b, 0xbc, 0x07, 0xb0, 0xc2, 0xe8, 0x0c, 0xac, 0xc8, 0x0b, 0x42, 0x64, 0xbb, 0x80, 0xcf,
	0x49, 0x21, 0xc8, 0x15, 0x89, 0x58, 0xb0, 0x0a, 0xcf, 0x5d, 0x07, 0x23, 0xc0, 0xe1, 0x82, 0xe4,
	0x50, 0x72, 0x5b, 0xa5, 0x07, 0x24, 0xf8, 0xfa, 0x23, 0xc5, 0xf2, 0xd8, 0x5b, 0x72, 0x7b, 0x83,
	0xec, 0x0a, 0xcd, 0x0c, 0xd6, 0xfe, 0x2b, 0x0f, 0x95, 0xf8, 0x89, 0xa4, 0x05, 0xe5, 0x5b, 0x16,
	0x84, 0xc2, 0xd7, 0x82, 0xcc, 0x3a, 0x8d, 0x97, 0xe4, 0x0b, 0xa8, 0xd8, 0x56, 0xc4, 0x16, 0x5e,
	0xb0, 0x91, 0x44, 0x36, 0xce, 0xde, 0x3d, 0x44, 0x52, 0xb7, 0xa7, 0x95, 0xe8, 0x56, 0x9d, 0x7c,
	0x03, 0xb5, 0xf8, 0x7b, 0xb2, 0x9e, 0x4b, 0x8a, 0x1b, 0x67, 0xa7, 0xff, 0x6b, 0x8d, 0x7a, 0x34,
	0x6d, 0x84, 0x51, 0x6b, 0xb0, 0x57, 0x3e, 0xc7, 0xc3, 0x15, 0xe7, 0x75, 0x61, 0x3e, 0xe5, 0x2b,
	0x16, 0x46, 0xd6, 0xca, 0xa7, 0x5a, 0xd8, 0x79, 0x04, 0xb5, 0xd4, 0x16, 0xc4, 0x80, 0xfc, 0xf0,
	0xfc, 0xe4, 0x0d, 0x72, 0x0c, 0xb5, 0x0b, 0xf3, 0xbb, 0x41, 0xff, 0x6a, 0x4c, 0xcd, 0xde, 0xe0,
	0x24, 0x47, 0x6a, 0x50, 0x3e, 0x9f, 0xf5, 0xa6, 0xe6, 0x68, 0x78, 0x92, 0xef, 0x98, 0x50, 0x89,
	0x8d, 0x84, 0x60, 0x36, 0x7c, 0x36, 0x1c, 0xbd, 0x18, 0xa2, 0x59, 0x13, 0xea, 0xe3, 0xa7, 0xdf,
	0x4f, 0xcc, 0xde, 0xf9, 0xe5, 0xd5, 0x93, 0xd1, 0xa8, 0x8f, 0x86, 0x27, 0x70, 0xd4, 0x37, 0x9f,
	0x98, 0xd3, 0x18, 0xc9, 0x0b, 0x8b, 0xc9, 0x80, 0x7e, 0x2b, 0xf6, 0x2d, 0xb4, 0xff, 0x2c, 0x40,
	0x51, 0x44, 0x00, 0xb9, 0x03, 0xa5, 0x88, 0x47, 0xcb, 0x38, 0x26, 0xd5, 0x42, 0xc4, 0x2b, 0xd6,
	0x1e, 0x3b, 0xe0, 0xbe, 0x4c, 0xa7, 0xbc, 0x8a, 0xd7, 0x14, 0x44, 0xee, 0x43, 0x03, 0x2b, 0x93,
	0xcd, 0xc2, 0x10, 0x79, 0x11, 0xef, 0x93, 0x74, 0x55, 0xe9, 0x0e, 0x4a, 0x3e, 0x87, 0x23, 0x3f,
	0xe0, 0x36, 0x1b, 0xb3, 0x60, 0xe6, 0xf2, 0x48, 0xb3, 0xd2, 0x4c, 0x93, 0x3a, 0x16, 0x72, 0x9a,
	0x51, 0x23, 0x04, 0x8a, 0x6e, 0x78, 0xfd, 0xa3, 0x0c, 0xc7, 0x0a, 0x95, 0xdf, 0x02, 0x8b, 0xac,
	0x45, 0x1c, 0x6e, 0xf2, 0x5b, 0x5c, 0x94, 0xaf, 0xac, 0x05, 0x7b, 0x6a, 0x85, 0x37, 0x4c, 0x44,
	0x98, 0x10, 0xa5, 0x21, 0x24, 0xa2, 0x30, 0x79, 0x36, 0xd3, 0x11, 0x55, 0x08, 0x9f, 0xcd, 0xc8,
	0x3b, 0x50, 0xb5, 0xe3, 0xd0, 0x6b, 0x55, 0x25, 0x9e, 0x00, 0xe4, 0x0c, 0xca, 0x9e, 0xaf, 0xe2,
	0x15, 0x64, 0x8e, 0xb7, 0x76, 0xb3, 0xa6, 0x3b, 0x52, 0x72, 0x1a, 0x2b, 0xb6, 0x03, 0x28, 0x6b,
	0x4c, 0x5e, 0x3c, 0x49, 0x71, 0xf9, 0x4d, 0xee, 0x82, 0x71, 0x6b, 0x2d, 0xd7, 0x78, 0xbf, 0xbc,
	0xbc, 0x9f, 0x5e, 0x91, 0x2f, 0xe1, 0x58, 0x3e, 0xfa, 0x05, 0x8f, 0x6e, 0x94, 0xbd, 0x4e, 0xeb,
	0x03, 0xf4, 0xec, 0x6a, 0xb6, 0xff, 0x28, 0x40, 0x25, 0xce, 0x4c, 0x91, 0x3f, 0xd7, 0x01, 0x63,
	0xf1, 0x5a, 0x9e, 0x5e, 0xa1, 0x19, 0x8c, 0x7c, 0x02, 0x15, 0xc7, 0x13, 0x71, 0xc8, 0x6d, 0x5d,
	0x61, 0x0e, 0x1c, 0xb3, 0x55, 0x21, 0x8f, 0xa1, 0xce, 0x5d, 0x4c, 0x55, 0x57, 0xd6, 0x6e, 0x6b,
	0xf9, 0xfa, 0xab, 0x65, 0xf5, 0x30, 0x01, 0x8f, 0xe3, 0x3a, 0x41, 0xd9, 0x42, 0x12, 0x59, 0xc4,
	0x67, 0x37, 0xce, 0x8e, 0x85, 0x69, 0x4f, 0xf5, 0xb6, 0x1e, 0xd6, 0x0a, 0xba, 0xab, 0x47, 0xa6,
	0xd0, 0x14, 0x87, 0x63, 0xbb, 0x60, 0x4e, 0x9f, 0x2d, 0x39, 0xe6, 0xf4, 0x46, 0x57, 0xa4, 0xfb,
	0x87, 0x2a, 0x52, 0x77, 0xb0, 0xab, 0x4d, 0xf7, 0x37, 0xc0, 0x97, 0x34, 0xe2, 0x83, 0x46, 0x01,
	0x5f, 0x70, 0xd5, 0x1e, 0x0e, 0xdc, 0x67, 0x47, 0xad, 0x3d, 0x83, 0xe6, 0xde, 0x01, 0xa4, 0x9d,
	0xa2, 0x51, 0x39, 0x39, 0xe1, 0xec, 0xde, 0x2e, 0x67, 0x2a, 0x71, 0xb2, 0x60, 0xfb, 0xa7, 0x1c,
	0x94, 0x24, 0x73, 0xa2, 0x8a, 0xcd, 0x79, 0x64, 0x7b, 0x7c, 0x5b, 0xc5, 0xf4, 0x92, 0x7c, 0x04,
	0xc5, 0x6b, 0x6e, 0x45, 0xda, 0x51, 0x77, 0xf7, 0x48, 0xef, 0x5e, 0xa0, 0x94, 0x4a, 0x9d, 0xf6,
	0xd7, 0x50, 0x14, 0x2b, 0x11, 0x04, 0xf6, 0x3a, 0x08, 0x98, 0x6b, 0xcb, 0xe7, 0xe8, 0xdb, 0x65,
	0x30, 0x91, 0xee, 0x32, 0x90, 0xe4, 0xc6, 0x79, 0xaa, 0x16, 0x9d, 0xbf, 0x0d, 0x28, 0xa9, 0x9e,
	0x8f, 0x2f, 0x50, 0x45, 0xf7, 0xdc, 0x71, 0xb0, 0x89, 0x85, 0x7a, 0x93, 0x2c, 0x88, 0xe5, 0x3d,
	0x69, 0x18, 0xea, 0x86, 0x64, 0x3b, 0x0a, 0x1c, 0x6a, 0x17, 0xa7, 0xf8, 0x4e, 0xd1, 0xd9, 0xb1,
	0xb7, 0x15, 0x32, 0xbd, 0x2d, 0x86, 0xc9, 0x43, 0xa8, 0x46, 0x71, 0x91, 0x3c, 0x5c, 0x39, 0x13,
	0x39, 0x3e, 0x54, 0x75, 0xb5, 0x52, 0xd2, 0x83, 0xd5, 0xd1, 0xa9, 0x9e, 0xf6, 0x10, 0xca, 0xbe,
	0xb5, 0x59, 0x31, 0x37, 0xd2, 0xc3, 0x40, 0x33, 0x51, 0x1b, 0x2b, 0x01, 0x8d, 0x35, 0xda, 0xbf,
	0xe7, 0x52, 0xb9, 0x84, 0xd9, 0x2a, 0x2e, 0x3e, 0xf5, 0xf4, 0xdb, 0xf5, 0x4a, 0x38, 0xcb, 0xd2,
	0xa4, 0x28, 0xb7, 0xc6, 0x4b, 0x91, 0xf3, 0x36, 0x8f, 0x36, 0xba, 0x02, 0xca, 0x6f, 0x41, 0x34,
	0x5e, 0x16, 0x67, 0x97, 0xa2, 0xaa, 0xab, 0x72, 0x21, 0xfa, 0xa6, 0xef, 0xe1, 0xe7, 0x52, 0x3a,
	0xa8, 0x24, 0x45, 0x29, 0x04, 0xa7, 0xb3, 0xb2, 0x1e, 0xfe, 0x5e, 0x17, 0xa3, 0xb1, 0xbc, 0xfd,
	0x4b, 0x4e, 0x57, 0xf0, 0x64, 0xb6, 0x10, 0x15, 0x6f, 0x67, 0xb6, 0x10, 0x90, 0x08, 0xd9, 0x1f,
	0xd6, 0x96, 0x1b, 0x89, 0x3b, 0xe6, 0x65, 0x9c, 0x6d, 0xd7, 0x58, 0x15, 0x0c, 0x2f, 0x2e, 0x3d,
	0x85, 0xcc, 0x4c, 0x97, 0xae, 0x75, 0x54, 0x2b, 0xb5, 0xcf, 0xc0, 0x50, 0xc8, 0xc1, 0x42, 0x87,
	0x8f, 0x96, 0xa5, 0x4d, 0x13, 0xa4, 0x16, 0xed, 0x7f, 0x73, 0x50, 0xd6, 0x94, 0xe3, 0xa8, 0x61,
	0xe0, 0x70, 0x72, 0xe3, 0x39, 0xd2, 0xae, 0xa1, 0x8a, 0x6b, 0xc6, 0x2b, 0xa2, 0xc7, 0xa2, 0x9c,
	0x6a, 0x3d, 0x51, 0xad, 0xb7, 0x83, 0x85, 0xde, 0x37, 0x01, 0x84, 0xb3, 0xac, 0x95, 0x60, 0x44,
	0x92, 0x5f, 0xa7, 0x7a, 0x25, 0x6b, 0xfc, 0x8d, 0xc5, 0x5d, 0x31, 0x43, 0x6b, 0x17, 0x24, 0x40,
	0xda, 0x95, 0xa5, 0xac, 0x2b, 0xe5, 0x20, 0xe2, 0x30, 0xb6, 0x9a, 0xc8, 0x4e, 0x27, 0xbd, 0x20,
	0x07, 0x91, 0x04, 0xeb, 0xbc, 0x0f, 0x86, 0xba, 0x23, 0x01, 0x30, 0xfa, 0x26, 0x1d, 0xf4, 0xa6,
	0xd8, 0x83, 0xeb, 0x50, 0x7d, 0x3e, 0xea, 0x0f, 0xe8, 0xf9, 0x74, 0x80, 0xfd, 0xb7, 0xf3, 0x26,
	0x34, 0xf7, 0x06, 0xe0, 0x4e, 0x05, 0x0c, 0x35, 0xdc, 0x76, 0xaa, 0x50, 0xd6, 0x83, 0xa9, 0xd0,
	0xdc, 0x9b, 0x51, 0xa5, 0xa6, 0x4c, 0xb9, 0xce, 0x57, 0x50, 0xdd, 0x26, 0x01, 0x36, 0xa7, 0x3b,
	0x21, 0x13, 0xbd, 0x2a, 0xbc, 0x08, 0xbc, 0x15, 0x76, 0xca, 0x57, 0x03, 0xdf, 0xb3, 0x95, 0xd3,
	0x0b, 0xf4, 0xa0, 0xac, 0xf3, 0x73, 0x0e, 0xf2, 0x98, 0x61, 0xe8, 0xaf, 0xc5, 0x9a, 0x3b, 0xb1,
	0xbf, 0xc4, 0xb7, 0x78, 0xed, 0x7c, 0xe9, 0xd9, 0x2f, 0x25, 0x33, 0x7a, 0xf0, 0xc4, 0xd7, 0xa6,
	0x31, 0xf2, 0x00, 0x13, 0x69, 0x3d, 0x7f, 0xc9, 0x36, 0xa1, 0xce, 0xdd, 0x86, 0xca, 0xdd, 0xee,
	0x58, 0xa1, 0x34, 0x16, 0xb7, 0x1f, 0xa3, 0x9b, 0xd5, 0x67, 0xe6, 0xb0, 0x23, 0x7d, 0x58, 0xaa,
	0xd8, 0xe5, 0x25, 0x1c, 0x2f, 0x3b, 0xff, 0xe4, 0x00, 0x92, 0x21, 0x1a, 0x63, 0xa4, 0x8c, 0x0f,
	0x89, 0xe2, 0xd9, 0xae, 0xa1, 0xca, 0x5f, 0xa2, 0xd0, 0x9d, 0x28, 0x29, 0x8d, 0xd5, 0xb6, 0xc7,
	0xe5, 0x0f, 0x1f, 0x57, 0xc8, 0x1e, 0x77, 0x8b, 0x83, 0x90, 0x36, 0xc4, 0x99, 0xe8, 0xd2, 0x9c,
	0x4c, 0xcd, 0xe1, 0x13, 0xf4, 0x60, 0x15, 0x8b, 0x20, 0x45, 0x0f, 0xe2, 0xf4, 0x74, 0x17, 0x88,
	0xfc, 0xbc, 0xea, 0x8d, 0x86, 0x17, 0x26, 0x7d, 0x7e, 0xae, 0x26, 0x30, 0xe1, 0x70, 0x74, 0xb0,
	0x50, 0x2f, 0x08, 0xdb, 0xbe, 0x39, 0x19, 0xcf, 0xa6, 0x83, 0x93, 0xa2, 0x30, 0xd0, 0x8b, 0x2b,
	0x3a, 0x98, 0x8c, 0x2e, 0x67, 0xd2, 0xa0, 0x24, 0x0d, 0x06, 0x17, 0xb3, 0x61, 0xff, 0xc4, 0x98,
	0x1b, 0xf2, 0x57, 0xee, 0xd1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x06, 0x7a, 0xf9, 0xf5,
	0x0d, 0x00, 0x00,
}