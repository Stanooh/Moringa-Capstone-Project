package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Item represents a simple resource
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// In-memory storage (simulating a database)
var items = []Item{
	{ID: 1, Name: "Laptop", Price: 700},
	{ID: 2, Name: "Phone", Price: 400},
}

var nextID = 3

// =====================
// Utility Function
// =====================
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// =====================
// GET All Items
// =====================
func getItems(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, items)
}

// =====================
// GET Single Item by ID
// =====================
func getItemByID(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
		return
	}

	for _, item := range items {
		if item.ID == id {
			respondWithJSON(w, http.StatusOK, item)
			return
		}
	}

	respondWithJSON(w, http.StatusNotFound, map[string]string{
		"error": "Item not found",
	})
}

// =====================
// POST New Item
// =====================
func createItem(w http.ResponseWriter, r *http.Request) {

	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	// Basic validation
	if newItem.Name == "" || newItem.Price <= 0 {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Name and Price are required",
		})
		return
	}

	newItem.ID = nextID
	nextID++

	items = append(items, newItem)

	respondWithJSON(w, http.StatusCreated, newItem)
}

// =====================
// Router
// =====================
func router(w http.ResponseWriter, r *http.Request) {

	switch {

	case r.Method == http.MethodGet && r.URL.Path == "/items":
		getItems(w, r)

	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/items/"):
		getItemByID(w, r)

	case r.Method == http.MethodPost && r.URL.Path == "/items":
		createItem(w, r)

	default:
		respondWithJSON(w, http.StatusNotFound, map[string]string{
			"error": "Route not found",
		})
	}
}

func main() {

	fmt.Println("🚀 Server running at http://localhost:8080")

	http.HandleFunc("/", router)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
