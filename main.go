package main

import(
	"net/http"
	"encoding/json"
	"log/slog"
)

const addr = "127.0.0.1:8080"

type Vinil struct{
	ID 		string  `json: "id"`
	Title 	string  `json: "title"`
	Artist 	string  `json: "artist"`
	Price	float64 `json: "price"`
}

var vinils = []Vinil{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func responseJSON(w http.ResponseWriter, code int, payload any){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /vinils", getVinils)
	mux.HandleFunc("GET /vinils/{id}", getVinilById)
	mux.HandleFunc("POST /vinils", createVinil)

	slog.Info("server start listening", "at", addr)
	http.ListenAndServe(addr, mux)
}

func getVinils(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	responseJSON(w, http.StatusOK, vinils)
}

func getVinilById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	for _, v := range vinils{
		if v.ID == id{
			responseJSON(w, http.StatusOK, v)
			return
		}
	}

	responseJSON(w, http.StatusNotFound, map[string]string{"message": "vinil doesn't exist"})
}

func createVinil(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var newVinil Vinil

	if err := json.NewDecoder(r.Body).Decode(&newVinil); err != nil{
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	vinils = append(vinils, newVinil)

	responseJSON(w, http.StatusCreated, newVinil)
}