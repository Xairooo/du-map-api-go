// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Du Map
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.10
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/Xairooo/du-map-api-go"
)

func main() {
	log.Printf("Server started")

	MapAPIService := openapi.NewMapAPIService()
	MapAPIController := openapi.NewMapAPIController(MapAPIService)

	ScanAPIService := openapi.NewScanAPIService()
	ScanAPIController := openapi.NewScanAPIController(ScanAPIService)

	SharesAPIService := openapi.NewSharesAPIService()
	SharesAPIController := openapi.NewSharesAPIController(SharesAPIService)

	router := openapi.NewRouter(MapAPIController, ScanAPIController, SharesAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
