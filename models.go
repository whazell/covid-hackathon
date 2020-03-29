package covid

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type Company struct {
	Id      string
	Name    string
	Logo    string
	Rating  float64
	Facts   []Fact
	Deleted bool
}

type Fact struct {
	Id        string
	Summary   string
	Citation  string
	CompanyId string
	Deleted   bool
}

type ProposedFact struct {
	Id          string
	Summary     string
	Citation    string
	CompanyId   string
	DateAdded   time.Time `json:"-"`
	CompanyName string

	// Non-approved proposed facts imply they have yet to be processed. When Rejected is
	// true, that means they have been processed but not accepted and should be ignored.
	Approved bool `json:"-"`
	Rejected bool `json:"-"`
}

// Save the company to the database, if a new company a id is generated and assigned to
// the struct
func (c *Company) Save() error {
	query := `INSERT INTO companies (id, name, logo, rating, deleted)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		logo = VALUES(logo),
		rating = VALUES(rating),
		deleted = VALUES(deleted)`

	if c.Id == "" {
		c.Id = genId()
	}
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Id, c.Name, c.Logo, c.Rating, c.Deleted)
	if err != nil {
		return err
	}

	// Update associations
	for _, fact := range c.Facts {
		if err := fact.Save(); err != nil {
			return err
		}
	}
	log.Debugf("Executing: %s", query)
	log.Debugf("Data: %+v", c)
	return nil
}

// Load the company into the given struct by id
func (c *Company) Load(id string) (bool, error) {
	rows, err := db.Query(`SELECT * FROM companies WHERE id = ?`, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// There should only be one result in the query
	if !rows.Next() {
		return false, nil
	}

	err = rows.Scan(&c.Id, &c.Name, &c.Logo, &c.Rating, &c.Deleted)
	if err != nil {
		return false, err
	}
	c.Facts, err = LoadFacts(id)
	log.Debugf("Loaded company from database: %+v", c)
	return true, err
}

func LoadAllCompanies() ([]Company, error) {
	log.Debug("Running: SELECT * FROM companies")
	rows, err := db.Query(`SELECT * FROM companies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []Company
	for rows.Next() {
		c := Company{}
		err := rows.Scan(&c.Id, &c.Name, &c.Logo, &c.Rating, &c.Deleted)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}
	return companies, nil
}

func (f *Fact) Save() error {
	query := `INSERT INTO facts (id, summary, citation, company_id, deleted)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		summary = VALUES(summary),
		citation = VALUES(citation),
		company_id = VALUES(company_id),
		deleted = VALUES(deleted)`

	if f.Id == "" {
		f.Id = genId()
	}
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	log.Debugf("Executing: %s", query)
	log.Debugf("Data: %+v", f)
	_, err = stmt.Exec(f.Id, f.Summary, f.Citation, f.CompanyId, f.Deleted)
	return err
}

func (f *Fact) Load(id string) (bool, error) {
	rows, err := db.Query(`SELECT * FROM facts WHERE id = ?`, id)
	log.Debugf("Executing: SELECT * FROM facts WHERE id = %s", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// There should only be one result in the query
	if !rows.Next() {
		return false, nil
	}

	err = rows.Scan(&f.Id, &f.Summary, &f.Citation, &f.CompanyId, &f.Deleted)
	if err != nil {
		return false, err
	}
	return true, nil
}

func LoadFacts(companyId string) ([]Fact, error) {
	rows, err := db.Query(`SELECT * FROM facts WHERE company_id = ?`, companyId)
	log.Debugf("Executing: SELECT * FROM facts WHERE company_id = %s", companyId)
	if err != nil {
		return []Fact{}, err
	}
	defer rows.Close()

	var facts []Fact
	for rows.Next() {
		f := Fact{}
		err := rows.Scan(&f.Id, &f.Summary, &f.Citation, &f.CompanyId, &f.Deleted)
		if err != nil {
			return []Fact{}, err
		}
		facts = append(facts, f)
	}
	return facts, nil
}

func (p *ProposedFact) Load(id string) (bool, error) {
	rows, err := db.Query(`SELECT * FROM proposed_facts WHERE id = ?`, id)
	log.Debugf("Executing: SELECT * FROM proposed_facts WHERE id = %s", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return p.scan(rows)
}

func (p *ProposedFact) LoadCitation(citation string) (bool, error) {
	rows, err := db.Query(`SELECT * FROM proposed_facts WHERE citation = ?`, citation)
	log.Debugf("Executing: SELECT * FROM proposed_facts WHERE citation = %s", citation)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return p.scan(rows)
}

func (p *ProposedFact) scan(rows *sql.Rows) (bool, error) {
	// There should only be one result in the query
	if !rows.Next() {
		return false, nil
	}

	err := rows.Scan(&p.Id, &p.Summary, &p.Citation, &p.CompanyId, &p.CompanyName,
		&p.DateAdded, &p.Approved, &p.Rejected)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUnprocessedProposedFacts() ([]ProposedFact, error) {
	query := `SELECT * FROM proposed_facts where approved = 0 AND rejected = 0`
	log.Debugf("Running: %s", query)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var facts []ProposedFact
	for {
		pf := &ProposedFact{}
		found, err := pf.scan(rows)
		if err != nil {
			return nil, err
		}
		if !found {
			break
		}
		facts = append(facts, *pf)
	}
	return facts, nil
}

func (p *ProposedFact) Save() error {
	query := `INSERT INTO proposed_facts 
		(id, summary, citation, company_id, company_name, date_added, approved, rejected)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		summary = VALUES(summary),
		approved = VALUES(approved),
		rejected = VALUES(rejected)
	`

	if p.Id == "" {
		p.Id = genId()
		p.DateAdded = time.Now()
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	log.Debugf("Executing: %s", query)
	log.Debugf("Data: %+v", p)
	_, err = stmt.Exec(p.Id, p.Summary, p.Citation, p.CompanyId, p.CompanyName, p.DateAdded, p.Approved, p.Rejected)
	return err
}

func genId() string {
	return uuid.NewV4().String()
}
