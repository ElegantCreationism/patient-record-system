package handlers

import "net/http"

// RootHandler returns an empty body status code
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func CreatePatient(w http.ResponseWriter, r *http.Request){

}

func GetPatient(w http.ResponseWriter, r *http.Request) {
	
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	
}

func DeletePatient(w http.ResponseWriter, r *http.Request){

}

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreatePatient(w, r)
	case http.MethodGet:
		GetPatient(w, r)
	case http.MethodPut:
		UpdatePatient(w, r)
	case http.MethodDelete:
		DeletePatient(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}

}