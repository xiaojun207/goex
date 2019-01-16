// Code generated by go-swagger; DO NOT EDIT.

package leaderboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// LeaderboardGetReader is a Reader for the LeaderboardGet structure.
type LeaderboardGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeaderboardGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLeaderboardGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLeaderboardGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewLeaderboardGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLeaderboardGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeaderboardGetOK creates a LeaderboardGetOK with default headers values
func NewLeaderboardGetOK() *LeaderboardGetOK {
	return &LeaderboardGetOK{}
}

/*LeaderboardGetOK handles this case with default header values.

Request was successful
*/
type LeaderboardGetOK struct {
	Payload []*models.Leaderboard
}

func (o *LeaderboardGetOK) Error() string {
	return fmt.Sprintf("[GET /leaderboard][%d] leaderboardGetOK  %+v", 200, o.Payload)
}

func (o *LeaderboardGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaderboardGetBadRequest creates a LeaderboardGetBadRequest with default headers values
func NewLeaderboardGetBadRequest() *LeaderboardGetBadRequest {
	return &LeaderboardGetBadRequest{}
}

/*LeaderboardGetBadRequest handles this case with default header values.

Parameter Error
*/
type LeaderboardGetBadRequest struct {
	Payload *models.Error
}

func (o *LeaderboardGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /leaderboard][%d] leaderboardGetBadRequest  %+v", 400, *o.Payload)
}

func (o *LeaderboardGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaderboardGetUnauthorized creates a LeaderboardGetUnauthorized with default headers values
func NewLeaderboardGetUnauthorized() *LeaderboardGetUnauthorized {
	return &LeaderboardGetUnauthorized{}
}

/*LeaderboardGetUnauthorized handles this case with default header values.

Unauthorized
*/
type LeaderboardGetUnauthorized struct {
	Payload *models.Error
}

func (o *LeaderboardGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /leaderboard][%d] leaderboardGetUnauthorized  %+v", 401, o.Payload)
}

func (o *LeaderboardGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLeaderboardGetNotFound creates a LeaderboardGetNotFound with default headers values
func NewLeaderboardGetNotFound() *LeaderboardGetNotFound {
	return &LeaderboardGetNotFound{}
}

/*LeaderboardGetNotFound handles this case with default header values.

Not Found
*/
type LeaderboardGetNotFound struct {
	Payload *models.Error
}

func (o *LeaderboardGetNotFound) Error() string {
	return fmt.Sprintf("[GET /leaderboard][%d] leaderboardGetNotFound  %+v", 404, o.Payload)
}

func (o *LeaderboardGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}