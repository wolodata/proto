// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: convert/v1/convert.proto

package convertv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/wolodata/proto/gen/convert/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ConvertServiceName is the fully-qualified name of the ConvertService service.
	ConvertServiceName = "convert.v1.ConvertService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ConvertServicePdfToHtmlProcedure is the fully-qualified name of the ConvertService's PdfToHtml
	// RPC.
	ConvertServicePdfToHtmlProcedure = "/convert.v1.ConvertService/PdfToHtml"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	convertServiceServiceDescriptor         = v1.File_convert_v1_convert_proto.Services().ByName("ConvertService")
	convertServicePdfToHtmlMethodDescriptor = convertServiceServiceDescriptor.Methods().ByName("PdfToHtml")
)

// ConvertServiceClient is a client for the convert.v1.ConvertService service.
type ConvertServiceClient interface {
	PdfToHtml(context.Context, *connect.Request[v1.PdfToHtmlRequest]) (*connect.Response[v1.PdfToHtmlResponse], error)
}

// NewConvertServiceClient constructs a client for the convert.v1.ConvertService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewConvertServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ConvertServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &convertServiceClient{
		pdfToHtml: connect.NewClient[v1.PdfToHtmlRequest, v1.PdfToHtmlResponse](
			httpClient,
			baseURL+ConvertServicePdfToHtmlProcedure,
			connect.WithSchema(convertServicePdfToHtmlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// convertServiceClient implements ConvertServiceClient.
type convertServiceClient struct {
	pdfToHtml *connect.Client[v1.PdfToHtmlRequest, v1.PdfToHtmlResponse]
}

// PdfToHtml calls convert.v1.ConvertService.PdfToHtml.
func (c *convertServiceClient) PdfToHtml(ctx context.Context, req *connect.Request[v1.PdfToHtmlRequest]) (*connect.Response[v1.PdfToHtmlResponse], error) {
	return c.pdfToHtml.CallUnary(ctx, req)
}

// ConvertServiceHandler is an implementation of the convert.v1.ConvertService service.
type ConvertServiceHandler interface {
	PdfToHtml(context.Context, *connect.Request[v1.PdfToHtmlRequest]) (*connect.Response[v1.PdfToHtmlResponse], error)
}

// NewConvertServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewConvertServiceHandler(svc ConvertServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	convertServicePdfToHtmlHandler := connect.NewUnaryHandler(
		ConvertServicePdfToHtmlProcedure,
		svc.PdfToHtml,
		connect.WithSchema(convertServicePdfToHtmlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/convert.v1.ConvertService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ConvertServicePdfToHtmlProcedure:
			convertServicePdfToHtmlHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedConvertServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedConvertServiceHandler struct{}

func (UnimplementedConvertServiceHandler) PdfToHtml(context.Context, *connect.Request[v1.PdfToHtmlRequest]) (*connect.Response[v1.PdfToHtmlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("convert.v1.ConvertService.PdfToHtml is not implemented"))
}