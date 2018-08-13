// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package search is a auto generated package.
Input file: svc_search.proto
*/
package search

import (
	"context"

	"github.com/pensando/sw/api"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

// ServiceSearchV1Client  is the client interface for the service.
type ServiceSearchV1Client interface {
	AutoWatchSvcSearchV1(ctx context.Context, in *api.ListWatchOptions) (SearchV1_AutoWatchSvcSearchV1Client, error)

	PolicyQuery(ctx context.Context, t *PolicySearchRequest) (*PolicySearchResponse, error)
	Query(ctx context.Context, t *SearchRequest) (*SearchResponse, error)
}

// ServiceSearchV1Server is the server interface for the service.
type ServiceSearchV1Server interface {
	AutoWatchSvcSearchV1(in *api.ListWatchOptions, stream SearchV1_AutoWatchSvcSearchV1Server) error

	PolicyQuery(ctx context.Context, t PolicySearchRequest) (PolicySearchResponse, error)
	Query(ctx context.Context, t SearchRequest) (SearchResponse, error)
}
