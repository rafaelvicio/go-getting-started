package pipelines

import (
	"encoding/json"
	"net/http"
	"pipenator-backend/common"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Pipeline struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	URL         string        `json:"url" bson:"url"`
	Email       string        `json:"email" bson:"email"`
	DataCriacao time.Time     `json:"data-criacao" bson:"data-criacao"`
}

const collectionName = "pipelines"

func SaveOne(w http.ResponseWriter, r *http.Request) {
	db := common.GetDBFromContext(r)

	var p Pipeline
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = bson.NewObjectId()
	p.DataCriacao = time.Now()

	if err := db.C(collectionName).Insert(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	db := common.GetDBFromContext(r)

	var pipelines []*Pipeline
	if err := db.C(collectionName).Find(nil).Sort("-data-criacao").All(&pipelines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pipelines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ShowOne(w http.ResponseWriter, r *http.Request) {
	db := common.GetDBFromContext(r)

	id := mux.Vars(r)["id"]

	var p *Pipeline

	if err := db.C(collectionName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	db := common.GetDBFromContext(r)

	id := mux.Vars(r)["id"]

	oid := bson.ObjectIdHex(id)

	if err := db.C(collectionName).RemoveId(oid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	db := common.GetDBFromContext(r)

	id := mux.Vars(r)["id"]

	var p Pipeline
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = bson.ObjectIdHex(id)

	if err := db.C(collectionName).UpdateId(&p.ID, &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
