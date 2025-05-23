package handlers

import "bisnode/internal/models"

// SearchPersonRequest represents the request body for searching a person
// swagger:parameters searchPerson
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type searchPersonRequest struct {
	// Mobile number of the person to search for
	// in: body
	MobileNumber string `json:"mobileNumber"`
}

// SearchOrganizationRequest represents the request body for searching an organization
// swagger:parameters searchOrganization
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type searchOrganizationRequest struct {
	// Organization number to search for
	// in: body
	OrganizationNumber string `json:"organizationNumber"`
}

// swagger:route GET /directory/persons/search directory searchPerson
//
// Search for a person by mobile number
//
// This endpoint allows you to search for a person using their mobile number.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: personSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError
//
// swagger:route POST /directory/persons/search directory searchPersonPost
//
// Search for a person by mobile number (POST)
//
// This endpoint allows you to search for a person using their mobile number.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: personSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError

// swagger:route GET /directory/organizations/search directory searchOrganization
//
// Search for an organization by organization number
//
// This endpoint allows you to search for an organization using its organization number.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: organizationSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError
//
// swagger:route POST /directory/organizations/search directory searchOrganizationPost
//
// Search for an organization by organization number (POST)
//
// This endpoint allows you to search for an organization using its organization number.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: organizationSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError

// swagger:response personSearchResponse
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type personSearchResponseWrapper struct {
	// in:body
	Body models.DirectorySearchResponse
}

// swagger:response organizationSearchResponse
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type organizationSearchResponseWrapper struct {
	// in:body
	Body models.DirectorySearchResponse
}
