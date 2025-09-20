package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
)

type Server struct {
	domain *application.Service
	mux    *http.ServeMux
}

func NewServer(domainInstance *application.Service) *Server {
	s := &Server{
		domain: domainInstance,
		mux:    http.NewServeMux(),
	}
	s.setupRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/accounts", s.handleAccounts)
	s.mux.HandleFunc("/accounts/", s.handleAccountsWithName)
	s.mux.HandleFunc("/clear", s.handleClear)
}

func (s *Server) handleAccounts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.createAccount(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleAccountsWithName(w http.ResponseWriter, r *http.Request) {
	// Extract account name from path
	path := strings.TrimPrefix(r.URL.Path, "/accounts/")
	parts := strings.Split(path, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Account name required", http.StatusBadRequest)
		return
	}

	accountName := parts[0]

	// Handle different endpoints
	if len(parts) == 1 {
		// /accounts/{name}
		switch r.Method {
		case "GET":
			s.getAccount(w, r, accountName)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if len(parts) == 2 {
		action := parts[1]
		switch action {
		case "activate":
			if r.Method == "POST" {
				s.activateAccount(w, r, accountName)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		case "authenticate":
			if r.Method == "POST" {
				s.authenticateAccount(w, r, accountName)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		case "authentication-status":
			if r.Method == "GET" {
				s.getAuthenticationStatus(w, r, accountName)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		case "projects":
			switch r.Method {
			case "GET":
				s.getProjects(w, r, accountName)
			case "POST":
				s.createProject(w, r, accountName)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (s *Server) handleClear(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		s.clearAll(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) createAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		s.writeError(w, "Name is required", http.StatusBadRequest)
		return
	}

	if err := s.domain.CreateAccount(req.Name); err != nil {
		s.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getAccount(w http.ResponseWriter, _ *http.Request, name string) {
	account, err := s.domain.GetAccount(name)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			s.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			s.writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := struct {
		Name          string `json:"name"`
		Activated     bool   `json:"activated"`
		Authenticated bool   `json:"authenticated"`
	}{
		Name:          account.Name(),
		Activated:     account.IsActivated(),
		Authenticated: account.IsAuthenticated(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (s *Server) activateAccount(w http.ResponseWriter, r *http.Request, name string) {
	if err := s.domain.Activate(name); err != nil {
		if strings.Contains(err.Error(), "not found") {
			s.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			s.writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) authenticateAccount(w http.ResponseWriter, r *http.Request, name string) {
	if err := s.domain.Authenticate(name); err != nil {
		if strings.Contains(err.Error(), "not found") {
			s.writeError(w, err.Error(), http.StatusNotFound)
		} else if strings.Contains(err.Error(), "activate") {
			s.writeError(w, err.Error(), http.StatusBadRequest)
		} else {
			s.writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := struct {
		Authenticated bool `json:"authenticated"`
	}{
		Authenticated: true,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (s *Server) getAuthenticationStatus(w http.ResponseWriter, r *http.Request, name string) {
	authenticated := s.domain.IsAuthenticated(name)

	response := struct {
		Authenticated bool `json:"authenticated"`
	}{
		Authenticated: authenticated,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (s *Server) getProjects(w http.ResponseWriter, r *http.Request, name string) {
	projects, err := s.domain.GetProjects(name)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			s.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			s.writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request, name string) {
	if err := s.domain.CreateProject(name); err != nil {
		if strings.Contains(err.Error(), "not found") {
			s.writeError(w, err.Error(), http.StatusNotFound)
		} else {
			s.writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) clearAll(w http.ResponseWriter, r *http.Request) {
	s.domain.ClearAll()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	_ = json.NewEncoder(w).Encode(errorResponse)
}
