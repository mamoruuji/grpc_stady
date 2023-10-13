package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/lib/pq"
	"github.com/mamoruuji/grpc_stady_api/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	todov1 "github.com/mamoruuji/grpc_stady_api/gen/proto/todo/v1"        // generated by protoc-gen-go
	"github.com/mamoruuji/grpc_stady_api/gen/proto/todo/v1/todov1connect" // generated by protoc-gen-connect-go
)

type todoServer struct {
	todov1connect.TodoServiceHandler
}

func (s *todoServer) ListTodos(
	ctx context.Context,
	_ *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {
	todos, err := models.Todos().All(ctx, boil.GetContextDB())
	if err != nil {
		log.Printf("failed to get todos: %v", err)
		return nil, err
	}

	var pbTodos []*todov1.TodoData
	for _, d := range todos {
		todoID := int32(d.ID)
		createdAT := timestamppb.New(d.CreatedAt)
		updatedAT := timestamppb.New(d.UpdatedAt)
		pbTodo := &todov1.TodoData{
			Id:        todoID,
			Title:     d.Title,
			Context:   d.Context,
			CreatedAt: createdAT,
			UpdatedAt: updatedAT,
		}
		pbTodos = append(pbTodos, pbTodo)
	}

	res := connect.NewResponse(&todov1.ListTodosResponse{
		Todos: pbTodos,
	})

	return res, nil
}

func (s *todoServer) AddTodo(
	ctx context.Context,
	req *connect.Request[todov1.AddTodoRequest],
) (*connect.Response[todov1.AddTodoResponse], error) {
	d := models.Todo{
		Title:   req.Msg.Title,
		Context: req.Msg.Context,
	}
	err := d.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Printf("failed to add todo: %v", err)
		return nil, err
	}

	todoID := int32(d.ID)
	res := connect.NewResponse(&todov1.AddTodoResponse{
		Id: todoID,
	})

	return res, nil
}

func (s *todoServer) DeleteTodo(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTodoRequest],
) (*connect.Response[todov1.DeleteTodoResponse], error) {
	todoID := int(req.Msg.Id)
	d := models.Todo{
		ID: todoID,
	}

	_, err := d.Delete(ctx, boil.GetContextDB())
	if err != nil {
		log.Printf("failed to delete todo: %v", err)
		return nil, err
	}

	return connect.NewResponse(&todov1.DeleteTodoResponse{}), nil
}

func (s *todoServer) UpdateTodoStatus(
	ctx context.Context,
	req *connect.Request[todov1.UpdateTodoStatusRequest],
) (*connect.Response[todov1.UpdateTodoStatusResponse], error) {
	todoID := int(req.Msg.Id)
	d := models.Todo{
		ID: todoID,
	}

	d.Title = req.Msg.Title
	d.Context = req.Msg.Context
	_, err := d.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Printf("failed to update todo status: %v", err)
		return nil, err
	}

	return connect.NewResponse(&todov1.UpdateTodoStatusResponse{}), nil
}

func server() http.Handler {
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTodoServiceHandler(&todoServer{})
	mux.Handle(path, handler)
	return mux
}

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=pass dbname=grpc_stady sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

	mux := server()

	err = http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		log.Fatalf("failed to listen(tcp, :8080)")
	}

}
