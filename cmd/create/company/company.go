package company

import (
	"bufio"
	"covid"
	"covid/sdk"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	Url         string
	CompanyName string
	Logo        string
	Rating      string
}

// Allow for the RunE function to be called externally, without making the flags struct
// public
func ExternalRunE(Url string) error {
	flags := &flags{
		Url: Url,
	}
	return RunE(flags)
}

func RunE(f *flags) error {
	reader := bufio.NewReader(os.Stdin)
	if f.CompanyName == "" {
		fmt.Printf("Enter company name: ")
		f.CompanyName, _ = reader.ReadString('\n')
		f.CompanyName = strings.TrimSuffix(f.CompanyName, "\n")
	}

	if f.Logo == "" {
		fmt.Printf("Enter company logo: ")
		f.Logo, _ = reader.ReadString('\n')
		f.Logo = strings.TrimSuffix(f.Logo, "\n")
	}

	if f.Rating == "" {
		fmt.Printf("Enter company rating: ")
		f.Rating, _ = reader.ReadString('\n')
		f.Rating = strings.TrimSuffix(f.Rating, "\n")
	}

	rating, err := strconv.ParseFloat(f.Rating, 8)
	if err != nil {
		return err
	}

	cbr := sdk.Client{
		AdminURL: f.Url,
	}

	company := &covid.Company{
		Name:   f.CompanyName,
		Logo:   f.Logo,
		Rating: rating,
	}

	err = cbr.CreateCompany(company)
	if err != nil {
		return err
	}
	fmt.Printf("Server returned:\n%+v\n", company)
	return nil
}

func NewCommand() *cobra.Command {
	flags := &flags{}
	var cmd = &cobra.Command{
		Use:   "company",
		Short: "Create a new company",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunE(flags)
		},
	}
	cmd.Flags().StringVarP(&flags.Url, "url", "u", "http://localhost:8081", "Service URL")
	cmd.Flags().StringVar(&flags.CompanyName, "name", "", "Name of company")
	cmd.Flags().StringVar(&flags.Logo, "logo", "", "Logo of company")
	cmd.Flags().StringVar(&flags.Rating, "rating", "", "Rating")
	return cmd
}
