// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     string        `json:"image_type"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Version string `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws   UploadTypes = "aws"
	UploadTypes_azure UploadTypes = "azure"
	UploadTypes_gcp   UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposes(ctx, params)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"7fZ/AQAA//+9kAz1PywAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
