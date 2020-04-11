package process

import (
	"bufio"
	"covid"
	"covid/cmd/create/company"
	"covid/sdk"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type flags struct {
	AdminUrl string
	Url      string
}

func RunE(f *flags) error {

	client := sdk.Client{
		AdminURL: f.AdminUrl,
		URL:      f.Url,
	}

	// Get all the unprocessed articles
	pfacts, err := client.GetAllFacts()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	for _, pfact := range pfacts {
		fmt.Printf("%s --- %s [%s %s]\n",
			pfact.Citation,
			pfact.Summary,
			pfact.CompanyId,
			pfact.CompanyName)
		fmt.Printf("Accept? [Y/n]\n")
		r, _ := reader.ReadString('\n')
		r = strings.TrimSuffix(r, "\n")

		if strings.ToUpper(r) != "Y" {
			if strings.ToUpper(r) == "E" {
				fmt.Printf("Exiting...\n")
				break
			}
			fmt.Printf("Rejecting...\n")
			pfact.Rejected = true
			err := client.UpdateProposedFact(pfact)
			if err != nil {
				fmt.Printf("Could not reject proposed fact: %s\n", pfact.Id)
			}
			continue
		}

		// Attempt to lookup company id
		createCompany, linkCompany := true, false
		if pfact.CompanyId != "" {
			_, err := client.GetCompany(pfact.CompanyId)
			if err == nil {
				createCompany = false
			}

			fmt.Printf("Couldn't find existing company by ID, link existing? [id/n] ")
			r, _ := reader.ReadString('\n')
			r = strings.TrimSuffix(r, "\n")

			if strings.ToUpper(r) == "N" {
				createCompany = false
				linkCompany = true
			}

		}

		// Create company if we have too
		if createCompany {
			err := company.ExternalRunE(f.Url)
			if err != nil {
				return err
			}
		}

		// Link Company
		if linkCompany {
			fmt.Printf("Enter ID: ")
			r, _ := reader.ReadString('\n')
			pfact.CompanyId = strings.TrimSuffix(r, "\n")
			_, err := client.GetCompany(pfact.CompanyId)
			if err == nil {
				fmt.Printf("Invalid company ID -- skipping...")
				continue
			}
		}

		fact := &covid.Fact{
			Summary:   pfact.Summary,
			Citation:  pfact.Citation,
			CompanyId: pfact.CompanyId,
		}
		err := client.CreateFact(fact)
		if err != nil {
			fmt.Printf("Could not create fact for: %s\n", pfact.Id)
		}

		pfact.Approved = true
		err = client.UpdateProposedFact(pfact)
		if err != nil {
			fmt.Printf("Could not accept fact: %s\n", pfact.Id)
		}
	}

	return nil
}

func NewCommand() *cobra.Command {
	flags := &flags{}
	cmd := &cobra.Command{
		Use:   "process",
		Short: "Process proposed facts and create real facts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunE(flags)
		},
	}
	cmd.Flags().StringVarP(&flags.Url, "url", "u", "http://localhost:8080", "URL of service")
	cmd.Flags().StringVarP(&flags.AdminUrl, "admin-url", "a", "http://localhost:8081", "Admin URL of service")
	return cmd
}
