// Define APIs here

package main
import (
	"github.com/gorilla/mux"
	"learn-golang/pkg/handler"
	"log"
	"net/http"
)

func main() {
	// Tạo một router mới từ gorilla/mux
	r := mux.NewRouter()
	log.Println("Router created")

	// Đăng ký handler cho endpoint /api/todo
	r.HandleFunc("/api/todo", handler.GetAllTodo).Methods(http.MethodGet)
	log.Println("Route /api/todo registered")

	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	log.Println("Route /api/todo/{id} registered")

	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	log.Println("Route /api/todo registered, created new todo")

	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)

	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)

	// Bắt đầu lắng nghe và phục vụ trên cổng 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
