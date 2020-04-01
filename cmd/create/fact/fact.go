package fact

import (
	"bufio"
	"covid"
	"covid/sdk"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type flags struct {
	AdminUrl    string
	Url         string
	Summary     string
	Citation    string
	CompanyId   string
	Propose     bool
	CompanyName string
}

func RunE(f *flags) error {
	reader := bufio.NewReader(os.Stdin)
	if f.Citation == "" {
		fmt.Printf("Enter fact citation: ")
		f.Citation, _ = reader.ReadString('\n')
		f.Citation = strings.TrimSuffix(f.Citation, "\n")
	}

	if f.Summary == "" {
		fmt.Printf("Enter fact summary: ")
		f.Summary, _ = reader.ReadString('\n')
		f.Summary = strings.TrimSuffix(f.Summary, "\n")
	}

	if f.CompanyId == "" && !f.Propose {
		fmt.Printf("Enter companyid: ")
		f.CompanyId, _ = reader.ReadString('\n')
		f.CompanyId = strings.TrimSuffix(f.CompanyId, "\n")
	}

	if f.CompanyName == "" && f.Propose {
		fmt.Printf("Enter company name: ")
		f.CompanyName, _ = reader.ReadString('\n')
		f.CompanyName = strings.TrimSuffix(f.CompanyName, "\n")
	}

	cbr := sdk.Client{
		URL:      f.Url,
		AdminURL: f.AdminUrl,
	}

	if !f.Propose {
		fact := &covid.Fact{
			Summary:   f.Summary,
			Citation:  f.Citation,
			CompanyId: f.CompanyId,
		}

		err := cbr.CreateFact(fact)
		if err != nil {
			return err
		}
		fmt.Printf("Server returned:\n%+v\n", fact)
		return nil
	}

	proposed := &covid.ProposedFact{
		Summary:     f.Summary,
		Citation:    f.Citation,
		CompanyName: f.CompanyName,
	}
	err := cbr.SubmitArticle(proposed)
	if err != nil {
		return err
	}
	fmt.Printf("Server returned:\n%+v\n", proposed)
	return nil
}

func NewCommand() *cobra.Command {
	flags := &flags{}
	var cmd = &cobra.Command{
		Use:   "fact",
		Short: "Create a new fact directly - skipping processing",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunE(flags)
		},
	}
	cmd.Flags().StringVarP(&flags.Url, "url", "u", "http://localhost:8080", "Service URL")
	cmd.Flags().StringVar(&flags.AdminUrl, "admin-url", "http://localhost:8081", "Admin Service URL")
	cmd.Flags().StringVar(&flags.Summary, "summary", "", "Short summary of fact")
	cmd.Flags().StringVar(&flags.Citation, "citation", "", "Citation for fact (ie news article)")
	cmd.Flags().StringVar(&flags.CompanyId, "company-id", "", "CompanyId to link fact too")
	cmd.Flags().BoolVarP(&flags.Propose, "propose", "p", false, "Propose fact instead of creating directly")
	cmd.Flags().StringVar(&flags.CompanyName, "name", "", "Company name when proposing fact")
	return cmd
}
