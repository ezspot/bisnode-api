package handlers

import "bisnode/internal/models"

// SearchRequest represents the request body for searching a motor vehicle
// swagger:parameters searchMotorVehicle
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type searchMotorVehicleRequest struct {
	// License number of the vehicle (optional if VIN is provided)
	// in: body
	LicenseNumber string `json:"licenseNumber,omitempty"`
	// VIN (Vehicle Identification Number) of the vehicle (optional if license number is provided)
	// in: body
	VIN string `json:"vin,omitempty"`
}

// swagger:route GET /motor-vehicles/search motor-vehicles searchMotorVehicle
//
// Search for a motor vehicle by license number or VIN
//
// This endpoint allows you to search for motor vehicle information using either the license number or VIN.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: motorVehicleSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError
//
// swagger:route POST /motor-vehicles/search motor-vehicles searchMotorVehiclePost
//
// Search for a motor vehicle by license number or VIN (POST)
//
// This endpoint allows you to search for motor vehicle information using either the license number or VIN.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//     - basicAuth: []
//
//     Responses:
//       200: motorVehicleSearchResponse
//       400: badRequestResponse
//       401: unauthorizedResponse
//       500: internalServerError

// swagger:response motorVehicleSearchResponse
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type motorVehicleSearchResponseWrapper struct {
	// in:body
	Body models.MotorVehicleSearchResponse
}

// swagger:response badRequestResponse
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type badRequestResponseWrapper struct {
	// in:body
	Body struct {
		// The error message
		Error string `json:"error"`
	}
}

// swagger:response unauthorizedResponse
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type unauthorizedResponseWrapper struct {
	// in:body
	Body struct {
		// The error message
		Error string `json:"error"`
	}
}

// swagger:response internalServerError
//
//lint:ignore U1000 Ignore unused code, it's used for Swagger documentation
type internalServerErrorWrapper struct {
	// in:body
	Body struct {
		// The error message
		Error string `json:"error"`
	}
}
