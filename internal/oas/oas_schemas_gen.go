// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"
	"time"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Address
type Address struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	LineOne    string `json:"line_one"`
	LineTwo    string `json:"line_two"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

// GetID returns the value of ID.
func (s *Address) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Address) GetName() string {
	return s.Name
}

// GetLineOne returns the value of LineOne.
func (s *Address) GetLineOne() string {
	return s.LineOne
}

// GetLineTwo returns the value of LineTwo.
func (s *Address) GetLineTwo() string {
	return s.LineTwo
}

// GetCity returns the value of City.
func (s *Address) GetCity() string {
	return s.City
}

// GetState returns the value of State.
func (s *Address) GetState() string {
	return s.State
}

// GetPostalCode returns the value of PostalCode.
func (s *Address) GetPostalCode() string {
	return s.PostalCode
}

// GetCountry returns the value of Country.
func (s *Address) GetCountry() string {
	return s.Country
}

// SetID sets the value of ID.
func (s *Address) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Address) SetName(val string) {
	s.Name = val
}

// SetLineOne sets the value of LineOne.
func (s *Address) SetLineOne(val string) {
	s.LineOne = val
}

// SetLineTwo sets the value of LineTwo.
func (s *Address) SetLineTwo(val string) {
	s.LineTwo = val
}

// SetCity sets the value of City.
func (s *Address) SetCity(val string) {
	s.City = val
}

// SetState sets the value of State.
func (s *Address) SetState(val string) {
	s.State = val
}

// SetPostalCode sets the value of PostalCode.
func (s *Address) SetPostalCode(val string) {
	s.PostalCode = val
}

// SetCountry sets the value of Country.
func (s *Address) SetCountry(val string) {
	s.Country = val
}

// Ref: #/components/schemas/Aggregation
type Aggregation map[string][]Bucket

func (s *Aggregation) init() Aggregation {
	m := *s
	if m == nil {
		m = map[string][]Bucket{}
		*s = m
	}
	return m
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/Branch
type Branch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *Branch) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Branch) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *Branch) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Branch) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/Bucket
type Bucket struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// GetValue returns the value of Value.
func (s *Bucket) GetValue() string {
	return s.Value
}

// GetCount returns the value of Count.
func (s *Bucket) GetCount() int {
	return s.Count
}

// SetValue sets the value of Value.
func (s *Bucket) SetValue(val string) {
	s.Value = val
}

// SetCount sets the value of Count.
func (s *Bucket) SetCount(val int) {
	s.Count = val
}

type CreateQuoteCreated struct {
	// ID of the created quote.
	QuoteID string `json:"quote_id"`
	// Status of the quote creation.
	Status OptString `json:"status"`
}

// GetQuoteID returns the value of QuoteID.
func (s *CreateQuoteCreated) GetQuoteID() string {
	return s.QuoteID
}

// GetStatus returns the value of Status.
func (s *CreateQuoteCreated) GetStatus() OptString {
	return s.Status
}

// SetQuoteID sets the value of QuoteID.
func (s *CreateQuoteCreated) SetQuoteID(val string) {
	s.QuoteID = val
}

// SetStatus sets the value of Status.
func (s *CreateQuoteCreated) SetStatus(val OptString) {
	s.Status = val
}

func (*CreateQuoteCreated) createQuoteRes() {}

type CreateQuoteReq struct {
	PurchaseOrder        string `json:"purchase_order"`
	DeliveryInstructions string `json:"delivery_instructions"`
	// Date of the request.  Should be in YYYY-MM-DD format (ISO 8601).
	DateRequested time.Time `json:"date_requested"`
	// List of items for the quote.
	Items []CreateQuoteReqItemsItem `json:"items"`
}

// GetPurchaseOrder returns the value of PurchaseOrder.
func (s *CreateQuoteReq) GetPurchaseOrder() string {
	return s.PurchaseOrder
}

// GetDeliveryInstructions returns the value of DeliveryInstructions.
func (s *CreateQuoteReq) GetDeliveryInstructions() string {
	return s.DeliveryInstructions
}

// GetDateRequested returns the value of DateRequested.
func (s *CreateQuoteReq) GetDateRequested() time.Time {
	return s.DateRequested
}

// GetItems returns the value of Items.
func (s *CreateQuoteReq) GetItems() []CreateQuoteReqItemsItem {
	return s.Items
}

// SetPurchaseOrder sets the value of PurchaseOrder.
func (s *CreateQuoteReq) SetPurchaseOrder(val string) {
	s.PurchaseOrder = val
}

// SetDeliveryInstructions sets the value of DeliveryInstructions.
func (s *CreateQuoteReq) SetDeliveryInstructions(val string) {
	s.DeliveryInstructions = val
}

// SetDateRequested sets the value of DateRequested.
func (s *CreateQuoteReq) SetDateRequested(val time.Time) {
	s.DateRequested = val
}

// SetItems sets the value of Items.
func (s *CreateQuoteReq) SetItems(val []CreateQuoteReqItemsItem) {
	s.Items = val
}

type CreateQuoteReqItemsItem struct {
	// ID of the product.
	ProductID string `json:"product_id"`
	// Quantity of the product.
	Quantity float64 `json:"quantity"`
}

// GetProductID returns the value of ProductID.
func (s *CreateQuoteReqItemsItem) GetProductID() string {
	return s.ProductID
}

// GetQuantity returns the value of Quantity.
func (s *CreateQuoteReqItemsItem) GetQuantity() float64 {
	return s.Quantity
}

// SetProductID sets the value of ProductID.
func (s *CreateQuoteReqItemsItem) SetProductID(val string) {
	s.ProductID = val
}

// SetQuantity sets the value of Quantity.
func (s *CreateQuoteReqItemsItem) SetQuantity(val float64) {
	s.Quantity = val
}

type CreateQuoteUnprocessableEntity struct {
	// List of validation errors.
	Errors []CreateQuoteUnprocessableEntityErrorsItem `json:"errors"`
}

// GetErrors returns the value of Errors.
func (s *CreateQuoteUnprocessableEntity) GetErrors() []CreateQuoteUnprocessableEntityErrorsItem {
	return s.Errors
}

// SetErrors sets the value of Errors.
func (s *CreateQuoteUnprocessableEntity) SetErrors(val []CreateQuoteUnprocessableEntityErrorsItem) {
	s.Errors = val
}

func (*CreateQuoteUnprocessableEntity) createQuoteRes() {}

type CreateQuoteUnprocessableEntityErrorsItem struct {
	// The field where the validation error occurred.
	Field OptString `json:"field"`
	// A descriptive error message.
	Message OptString `json:"message"`
}

// GetField returns the value of Field.
func (s *CreateQuoteUnprocessableEntityErrorsItem) GetField() OptString {
	return s.Field
}

// GetMessage returns the value of Message.
func (s *CreateQuoteUnprocessableEntityErrorsItem) GetMessage() OptString {
	return s.Message
}

// SetField sets the value of Field.
func (s *CreateQuoteUnprocessableEntityErrorsItem) SetField(val OptString) {
	s.Field = val
}

// SetMessage sets the value of Message.
func (s *CreateQuoteUnprocessableEntityErrorsItem) SetMessage(val OptString) {
	s.Message = val
}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// GetOrderNotFound is response for GetOrder operation.
type GetOrderNotFound struct{}

func (*GetOrderNotFound) getOrderRes() {}

type GetOrderOK struct {
	Order Order `json:"order"`
}

// GetOrder returns the value of Order.
func (s *GetOrderOK) GetOrder() Order {
	return s.Order
}

// SetOrder sets the value of Order.
func (s *GetOrderOK) SetOrder(val Order) {
	s.Order = val
}

func (*GetOrderOK) getOrderRes() {}

// GetOrderUnauthorized is response for GetOrder operation.
type GetOrderUnauthorized struct{}

func (*GetOrderUnauthorized) getOrderRes() {}

// GetProductNotFound is response for GetProduct operation.
type GetProductNotFound struct{}

func (*GetProductNotFound) getProductRes() {}

type GetProductOK struct {
	Product Product `json:"product"`
}

// GetProduct returns the value of Product.
func (s *GetProductOK) GetProduct() Product {
	return s.Product
}

// SetProduct sets the value of Product.
func (s *GetProductOK) SetProduct(val Product) {
	s.Product = val
}

func (*GetProductOK) getProductRes() {}

type GetQuoteOK struct {
	Quote Quote `json:"quote"`
}

// GetQuote returns the value of Quote.
func (s *GetQuoteOK) GetQuote() Quote {
	return s.Quote
}

// SetQuote sets the value of Quote.
func (s *GetQuoteOK) SetQuote(val Quote) {
	s.Quote = val
}

type ListCustomerBranchesOK struct {
	Branches []Branch `json:"branches"`
}

// GetBranches returns the value of Branches.
func (s *ListCustomerBranchesOK) GetBranches() []Branch {
	return s.Branches
}

// SetBranches sets the value of Branches.
func (s *ListCustomerBranchesOK) SetBranches(val []Branch) {
	s.Branches = val
}

type ListOrdersOK struct {
	TotalRecords int            `json:"total_records"`
	Orders       []OrderSummary `json:"orders"`
}

// GetTotalRecords returns the value of TotalRecords.
func (s *ListOrdersOK) GetTotalRecords() int {
	return s.TotalRecords
}

// GetOrders returns the value of Orders.
func (s *ListOrdersOK) GetOrders() []OrderSummary {
	return s.Orders
}

// SetTotalRecords sets the value of TotalRecords.
func (s *ListOrdersOK) SetTotalRecords(val int) {
	s.TotalRecords = val
}

// SetOrders sets the value of Orders.
func (s *ListOrdersOK) SetOrders(val []OrderSummary) {
	s.Orders = val
}

type ListQuotesOK struct {
	TotalRecords int            `json:"total_records"`
	Quotes       []QuoteSummary `json:"quotes"`
}

// GetTotalRecords returns the value of TotalRecords.
func (s *ListQuotesOK) GetTotalRecords() int {
	return s.TotalRecords
}

// GetQuotes returns the value of Quotes.
func (s *ListQuotesOK) GetQuotes() []QuoteSummary {
	return s.Quotes
}

// SetTotalRecords sets the value of TotalRecords.
func (s *ListQuotesOK) SetTotalRecords(val int) {
	s.TotalRecords = val
}

// SetQuotes sets the value of Quotes.
func (s *ListQuotesOK) SetQuotes(val []QuoteSummary) {
	s.Quotes = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSearchProductsReqFilters returns new OptSearchProductsReqFilters with value set to v.
func NewOptSearchProductsReqFilters(v SearchProductsReqFilters) OptSearchProductsReqFilters {
	return OptSearchProductsReqFilters{
		Value: v,
		Set:   true,
	}
}

// OptSearchProductsReqFilters is optional SearchProductsReqFilters.
type OptSearchProductsReqFilters struct {
	Value SearchProductsReqFilters
	Set   bool
}

// IsSet returns true if OptSearchProductsReqFilters was set.
func (o OptSearchProductsReqFilters) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSearchProductsReqFilters) Reset() {
	var v SearchProductsReqFilters
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSearchProductsReqFilters) SetTo(v SearchProductsReqFilters) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSearchProductsReqFilters) Get() (v SearchProductsReqFilters, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSearchProductsReqFilters) Or(d SearchProductsReqFilters) SearchProductsReqFilters {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Order
type Order struct {
	ID                   string      `json:"id"`
	ContactID            string      `json:"contact_id"`
	BranchID             string      `json:"branch_id"`
	PurchaseOrder        string      `json:"purchase_order"`
	Status               OrderStatus `json:"status"`
	DateCreated          time.Time   `json:"date_created"`
	DateRequested        time.Time   `json:"date_requested"`
	Taker                OptString   `json:"taker"`
	DeliveryInstructions string      `json:"delivery_instructions"`
	ShippingAddress      Address     `json:"shipping_address"`
	Total                float64     `json:"total"`
	Items                []OrderItem `json:"items"`
}

// GetID returns the value of ID.
func (s *Order) GetID() string {
	return s.ID
}

// GetContactID returns the value of ContactID.
func (s *Order) GetContactID() string {
	return s.ContactID
}

// GetBranchID returns the value of BranchID.
func (s *Order) GetBranchID() string {
	return s.BranchID
}

// GetPurchaseOrder returns the value of PurchaseOrder.
func (s *Order) GetPurchaseOrder() string {
	return s.PurchaseOrder
}

// GetStatus returns the value of Status.
func (s *Order) GetStatus() OrderStatus {
	return s.Status
}

// GetDateCreated returns the value of DateCreated.
func (s *Order) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateRequested returns the value of DateRequested.
func (s *Order) GetDateRequested() time.Time {
	return s.DateRequested
}

// GetTaker returns the value of Taker.
func (s *Order) GetTaker() OptString {
	return s.Taker
}

// GetDeliveryInstructions returns the value of DeliveryInstructions.
func (s *Order) GetDeliveryInstructions() string {
	return s.DeliveryInstructions
}

// GetShippingAddress returns the value of ShippingAddress.
func (s *Order) GetShippingAddress() Address {
	return s.ShippingAddress
}

// GetTotal returns the value of Total.
func (s *Order) GetTotal() float64 {
	return s.Total
}

// GetItems returns the value of Items.
func (s *Order) GetItems() []OrderItem {
	return s.Items
}

// SetID sets the value of ID.
func (s *Order) SetID(val string) {
	s.ID = val
}

// SetContactID sets the value of ContactID.
func (s *Order) SetContactID(val string) {
	s.ContactID = val
}

// SetBranchID sets the value of BranchID.
func (s *Order) SetBranchID(val string) {
	s.BranchID = val
}

// SetPurchaseOrder sets the value of PurchaseOrder.
func (s *Order) SetPurchaseOrder(val string) {
	s.PurchaseOrder = val
}

// SetStatus sets the value of Status.
func (s *Order) SetStatus(val OrderStatus) {
	s.Status = val
}

// SetDateCreated sets the value of DateCreated.
func (s *Order) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateRequested sets the value of DateRequested.
func (s *Order) SetDateRequested(val time.Time) {
	s.DateRequested = val
}

// SetTaker sets the value of Taker.
func (s *Order) SetTaker(val OptString) {
	s.Taker = val
}

// SetDeliveryInstructions sets the value of DeliveryInstructions.
func (s *Order) SetDeliveryInstructions(val string) {
	s.DeliveryInstructions = val
}

// SetShippingAddress sets the value of ShippingAddress.
func (s *Order) SetShippingAddress(val Address) {
	s.ShippingAddress = val
}

// SetTotal sets the value of Total.
func (s *Order) SetTotal(val float64) {
	s.Total = val
}

// SetItems sets the value of Items.
func (s *Order) SetItems(val []OrderItem) {
	s.Items = val
}

// Ref: #/components/schemas/OrderItem
type OrderItem struct {
	ProductSn           string  `json:"product_sn"`
	ProductName         string  `json:"product_name"`
	ProductID           string  `json:"product_id"`
	CustomerProductSn   string  `json:"customer_product_sn"`
	OrderedQuantity     float64 `json:"ordered_quantity"`
	ShippedQuantity     float64 `json:"shipped_quantity"`
	UnitType            string  `json:"unit_type"`
	UnitPrice           float64 `json:"unit_price"`
	TotalPrice          float64 `json:"total_price"`
	BackOrderedQuantity float64 `json:"back_ordered_quantity"`
}

// GetProductSn returns the value of ProductSn.
func (s *OrderItem) GetProductSn() string {
	return s.ProductSn
}

// GetProductName returns the value of ProductName.
func (s *OrderItem) GetProductName() string {
	return s.ProductName
}

// GetProductID returns the value of ProductID.
func (s *OrderItem) GetProductID() string {
	return s.ProductID
}

// GetCustomerProductSn returns the value of CustomerProductSn.
func (s *OrderItem) GetCustomerProductSn() string {
	return s.CustomerProductSn
}

// GetOrderedQuantity returns the value of OrderedQuantity.
func (s *OrderItem) GetOrderedQuantity() float64 {
	return s.OrderedQuantity
}

// GetShippedQuantity returns the value of ShippedQuantity.
func (s *OrderItem) GetShippedQuantity() float64 {
	return s.ShippedQuantity
}

// GetUnitType returns the value of UnitType.
func (s *OrderItem) GetUnitType() string {
	return s.UnitType
}

// GetUnitPrice returns the value of UnitPrice.
func (s *OrderItem) GetUnitPrice() float64 {
	return s.UnitPrice
}

// GetTotalPrice returns the value of TotalPrice.
func (s *OrderItem) GetTotalPrice() float64 {
	return s.TotalPrice
}

// GetBackOrderedQuantity returns the value of BackOrderedQuantity.
func (s *OrderItem) GetBackOrderedQuantity() float64 {
	return s.BackOrderedQuantity
}

// SetProductSn sets the value of ProductSn.
func (s *OrderItem) SetProductSn(val string) {
	s.ProductSn = val
}

// SetProductName sets the value of ProductName.
func (s *OrderItem) SetProductName(val string) {
	s.ProductName = val
}

// SetProductID sets the value of ProductID.
func (s *OrderItem) SetProductID(val string) {
	s.ProductID = val
}

// SetCustomerProductSn sets the value of CustomerProductSn.
func (s *OrderItem) SetCustomerProductSn(val string) {
	s.CustomerProductSn = val
}

// SetOrderedQuantity sets the value of OrderedQuantity.
func (s *OrderItem) SetOrderedQuantity(val float64) {
	s.OrderedQuantity = val
}

// SetShippedQuantity sets the value of ShippedQuantity.
func (s *OrderItem) SetShippedQuantity(val float64) {
	s.ShippedQuantity = val
}

// SetUnitType sets the value of UnitType.
func (s *OrderItem) SetUnitType(val string) {
	s.UnitType = val
}

// SetUnitPrice sets the value of UnitPrice.
func (s *OrderItem) SetUnitPrice(val float64) {
	s.UnitPrice = val
}

// SetTotalPrice sets the value of TotalPrice.
func (s *OrderItem) SetTotalPrice(val float64) {
	s.TotalPrice = val
}

// SetBackOrderedQuantity sets the value of BackOrderedQuantity.
func (s *OrderItem) SetBackOrderedQuantity(val float64) {
	s.BackOrderedQuantity = val
}

// Ref: #/components/schemas/OrderStatus
type OrderStatus string

const (
	OrderStatusUnspecified     OrderStatus = "unspecified"
	OrderStatusApproved        OrderStatus = "approved"
	OrderStatusCompleted       OrderStatus = "completed"
	OrderStatusCancelled       OrderStatus = "cancelled"
	OrderStatusPendingApproval OrderStatus = "pending_approval"
)

// AllValues returns all OrderStatus values.
func (OrderStatus) AllValues() []OrderStatus {
	return []OrderStatus{
		OrderStatusUnspecified,
		OrderStatusApproved,
		OrderStatusCompleted,
		OrderStatusCancelled,
		OrderStatusPendingApproval,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s OrderStatus) MarshalText() ([]byte, error) {
	switch s {
	case OrderStatusUnspecified:
		return []byte(s), nil
	case OrderStatusApproved:
		return []byte(s), nil
	case OrderStatusCompleted:
		return []byte(s), nil
	case OrderStatusCancelled:
		return []byte(s), nil
	case OrderStatusPendingApproval:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *OrderStatus) UnmarshalText(data []byte) error {
	switch OrderStatus(data) {
	case OrderStatusUnspecified:
		*s = OrderStatusUnspecified
		return nil
	case OrderStatusApproved:
		*s = OrderStatusApproved
		return nil
	case OrderStatusCompleted:
		*s = OrderStatusCompleted
		return nil
	case OrderStatusCancelled:
		*s = OrderStatusCancelled
		return nil
	case OrderStatusPendingApproval:
		*s = OrderStatusPendingApproval
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/OrderSummary
type OrderSummary struct {
	ID            string      `json:"id"`
	ContactID     string      `json:"contact_id"`
	BranchID      string      `json:"branch_id"`
	PurchaseOrder string      `json:"purchase_order"`
	Status        OrderStatus `json:"status"`
	DateCreated   time.Time   `json:"date_created"`
	DateRequested time.Time   `json:"date_requested"`
}

// GetID returns the value of ID.
func (s *OrderSummary) GetID() string {
	return s.ID
}

// GetContactID returns the value of ContactID.
func (s *OrderSummary) GetContactID() string {
	return s.ContactID
}

// GetBranchID returns the value of BranchID.
func (s *OrderSummary) GetBranchID() string {
	return s.BranchID
}

// GetPurchaseOrder returns the value of PurchaseOrder.
func (s *OrderSummary) GetPurchaseOrder() string {
	return s.PurchaseOrder
}

// GetStatus returns the value of Status.
func (s *OrderSummary) GetStatus() OrderStatus {
	return s.Status
}

// GetDateCreated returns the value of DateCreated.
func (s *OrderSummary) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateRequested returns the value of DateRequested.
func (s *OrderSummary) GetDateRequested() time.Time {
	return s.DateRequested
}

// SetID sets the value of ID.
func (s *OrderSummary) SetID(val string) {
	s.ID = val
}

// SetContactID sets the value of ContactID.
func (s *OrderSummary) SetContactID(val string) {
	s.ContactID = val
}

// SetBranchID sets the value of BranchID.
func (s *OrderSummary) SetBranchID(val string) {
	s.BranchID = val
}

// SetPurchaseOrder sets the value of PurchaseOrder.
func (s *OrderSummary) SetPurchaseOrder(val string) {
	s.PurchaseOrder = val
}

// SetStatus sets the value of Status.
func (s *OrderSummary) SetStatus(val OrderStatus) {
	s.Status = val
}

// SetDateCreated sets the value of DateCreated.
func (s *OrderSummary) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateRequested sets the value of DateRequested.
func (s *OrderSummary) SetDateRequested(val time.Time) {
	s.DateRequested = val
}

// Ref: #/components/schemas/PageMetadata
type PageMetadata struct {
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

// GetTotalPages returns the value of TotalPages.
func (s *PageMetadata) GetTotalPages() int {
	return s.TotalPages
}

// GetTotalRecords returns the value of TotalRecords.
func (s *PageMetadata) GetTotalRecords() int {
	return s.TotalRecords
}

// SetTotalPages sets the value of TotalPages.
func (s *PageMetadata) SetTotalPages(val int) {
	s.TotalPages = val
}

// SetTotalRecords sets the value of TotalRecords.
func (s *PageMetadata) SetTotalRecords(val int) {
	s.TotalRecords = val
}

// Ref: #/components/schemas/Product
type Product struct {
	ID          string    `json:"id"`
	Sn          string    `json:"sn"`
	Name        string    `json:"name"`
	Description OptString `json:"description"`
	ImageURL    OptString `json:"image_url"`
}

// GetID returns the value of ID.
func (s *Product) GetID() string {
	return s.ID
}

// GetSn returns the value of Sn.
func (s *Product) GetSn() string {
	return s.Sn
}

// GetName returns the value of Name.
func (s *Product) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Product) GetDescription() OptString {
	return s.Description
}

// GetImageURL returns the value of ImageURL.
func (s *Product) GetImageURL() OptString {
	return s.ImageURL
}

// SetID sets the value of ID.
func (s *Product) SetID(val string) {
	s.ID = val
}

// SetSn sets the value of Sn.
func (s *Product) SetSn(val string) {
	s.Sn = val
}

// SetName sets the value of Name.
func (s *Product) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Product) SetDescription(val OptString) {
	s.Description = val
}

// SetImageURL sets the value of ImageURL.
func (s *Product) SetImageURL(val OptString) {
	s.ImageURL = val
}

// Ref: #/components/schemas/Quote
type Quote struct {
	ID            string      `json:"id"`
	PurchaseOrder string      `json:"purchase_order"`
	DateCreated   time.Time   `json:"date_created"`
	DateExpires   time.Time   `json:"date_expires"`
	Status        QuoteStatus `json:"status"`
	Items         []QuoteItem `json:"items"`
}

// GetID returns the value of ID.
func (s *Quote) GetID() string {
	return s.ID
}

// GetPurchaseOrder returns the value of PurchaseOrder.
func (s *Quote) GetPurchaseOrder() string {
	return s.PurchaseOrder
}

// GetDateCreated returns the value of DateCreated.
func (s *Quote) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateExpires returns the value of DateExpires.
func (s *Quote) GetDateExpires() time.Time {
	return s.DateExpires
}

// GetStatus returns the value of Status.
func (s *Quote) GetStatus() QuoteStatus {
	return s.Status
}

// GetItems returns the value of Items.
func (s *Quote) GetItems() []QuoteItem {
	return s.Items
}

// SetID sets the value of ID.
func (s *Quote) SetID(val string) {
	s.ID = val
}

// SetPurchaseOrder sets the value of PurchaseOrder.
func (s *Quote) SetPurchaseOrder(val string) {
	s.PurchaseOrder = val
}

// SetDateCreated sets the value of DateCreated.
func (s *Quote) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateExpires sets the value of DateExpires.
func (s *Quote) SetDateExpires(val time.Time) {
	s.DateExpires = val
}

// SetStatus sets the value of Status.
func (s *Quote) SetStatus(val QuoteStatus) {
	s.Status = val
}

// SetItems sets the value of Items.
func (s *Quote) SetItems(val []QuoteItem) {
	s.Items = val
}

// Ref: #/components/schemas/QuoteItem
type QuoteItem struct {
	ProductID         string  `json:"product_id"`
	ProductSn         string  `json:"product_sn"`
	ProductName       string  `json:"product_name"`
	CustomerProductSn string  `json:"customer_product_sn"`
	OrderedQuantity   float64 `json:"ordered_quantity"`
	UnitType          string  `json:"unit_type"`
	UnitPrice         float64 `json:"unit_price"`
	TotalPrice        float64 `json:"total_price"`
}

// GetProductID returns the value of ProductID.
func (s *QuoteItem) GetProductID() string {
	return s.ProductID
}

// GetProductSn returns the value of ProductSn.
func (s *QuoteItem) GetProductSn() string {
	return s.ProductSn
}

// GetProductName returns the value of ProductName.
func (s *QuoteItem) GetProductName() string {
	return s.ProductName
}

// GetCustomerProductSn returns the value of CustomerProductSn.
func (s *QuoteItem) GetCustomerProductSn() string {
	return s.CustomerProductSn
}

// GetOrderedQuantity returns the value of OrderedQuantity.
func (s *QuoteItem) GetOrderedQuantity() float64 {
	return s.OrderedQuantity
}

// GetUnitType returns the value of UnitType.
func (s *QuoteItem) GetUnitType() string {
	return s.UnitType
}

// GetUnitPrice returns the value of UnitPrice.
func (s *QuoteItem) GetUnitPrice() float64 {
	return s.UnitPrice
}

// GetTotalPrice returns the value of TotalPrice.
func (s *QuoteItem) GetTotalPrice() float64 {
	return s.TotalPrice
}

// SetProductID sets the value of ProductID.
func (s *QuoteItem) SetProductID(val string) {
	s.ProductID = val
}

// SetProductSn sets the value of ProductSn.
func (s *QuoteItem) SetProductSn(val string) {
	s.ProductSn = val
}

// SetProductName sets the value of ProductName.
func (s *QuoteItem) SetProductName(val string) {
	s.ProductName = val
}

// SetCustomerProductSn sets the value of CustomerProductSn.
func (s *QuoteItem) SetCustomerProductSn(val string) {
	s.CustomerProductSn = val
}

// SetOrderedQuantity sets the value of OrderedQuantity.
func (s *QuoteItem) SetOrderedQuantity(val float64) {
	s.OrderedQuantity = val
}

// SetUnitType sets the value of UnitType.
func (s *QuoteItem) SetUnitType(val string) {
	s.UnitType = val
}

// SetUnitPrice sets the value of UnitPrice.
func (s *QuoteItem) SetUnitPrice(val float64) {
	s.UnitPrice = val
}

// SetTotalPrice sets the value of TotalPrice.
func (s *QuoteItem) SetTotalPrice(val float64) {
	s.TotalPrice = val
}

// Ref: #/components/schemas/QuoteStatus
type QuoteStatus string

const (
	QuoteStatusUnspecified     QuoteStatus = "unspecified"
	QuoteStatusPendingApproval QuoteStatus = "pending_approval"
	QuoteStatusApproved        QuoteStatus = "approved"
	QuoteStatusCancelled       QuoteStatus = "cancelled"
	QuoteStatusExpired         QuoteStatus = "expired"
)

// AllValues returns all QuoteStatus values.
func (QuoteStatus) AllValues() []QuoteStatus {
	return []QuoteStatus{
		QuoteStatusUnspecified,
		QuoteStatusPendingApproval,
		QuoteStatusApproved,
		QuoteStatusCancelled,
		QuoteStatusExpired,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s QuoteStatus) MarshalText() ([]byte, error) {
	switch s {
	case QuoteStatusUnspecified:
		return []byte(s), nil
	case QuoteStatusPendingApproval:
		return []byte(s), nil
	case QuoteStatusApproved:
		return []byte(s), nil
	case QuoteStatusCancelled:
		return []byte(s), nil
	case QuoteStatusExpired:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *QuoteStatus) UnmarshalText(data []byte) error {
	switch QuoteStatus(data) {
	case QuoteStatusUnspecified:
		*s = QuoteStatusUnspecified
		return nil
	case QuoteStatusPendingApproval:
		*s = QuoteStatusPendingApproval
		return nil
	case QuoteStatusApproved:
		*s = QuoteStatusApproved
		return nil
	case QuoteStatusCancelled:
		*s = QuoteStatusCancelled
		return nil
	case QuoteStatusExpired:
		*s = QuoteStatusExpired
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/QuoteSummary
type QuoteSummary struct {
	ID            string      `json:"id"`
	ContactID     string      `json:"contact_id"`
	BranchID      string      `json:"branch_id"`
	PurchaseOrder string      `json:"purchase_order"`
	Status        QuoteStatus `json:"status"`
	DateCreated   time.Time   `json:"date_created"`
	DateExpires   time.Time   `json:"date_expires"`
}

// GetID returns the value of ID.
func (s *QuoteSummary) GetID() string {
	return s.ID
}

// GetContactID returns the value of ContactID.
func (s *QuoteSummary) GetContactID() string {
	return s.ContactID
}

// GetBranchID returns the value of BranchID.
func (s *QuoteSummary) GetBranchID() string {
	return s.BranchID
}

// GetPurchaseOrder returns the value of PurchaseOrder.
func (s *QuoteSummary) GetPurchaseOrder() string {
	return s.PurchaseOrder
}

// GetStatus returns the value of Status.
func (s *QuoteSummary) GetStatus() QuoteStatus {
	return s.Status
}

// GetDateCreated returns the value of DateCreated.
func (s *QuoteSummary) GetDateCreated() time.Time {
	return s.DateCreated
}

// GetDateExpires returns the value of DateExpires.
func (s *QuoteSummary) GetDateExpires() time.Time {
	return s.DateExpires
}

// SetID sets the value of ID.
func (s *QuoteSummary) SetID(val string) {
	s.ID = val
}

// SetContactID sets the value of ContactID.
func (s *QuoteSummary) SetContactID(val string) {
	s.ContactID = val
}

// SetBranchID sets the value of BranchID.
func (s *QuoteSummary) SetBranchID(val string) {
	s.BranchID = val
}

// SetPurchaseOrder sets the value of PurchaseOrder.
func (s *QuoteSummary) SetPurchaseOrder(val string) {
	s.PurchaseOrder = val
}

// SetStatus sets the value of Status.
func (s *QuoteSummary) SetStatus(val QuoteStatus) {
	s.Status = val
}

// SetDateCreated sets the value of DateCreated.
func (s *QuoteSummary) SetDateCreated(val time.Time) {
	s.DateCreated = val
}

// SetDateExpires sets the value of DateExpires.
func (s *QuoteSummary) SetDateExpires(val time.Time) {
	s.DateExpires = val
}

type SearchProductsOK struct {
	Aggregations Aggregation  `json:"aggregations"`
	Metadata     PageMetadata `json:"metadata"`
	Products     []Product    `json:"products"`
}

// GetAggregations returns the value of Aggregations.
func (s *SearchProductsOK) GetAggregations() Aggregation {
	return s.Aggregations
}

// GetMetadata returns the value of Metadata.
func (s *SearchProductsOK) GetMetadata() PageMetadata {
	return s.Metadata
}

// GetProducts returns the value of Products.
func (s *SearchProductsOK) GetProducts() []Product {
	return s.Products
}

// SetAggregations sets the value of Aggregations.
func (s *SearchProductsOK) SetAggregations(val Aggregation) {
	s.Aggregations = val
}

// SetMetadata sets the value of Metadata.
func (s *SearchProductsOK) SetMetadata(val PageMetadata) {
	s.Metadata = val
}

// SetProducts sets the value of Products.
func (s *SearchProductsOK) SetProducts(val []Product) {
	s.Products = val
}

type SearchProductsReq struct {
	SortBy  OptString                   `json:"sort_by"`
	Filters OptSearchProductsReqFilters `json:"filters"`
	Page    OptInt                      `json:"page"`
	Query   OptString                   `json:"query"`
}

// GetSortBy returns the value of SortBy.
func (s *SearchProductsReq) GetSortBy() OptString {
	return s.SortBy
}

// GetFilters returns the value of Filters.
func (s *SearchProductsReq) GetFilters() OptSearchProductsReqFilters {
	return s.Filters
}

// GetPage returns the value of Page.
func (s *SearchProductsReq) GetPage() OptInt {
	return s.Page
}

// GetQuery returns the value of Query.
func (s *SearchProductsReq) GetQuery() OptString {
	return s.Query
}

// SetSortBy sets the value of SortBy.
func (s *SearchProductsReq) SetSortBy(val OptString) {
	s.SortBy = val
}

// SetFilters sets the value of Filters.
func (s *SearchProductsReq) SetFilters(val OptSearchProductsReqFilters) {
	s.Filters = val
}

// SetPage sets the value of Page.
func (s *SearchProductsReq) SetPage(val OptInt) {
	s.Page = val
}

// SetQuery sets the value of Query.
func (s *SearchProductsReq) SetQuery(val OptString) {
	s.Query = val
}

type SearchProductsReqFilters map[string][]string

func (s *SearchProductsReqFilters) init() SearchProductsReqFilters {
	m := *s
	if m == nil {
		m = map[string][]string{}
		*s = m
	}
	return m
}

type SetActiveBranchBadRequest Error

func (*SetActiveBranchBadRequest) setActiveBranchRes() {}

type SetActiveBranchForbidden Error

func (*SetActiveBranchForbidden) setActiveBranchRes() {}

type SetActiveBranchOK struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *SetActiveBranchOK) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *SetActiveBranchOK) SetMessage(val OptString) {
	s.Message = val
}

func (*SetActiveBranchOK) setActiveBranchRes() {}

type SetActiveBranchReq struct {
	// ID of the branch to set as active.
	BranchID string `json:"branch_id"`
}

// GetBranchID returns the value of BranchID.
func (s *SetActiveBranchReq) GetBranchID() string {
	return s.BranchID
}

// SetBranchID sets the value of BranchID.
func (s *SetActiveBranchReq) SetBranchID(val string) {
	s.BranchID = val
}
