package api

import (
	"covid"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"net/http"
)

type flags struct {
	Config string
}

type CmdApi struct {
	Config covid.Config
}

func (a CmdApi) RunE() error {
	// connect to database
	covid.ConnectDatabase(a.Config.Database)

	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/company/{id}", covid.HandleGetCompany).Methods("GET")
	r.HandleFunc("/api/v1/company/{id}", covid.HandleSubmitArticle).Methods("POST")
	r.HandleFunc("/api/v1/company", covid.HandleGetCompanies).Methods("GET")
	r.HandleFunc("/api/v1/fact", covid.HandleSubmitArticle).Methods("POST")

	return http.ListenAndServe(a.Config.Bind, r)
}

func NewCommand() *cobra.Command {
	flags := &flags{}
	var cmd = &cobra.Command{
		Use:   "api",
		Short: "Run the Covid-19 Business Tracker API",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := covid.LoadConfigFile(flags.Config)
			if err != nil {
				return err
			}
			err = config.Log.SetupLogger()
			if err != nil {
				return err
			}

			api := CmdApi{config}
			return api.RunE()
		},
	}
	cmd.Flags().StringVarP(&flags.Config, "config", "c", "covid.yaml", "Config file")
	return cmd
}
