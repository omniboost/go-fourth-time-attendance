package fourth_time_attendance

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-fourth-time-attendance/utils"
)

func (c *Client) NewPostClockRequest() PostClockRequest {
	return PostClockRequest{
		client:      c,
		queryParams: c.NewPostClockQueryParams(),
		pathParams:  c.NewPostClockPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostClockRequestBody(),
	}
}

type PostClockRequest struct {
	client      *Client
	queryParams *PostClockQueryParams
	pathParams  *PostClockPathParams
	method      string
	headers     http.Header
	requestBody PostClockRequestBody
}

func (c *Client) NewPostClockQueryParams() *PostClockQueryParams {
	return &PostClockQueryParams{}
}

type PostClockQueryParams struct{}

func (p PostClockQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostClockRequest) QueryParams() *PostClockQueryParams {
	return r.queryParams
}

func (c *Client) NewPostClockPathParams() *PostClockPathParams {
	return &PostClockPathParams{}
}

type PostClockPathParams struct {
	ClientID string `schema:"client_id"`
}

func (p *PostClockPathParams) Params() map[string]string {
	return map[string]string{
		"client_id": p.ClientID,
	}
}

func (r *PostClockRequest) PathParams() *PostClockPathParams {
	return r.pathParams
}

func (r *PostClockRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostClockRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostClockRequest) Method() string {
	return r.method
}

func (s *Client) NewPostClockRequestBody() PostClockRequestBody {
	return PostClockRequestBody{}
}

type PostClockRequestBody ClockingPostRequest

func (rb PostClockRequestBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(rb.Root)
}

func (r *PostClockRequest) RequestBody() *PostClockRequestBody {
	return &r.requestBody
}

func (r *PostClockRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostClockRequest) SetRequestBody(body PostClockRequestBody) {
	r.requestBody = body
}

func (r *PostClockRequest) NewResponseBody() *PostClockResponseBody {
	return &PostClockResponseBody{}
}

type PostClockResponseBody ClockingPostResponse

func (r *PostClockRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.client_id}}/tnasubmission", r.PathParams())
	return &u
}

func (r *PostClockRequest) Do() (PostClockResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
