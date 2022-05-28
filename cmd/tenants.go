package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
)

// TenantsCmd represents the Tenants command
var TenantsCmd = &cobra.Command{
	Use:   "tenants",
	Short: "List of all tenants",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTenants()
	},
}

func init() {
	getCmd.AddCommand(TenantsCmd)
}

type Tenants struct {
	ID       int      `header:"id"`
	Name     string   `header:"name"`
	Networks []string `header:"networks"`
}

func getTenants() {
	url := "https://api.nmsDomain.com/magma/v1/tenants"
	responseBytes := getTenantsData(url)

	var tenants []Tenants

	if err := json.Unmarshal(responseBytes, &tenants); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	tableprinter.Print(os.Stdout, tenants)

}

func checkError(err error, hdr string) {
	if err != nil {
		fmt.Printf("[%s] Fatal error: %v\n", hdr, err.Error())
		os.Exit(1)
	}
}

func getTenantsData(baseAPI string) []byte {

	cert, err := tls.LoadX509KeyPair("/path/to/admin_operator.pem", "/path/to/admin_operator.key.pem")
	checkError(err, "loadcert")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
			ClientAuth:         tls.RequireAnyClientCert,
		},
	}
	client := &http.Client{Transport: tr}

	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a magma api. %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
