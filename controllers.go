package covid

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func HandleGetCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	company := &Company{}
	found, err := company.Load(id)
	if err != nil {
		log.Errorf("Could not load company %s: %s", id, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not load company"}`))
		return
	}

	if !found {
		log.Warnf("Could not find company: %s\n", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	c, _ := json.Marshal(company)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(c)
}

func HandleGetCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := LoadAllCompanies()
	if err != nil {
		log.Errorf("Could not load all companies: %s", err.Error())
		http.Error(w, `{"error": "could not load companies"}`, http.StatusInternalServerError)
		return
	}
	c, _ := json.Marshal(companies)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(c)
}

// HandleCreateCompany is an internal endpoint that allows for creating a new company
// through the API. Should only be exposed on admin port.
func HandleCreateCompany(w http.ResponseWriter, r *http.Request) {
	company := &Company{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Could not read request: %s", err.Error())
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	log.Debugf("Read request body: %s", string(b))
	err = json.Unmarshal(b, &company)
	if err != nil {
		log.Errorf("Could not read json: %s", err.Error())
		http.Error(w, "malformed json body", http.StatusBadRequest)
		return
	}

	err = company.Save()
	if err != nil {
		log.Errorf("Could not save company: %s", err.Error())
		http.Error(w, "Could not save company", http.StatusInternalServerError)
		return
	}

	// Return created data
	c, _ := json.Marshal(company)
	w.WriteHeader(http.StatusCreated)
	w.Write(c)
}

// HandleSubmitArticle is the endpoint that allows for users to submit facts about
// companies. It will store in a sql table that can later be processed internally.
func HandleSubmitArticle(w http.ResponseWriter, r *http.Request) {
	fact := &ProposedFact{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Could not read request: %s", err.Error())
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	log.Debugf("Read request body: %s", string(b))
	err = json.Unmarshal(b, &fact)
	if err != nil {
		log.Errorf("Could not read json: %s", err.Error())
		http.Error(w, "malformed json body", http.StatusBadRequest)
		return
	}

	check := &ProposedFact{}
	exists, err := check.LoadCitation(fact.Citation)
	if err != nil {
		log.Errorf("Error loading citation: %s", err.Error())
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	log.Debugf("Fact exists: %v", exists)

	// Return 200 but secretly we don't need this article because we already have it
	if exists {
		log.Infof("Ignoring already submitted fact: %s", fact.Citation)
		submitted, _ := json.Marshal(check)
		w.Write(submitted)
		return
	}

	// If it doesn't exist, save the given one to the database
	err = fact.Save()
	if err != nil {
		log.Errorf("Error saving citation: %s", err.Error())
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Just fall through to a 200 here
	w.Header().Set("Access-Control-Allow-Origin", "*")
	submitted, err := json.Marshal(fact)
	if err == nil {
		w.Write(submitted)
		return
	}
	log.Errorf("Could not marshal saved fact: %s", err.Error())
}

// HandleGetProposedFacts is the internal handler for getting all unproccessed proposed
// facts. Should only be ran on the internal port
func HandleGetProposedFacts(w http.ResponseWriter, r *http.Request) {
	facts, err := GetUnprocessedProposedFacts()
	if err != nil {
		log.Errorf("Could not get proposed facts: %s", err.Error())
		http.Error(w, "Could not get proposed facts", http.StatusInternalServerError)
		return
	}
	f, err := json.Marshal(facts)
	if err != nil {
		log.Errorf("Could not marshal proposed facts: %s", err.Error())
		http.Error(w, "Could not marshal proposed facts", http.StatusInternalServerError)
		return
	}
	w.Write(f)
}

// HandleUpdateProposedFact is the internal handler for updating a proposed fact's
// state and should only be ran on the internal port
func HandleUpdateProposedFact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	pf := &ProposedFact{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Could not read request: %s", err.Error())
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	log.Debugf("Read request body: %s", string(b))
	err = json.Unmarshal(b, &pf)
	if err != nil {
		log.Errorf("Could not read json: %s", err.Error())
		http.Error(w, "malformed json body", http.StatusBadRequest)
		return
	}
	pf.Id = id

	err = pf.Save()
	if err != nil {
		log.Errorf("Could not save proposed fact: %s", err.Error())
		http.Error(w, "Could not save proposed fact", http.StatusInternalServerError)
		return
	}

	// Return created data
	c, _ := json.Marshal(pf)
	w.WriteHeader(http.StatusOK)
	w.Write(c)
}

// HandleCreateFact is the internal handler for creating a fact
func HandleCreateFact(w http.ResponseWriter, r *http.Request) {
	fact := &Fact{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Could not read request: %s", err.Error())
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	log.Debugf("Read request body: %s", string(b))
	err = json.Unmarshal(b, &fact)
	if err != nil {
		log.Errorf("Could not read json: %s", err.Error())
		http.Error(w, "malformed json body", http.StatusBadRequest)
		return
	}

	err = fact.Save()
	if err != nil {
		log.Errorf("Could not save company: %s", err.Error())
		http.Error(w, "Could not save company", http.StatusInternalServerError)
		return
	}

	// Return created data
	c, _ := json.Marshal(fact)
	w.WriteHeader(http.StatusCreated)
	w.Write(c)
}
