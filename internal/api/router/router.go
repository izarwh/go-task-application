package router

import (
	"encoding/json"
	"net/http"
	"task_planner_application/internal/task/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewChiRouter(taskHandler handlers.ITaskHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "up",
		})
	})

	// api contracts
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", taskHandler.CreateTask)
			r.Get("/", taskHandler.GetTasks)
			r.Get("/{task_id}", taskHandler.GetTask)
			r.Put("/{task_id}", taskHandler.UpdateTask)
			r.Delete("/{task_id}", taskHandler.DeleteTask)
		})
	})

	return r
}
