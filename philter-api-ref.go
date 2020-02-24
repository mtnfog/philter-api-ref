/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Status struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type Explain struct {
	FilteredText  string `json:"filteredText"`
	Context string `json:"context"`
	DocumentId string `json:"documentI"`
}

func main() {

	fmt.Println("Starting service...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/filter", filter).Methods("POST")
	router.HandleFunc("/api/explain", explain).Methods("POST")
	router.HandleFunc("/api/status", status).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func filter(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("x-document-id", "asdfghjkl12345678")

	fmt.Fprintln(w, "{{{REDACTED-entity}}} was a patient.")

}

func explain(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	context := params["context"]
	documentId := params["documentId"]

	explain := Explain{FilteredText: "{{{REDACTED-entity}}} was a patient.", Context: context, DocumentId: documentId}

	if err := json.NewEncoder(w).Encode(explain); err != nil {
		panic(err)
	}

}

func status(w http.ResponseWriter, r *http.Request) {

	status := Status{Status: "Healthy", Version: "1.0.0"}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}

}
