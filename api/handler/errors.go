package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//ErrInternalServer indicates that an unexpected error was thrown
	ErrInternalServer = NewRFC7807Failure("about:blank", "The server encountered an unexpected condition that prevented it from fulfilling the request", "", "", http.StatusInternalServerError)

	//ErrNotFound indicates that the resource which was asked was not found
	ErrNotFound = NewRFC7807Failure("about:blank", "Could not find the requested resource(s)", "", "", http.StatusNotFound)

	//ErrMultiFailure should be used when you need to return multi errors in a single struct
	ErrMultiFailure = NewRFC7807Failure("about:blank", "You have requested some invalid fields", "", "", http.StatusBadRequest)

	//ErrVisibilityFieldsMissing indicates that the client is ordering a data package which has not set properly
	ErrVisibilityFieldsMissing = NewRFC7807Failure("about:blank", "There's a Data Package settings missing",
		"You have not set the fields configuration yet. Please contact Newoway support", "", http.StatusBadRequest)

	//ErrBadRequest indicates that the client applied a wrong request
	ErrBadRequest = NewRFC7807Failure("about:blank", "bad request", "You have applied a wrong request it might be missing some field.", "", http.StatusBadRequest)
)

//BaseError is the simple error based on RFC-7807
type BaseError struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when
	// dereferenced
	Type string `json:"type"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except for
	// purposes of localization
	Title string `json:"title"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`
}

//NewBaseError creates a intance of BaseError
func NewBaseError(tp, title, detail string) BaseError {
	return BaseError{
		Type:   tp,
		Title:  title,
		Detail: detail,
	}
}

//NewRFC7807Failure creates a new instance of RFC7807Failure
func NewRFC7807Failure(tp, tittle, detail, instance string, status int) *RFC7807Failure {
	err := &RFC7807Failure{}
	err.Type = tp
	err.Title = tittle
	err.Status = status
	err.Instance = instance
	err.Detail = detail
	return err
}

//RFC7807Failure wraps the problem json that we send to client according to https://tools.ietf.org/html/rfc7807
type RFC7807Failure struct {
	//BaseError wraps the required fields shuch as type, title and detail
	BaseError

	// The HTTP status code for this occurrence of the problem
	Status int `json:"status,omitempty"`

	// A URI that identifies the specific occurrence of the problem. This URI
	// may or may not yield further information if dereferenced.
	Instance string `json:"instance,omitempty"`

	//Erros is the list of generic erros my happend
	Errors []BaseError `json:"errors,omitempty"`
}

//StatusCode is the http status in int format
func (f *RFC7807Failure) StatusCode() string {
	return http.StatusText(f.Status)
}

//Error is the type of error
func (f *RFC7807Failure) Error() string {
	return f.Type
}

//WithDetail add some detail to error
func (f *RFC7807Failure) WithDetail(detail string) *RFC7807Failure {
	f.Detail = detail
	return f
}

//WithErrors add a list of errors
func (f *RFC7807Failure) WithErrors(errors []BaseError) *RFC7807Failure {
	f.Errors = errors
	return f
}

//AddError add a single BaseError into the list of errors
func (f *RFC7807Failure) AddError(be BaseError) *RFC7807Failure {
	f.Errors = append(f.Errors, be)
	return f
}

// responseFailure respond with json error content RFC7807Failure
func responseFailure(c *gin.Context, tp, title, details, instance string, status int) {
	failure := NewRFC7807Failure(tp, title, details, instance, status)
	c.JSON(status, failure)
}
