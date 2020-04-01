package sdk

import (
	"bytes"
	"covid"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

//var defaultURL = "http://34.68.214.180"
var defaultURL = "http://localhost:8080"
var defaultAdminURL = "http://localhost:8081"

type Client struct {
	URL      string
	AdminURL string
}

// CreateCompany will take a covid.Company struct and will make a call to the defined url
// to submit the fact. Once done, it will update the given struct to have the generated ID
// in it.
func (cbr Client) CreateCompany(company *covid.Company) error {
	b, err := json.Marshal(company)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(cbr.AdminURL+"/api/v1/company", "application/json", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Service returned non 201 code: " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, company)
	return err
}

// GetCompany will get a given company by ID
func (cbr Client) GetCompany(id string) (covid.Company, error) {
	var company covid.Company
	resp, err := http.Get(cbr.URL + "/api/v1/company/" + id)
	if err != nil {
		return company, err
	}
	if resp.StatusCode != http.StatusOK {
		return company, errors.New("Invalid response code: " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return company, err
	}
	err = json.Unmarshal(body, &company)
	return company, err
}

// SubmitArticle will submit a given fact
func (cbr Client) SubmitArticle(fact *covid.ProposedFact) error {
	b, err := json.Marshal(fact)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(cbr.URL+"/api/v1/fact/propose", "application/json", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid response code: " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, fact)
	return err
}

// GetAllFacts will query the service to get all _unprocessed_ proposed facts
func (cbr Client) GetAllFacts() ([]covid.ProposedFact, error) {
	var facts []covid.ProposedFact
	resp, err := http.Get(cbr.AdminURL + "/api/v1/fact/proposed")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid response code: " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &facts)
	return facts, err
}

// UpdateProposedFact will update a given proposed fact
func (cbr Client) UpdateProposedFact(fact covid.ProposedFact) error {
	b, err := json.Marshal(fact)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", cbr.AdminURL+"/api/v1/fact/proposed/"+fact.Id, r)
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid response code: " + string(resp.StatusCode))
	}
	return nil
}

// CreateFact will create a given fact, and update the struct with the fact's Id
func (cbr Client) CreateFact(fact *covid.Fact) error {
	b, err := json.Marshal(fact)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(cbr.AdminURL+"/api/v1/fact", "application/json", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Service returned non 201 code: " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, fact)
	return err
}
