package pipelines

import (
	"net/http"
)

func CreatePipeline(w http.ResponseWriter, r *http.Request) {
	SaveOne(w, r)
}

func ListPipelines(w http.ResponseWriter, r *http.Request) {
	ListAll(w, r)
}

func ShowPipeline(w http.ResponseWriter, r *http.Request) {
	ShowOne(w, r)
}

func DeletePipeline(w http.ResponseWriter, r *http.Request) {
	DeleteOne(w, r)
}

func UpdatePipeline(w http.ResponseWriter, r *http.Request) {
	UpdateOne(w, r)
}
