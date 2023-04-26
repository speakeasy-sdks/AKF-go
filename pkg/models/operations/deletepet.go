// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"PetStore/pkg/models/shared"
	"net/http"
)

type DeletePetRequest struct {
	// ID of pet to delete
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePetResponse struct {
	ContentType string
	// unexpected error
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}