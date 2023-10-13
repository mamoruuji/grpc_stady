// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/todo/v1/todo.proto

package todov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/mamoruuji/grpc_stady_api/gen/proto/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "proto.todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceListTodosProcedure is the fully-qualified name of the TodoService's ListTodos RPC.
	TodoServiceListTodosProcedure = "/proto.todo.v1.TodoService/ListTodos"
	// TodoServiceAddTodoProcedure is the fully-qualified name of the TodoService's AddTodo RPC.
	TodoServiceAddTodoProcedure = "/proto.todo.v1.TodoService/AddTodo"
	// TodoServiceDeleteTodoProcedure is the fully-qualified name of the TodoService's DeleteTodo RPC.
	TodoServiceDeleteTodoProcedure = "/proto.todo.v1.TodoService/DeleteTodo"
	// TodoServiceUpdateTodoStatusProcedure is the fully-qualified name of the TodoService's
	// UpdateTodoStatus RPC.
	TodoServiceUpdateTodoStatusProcedure = "/proto.todo.v1.TodoService/UpdateTodoStatus"
)

// TodoServiceClient is a client for the proto.todo.v1.TodoService service.
type TodoServiceClient interface {
	ListTodos(context.Context, *connect_go.Request[v1.ListTodosRequest]) (*connect_go.Response[v1.ListTodosResponse], error)
	AddTodo(context.Context, *connect_go.Request[v1.AddTodoRequest]) (*connect_go.Response[v1.AddTodoResponse], error)
	DeleteTodo(context.Context, *connect_go.Request[v1.DeleteTodoRequest]) (*connect_go.Response[v1.DeleteTodoResponse], error)
	UpdateTodoStatus(context.Context, *connect_go.Request[v1.UpdateTodoStatusRequest]) (*connect_go.Response[v1.UpdateTodoStatusResponse], error)
}

// NewTodoServiceClient constructs a client for the proto.todo.v1.TodoService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		listTodos: connect_go.NewClient[v1.ListTodosRequest, v1.ListTodosResponse](
			httpClient,
			baseURL+TodoServiceListTodosProcedure,
			opts...,
		),
		addTodo: connect_go.NewClient[v1.AddTodoRequest, v1.AddTodoResponse](
			httpClient,
			baseURL+TodoServiceAddTodoProcedure,
			opts...,
		),
		deleteTodo: connect_go.NewClient[v1.DeleteTodoRequest, v1.DeleteTodoResponse](
			httpClient,
			baseURL+TodoServiceDeleteTodoProcedure,
			opts...,
		),
		updateTodoStatus: connect_go.NewClient[v1.UpdateTodoStatusRequest, v1.UpdateTodoStatusResponse](
			httpClient,
			baseURL+TodoServiceUpdateTodoStatusProcedure,
			opts...,
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	listTodos        *connect_go.Client[v1.ListTodosRequest, v1.ListTodosResponse]
	addTodo          *connect_go.Client[v1.AddTodoRequest, v1.AddTodoResponse]
	deleteTodo       *connect_go.Client[v1.DeleteTodoRequest, v1.DeleteTodoResponse]
	updateTodoStatus *connect_go.Client[v1.UpdateTodoStatusRequest, v1.UpdateTodoStatusResponse]
}

// ListTodos calls proto.todo.v1.TodoService.ListTodos.
func (c *todoServiceClient) ListTodos(ctx context.Context, req *connect_go.Request[v1.ListTodosRequest]) (*connect_go.Response[v1.ListTodosResponse], error) {
	return c.listTodos.CallUnary(ctx, req)
}

// AddTodo calls proto.todo.v1.TodoService.AddTodo.
func (c *todoServiceClient) AddTodo(ctx context.Context, req *connect_go.Request[v1.AddTodoRequest]) (*connect_go.Response[v1.AddTodoResponse], error) {
	return c.addTodo.CallUnary(ctx, req)
}

// DeleteTodo calls proto.todo.v1.TodoService.DeleteTodo.
func (c *todoServiceClient) DeleteTodo(ctx context.Context, req *connect_go.Request[v1.DeleteTodoRequest]) (*connect_go.Response[v1.DeleteTodoResponse], error) {
	return c.deleteTodo.CallUnary(ctx, req)
}

// UpdateTodoStatus calls proto.todo.v1.TodoService.UpdateTodoStatus.
func (c *todoServiceClient) UpdateTodoStatus(ctx context.Context, req *connect_go.Request[v1.UpdateTodoStatusRequest]) (*connect_go.Response[v1.UpdateTodoStatusResponse], error) {
	return c.updateTodoStatus.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the proto.todo.v1.TodoService service.
type TodoServiceHandler interface {
	ListTodos(context.Context, *connect_go.Request[v1.ListTodosRequest]) (*connect_go.Response[v1.ListTodosResponse], error)
	AddTodo(context.Context, *connect_go.Request[v1.AddTodoRequest]) (*connect_go.Response[v1.AddTodoResponse], error)
	DeleteTodo(context.Context, *connect_go.Request[v1.DeleteTodoRequest]) (*connect_go.Response[v1.DeleteTodoResponse], error)
	UpdateTodoStatus(context.Context, *connect_go.Request[v1.UpdateTodoStatusRequest]) (*connect_go.Response[v1.UpdateTodoStatusResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	todoServiceListTodosHandler := connect_go.NewUnaryHandler(
		TodoServiceListTodosProcedure,
		svc.ListTodos,
		opts...,
	)
	todoServiceAddTodoHandler := connect_go.NewUnaryHandler(
		TodoServiceAddTodoProcedure,
		svc.AddTodo,
		opts...,
	)
	todoServiceDeleteTodoHandler := connect_go.NewUnaryHandler(
		TodoServiceDeleteTodoProcedure,
		svc.DeleteTodo,
		opts...,
	)
	todoServiceUpdateTodoStatusHandler := connect_go.NewUnaryHandler(
		TodoServiceUpdateTodoStatusProcedure,
		svc.UpdateTodoStatus,
		opts...,
	)
	return "/proto.todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceListTodosProcedure:
			todoServiceListTodosHandler.ServeHTTP(w, r)
		case TodoServiceAddTodoProcedure:
			todoServiceAddTodoHandler.ServeHTTP(w, r)
		case TodoServiceDeleteTodoProcedure:
			todoServiceDeleteTodoHandler.ServeHTTP(w, r)
		case TodoServiceUpdateTodoStatusProcedure:
			todoServiceUpdateTodoStatusHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) ListTodos(context.Context, *connect_go.Request[v1.ListTodosRequest]) (*connect_go.Response[v1.ListTodosResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.ListTodos is not implemented"))
}

func (UnimplementedTodoServiceHandler) AddTodo(context.Context, *connect_go.Request[v1.AddTodoRequest]) (*connect_go.Response[v1.AddTodoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.AddTodo is not implemented"))
}

func (UnimplementedTodoServiceHandler) DeleteTodo(context.Context, *connect_go.Request[v1.DeleteTodoRequest]) (*connect_go.Response[v1.DeleteTodoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.DeleteTodo is not implemented"))
}

func (UnimplementedTodoServiceHandler) UpdateTodoStatus(context.Context, *connect_go.Request[v1.UpdateTodoStatusRequest]) (*connect_go.Response[v1.UpdateTodoStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.UpdateTodoStatus is not implemented"))
}