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
	company1 := Company{
		Id:     "1",
		Name:   "Walmart",
		Logo:   "https://shelbyreport.nyc3.cdn.digitaloceanspaces.com/wp-content/uploads/2017/02/New_Walmart_Logo.svg-e1519222902338.jpg",
		Rating: 4.5,
		Facts: []Fact{
			{"4", "Laid off 10 employees", "https://google.com", "", false},
			{"5", "Pinky promised not to lay off employees", "https://google.com", "", false},
			{"6", "Forced workers to keep coming in", "https://google.com", "", false},
		},
	}
	company2 := Company{
		Id:     "2",
		Name:   "Amazon",
		Logo:   "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.businessinsider.com%2Fhidden-meanings-in-tech-company-logos-2014-6&psig=AOvVaw1YDGp2XBbcfdTNEzxaKIoO&ust=1585503786193000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCNiQ04HcvegCFQAAAAAdAAAAABAR",
		Rating: 5.5,
		Facts: []Fact{
			{"7", "Hired 1000 new workers", "https://google.com", "", false},
			{"8", "No longer forcing workers to pee in bottles", "https://google.com", "", false},
			{"9", "Started to pay overtime", "https://google.com", "", false},
		},
	}
	companies := []Company{company1, company2}
	c, _ := json.Marshal(companies)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(c)
}

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
		http.Error(w, "", http.StatusOK)
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
}
