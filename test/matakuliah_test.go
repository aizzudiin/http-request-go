package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nama_npm_pert4/model" // Adjust with your folder name
)

func HandlerMatkulGet(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error
	Kd_mk := r.URL.Query()["Kd_mk"]
	if len(Kd_mk) != 0 {
		data, err = model.GetMatkul(db, Kd_mk[0])
	} else {
		data, err = model.GetAllMatkul(db)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func HandlerMatkulPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Matkul
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func HandlerMatkulDelete(w http.ResponseWriter, r *http.Request) {
	kd_mk := r.URL.Query()["kd_mk"]
	if len(kd_mk) != 0 {
		data := model.Matkul{Kd_mk: kd_mk[0]}
		if err := data.Delete(db); err != nil {
			http.Error(w, "ID not found", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Data has been deleted"))
	} else {
		http.Error(w, "ID not found", http.StatusInternalServerError)
	}
}

func HandlerMatkulPut(w http.ResponseWriter, r *http.Request) {
	kd_mk := r.URL.Query()["kd_mk"]
	if len(kd_mk) == 0 {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonmap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := model.Matkul{Kd_mk: kd_mk[0]}
	err = data.Update(db, jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := model.GetMatkul(db, kd_mk[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
