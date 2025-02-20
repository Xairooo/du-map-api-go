// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Du Map
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.10
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// MapAPIService is a service that implements the logic for the MapAPIServicer
// This service should implement the business logic for every endpoint for the MapAPI API.
// Include any external packages or services that will be required by this service.
type MapAPIService struct {
}

// NewMapAPIService creates a default api service
func NewMapAPIService() *MapAPIService {
	return &MapAPIService{}
}

// GetFaces - 
func (s *MapAPIService) GetFaces(ctx context.Context, tileId int64, celestialId int64, scale int64) (ImplResponse, error) {
	// TODO - update GetFaces with the required logic for this service method.
	// Add api_map_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Face{}) or use other options such as http.Ok ...
	// return Response(200, []Face{}), nil

	// TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	// return Response(401, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFaces method not implemented")
}

// UpdateTile - 
func (s *MapAPIService) UpdateTile(ctx context.Context, planet string, tileId int64) (ImplResponse, error) {
	// TODO - update UpdateTile with the required logic for this service method.
	// Add api_map_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Face{}) or use other options such as http.Ok ...
	// return Response(200, Face{}), nil

	// TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	// return Response(401, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateTile method not implemented")
}
