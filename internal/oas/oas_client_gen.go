// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// CreateQuote invokes createQuote operation.
	//
	// Create a new quote. The `customer_id` and `contact_id` are extracted from the provided
	// authentication token. Make sure to include a valid bearer token in the `Authorization` header.
	//
	// POST /account/quotes
	CreateQuote(ctx context.Context, request *CreateQuoteReq) (CreateQuoteRes, error)
	// GetOrder invokes getOrder operation.
	//
	// Get an order by ID.
	//
	// GET /orders/{id}
	GetOrder(ctx context.Context, params GetOrderParams) (GetOrderRes, error)
	// GetProduct invokes getProduct operation.
	//
	// Get a product by ID.
	//
	// GET /products/{id}
	GetProduct(ctx context.Context, params GetProductParams) (GetProductRes, error)
	// ListCustomerBranches invokes listCustomerBranches operation.
	//
	// Get available branches for customer.
	//
	// GET /account/branches
	ListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) (ListCustomerBranchesRes, error)
	// ListOrderInvoices invokes listOrderInvoices operation.
	//
	// Get invoices for an order.
	//
	// GET /orders/{id}/invoices
	ListOrderInvoices(ctx context.Context, params ListOrderInvoicesParams) (ListOrderInvoicesRes, error)
	// ListOrderShipments invokes listOrderShipments operation.
	//
	// Get shipments for an order.
	//
	// GET /orders/{id}/shipments
	ListOrderShipments(ctx context.Context, params ListOrderShipmentsParams) (ListOrderShipmentsRes, error)
	// ListOrders invokes listOrders operation.
	//
	// Get a list of orders.
	//
	// GET /account/orders
	ListOrders(ctx context.Context, params ListOrdersParams) ([]OrderSummary, error)
	// SearchProducts invokes searchProducts operation.
	//
	// Search for products.
	//
	// POST /search/products
	SearchProducts(ctx context.Context, request *SearchProductsReq) (SearchProductsRes, error)
	// SetActiveBranch invokes setActiveBranch operation.
	//
	// Set active branch for current user.
	//
	// PUT /account/branch
	SetActiveBranch(ctx context.Context, request *SetActiveBranchReq) (SetActiveBranchRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// CreateQuote invokes createQuote operation.
//
// Create a new quote. The `customer_id` and `contact_id` are extracted from the provided
// authentication token. Make sure to include a valid bearer token in the `Authorization` header.
//
// POST /account/quotes
func (c *Client) CreateQuote(ctx context.Context, request *CreateQuoteReq) (CreateQuoteRes, error) {
	res, err := c.sendCreateQuote(ctx, request)
	return res, err
}

func (c *Client) sendCreateQuote(ctx context.Context, request *CreateQuoteReq) (res CreateQuoteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createQuote"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/account/quotes"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateQuoteOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/account/quotes"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateQuoteRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateQuoteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetOrder invokes getOrder operation.
//
// Get an order by ID.
//
// GET /orders/{id}
func (c *Client) GetOrder(ctx context.Context, params GetOrderParams) (GetOrderRes, error) {
	res, err := c.sendGetOrder(ctx, params)
	return res, err
}

func (c *Client) sendGetOrder(ctx context.Context, params GetOrderParams) (res GetOrderRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getOrder"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/orders/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetOrderOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/orders/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetOrderResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetProduct invokes getProduct operation.
//
// Get a product by ID.
//
// GET /products/{id}
func (c *Client) GetProduct(ctx context.Context, params GetProductParams) (GetProductRes, error) {
	res, err := c.sendGetProduct(ctx, params)
	return res, err
}

func (c *Client) sendGetProduct(ctx context.Context, params GetProductParams) (res GetProductRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getProduct"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/products/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetProductOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/products/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetProductResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListCustomerBranches invokes listCustomerBranches operation.
//
// Get available branches for customer.
//
// GET /account/branches
func (c *Client) ListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) (ListCustomerBranchesRes, error) {
	res, err := c.sendListCustomerBranches(ctx, params)
	return res, err
}

func (c *Client) sendListCustomerBranches(ctx context.Context, params ListCustomerBranchesParams) (res ListCustomerBranchesRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listCustomerBranches"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/account/branches"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListCustomerBranchesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/account/branches"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "contact_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "contact_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.ContactID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, ListCustomerBranchesOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListCustomerBranchesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListOrderInvoices invokes listOrderInvoices operation.
//
// Get invoices for an order.
//
// GET /orders/{id}/invoices
func (c *Client) ListOrderInvoices(ctx context.Context, params ListOrderInvoicesParams) (ListOrderInvoicesRes, error) {
	res, err := c.sendListOrderInvoices(ctx, params)
	return res, err
}

func (c *Client) sendListOrderInvoices(ctx context.Context, params ListOrderInvoicesParams) (res ListOrderInvoicesRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listOrderInvoices"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/orders/{id}/invoices"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListOrderInvoicesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/orders/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/invoices"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, ListOrderInvoicesOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListOrderInvoicesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListOrderShipments invokes listOrderShipments operation.
//
// Get shipments for an order.
//
// GET /orders/{id}/shipments
func (c *Client) ListOrderShipments(ctx context.Context, params ListOrderShipmentsParams) (ListOrderShipmentsRes, error) {
	res, err := c.sendListOrderShipments(ctx, params)
	return res, err
}

func (c *Client) sendListOrderShipments(ctx context.Context, params ListOrderShipmentsParams) (res ListOrderShipmentsRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listOrderShipments"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/orders/{id}/shipments"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListOrderShipmentsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/orders/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/shipments"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, ListOrderShipmentsOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListOrderShipmentsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListOrders invokes listOrders operation.
//
// Get a list of orders.
//
// GET /account/orders
func (c *Client) ListOrders(ctx context.Context, params ListOrdersParams) ([]OrderSummary, error) {
	res, err := c.sendListOrders(ctx, params)
	return res, err
}

func (c *Client) sendListOrders(ctx context.Context, params ListOrdersParams) (res []OrderSummary, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listOrders"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/account/orders"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListOrdersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/account/orders"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Page))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "per_page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.PerPage))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListOrdersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SearchProducts invokes searchProducts operation.
//
// Search for products.
//
// POST /search/products
func (c *Client) SearchProducts(ctx context.Context, request *SearchProductsReq) (SearchProductsRes, error) {
	res, err := c.sendSearchProducts(ctx, request)
	return res, err
}

func (c *Client) sendSearchProducts(ctx context.Context, request *SearchProductsReq) (res SearchProductsRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchProducts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/search/products"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SearchProductsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/search/products"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSearchProductsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSearchProductsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SetActiveBranch invokes setActiveBranch operation.
//
// Set active branch for current user.
//
// PUT /account/branch
func (c *Client) SetActiveBranch(ctx context.Context, request *SetActiveBranchReq) (SetActiveBranchRes, error) {
	res, err := c.sendSetActiveBranch(ctx, request)
	return res, err
}

func (c *Client) sendSetActiveBranch(ctx context.Context, request *SetActiveBranchReq) (res SetActiveBranchRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setActiveBranch"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/account/branch"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SetActiveBranchOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/account/branch"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSetActiveBranchRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, SetActiveBranchOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSetActiveBranchResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
