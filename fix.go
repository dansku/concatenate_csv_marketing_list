package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Company build the desired structure of company info
type Company struct {
	Name      string `json:"name"`
	Website   string `json:"website"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Linkedin  string `json:"linkedin"`
	Youtube   string `json:"youtube"`
	Github    string `json:"github"`
	Change    string `json:"change"`
}

// Contains check if string is present in slice.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Reads the CSV File and return a list of companies
func getNames(fileName string) []string {
	var oldList []string
	csvFileOld, _ := os.Open(fileName)
	readerOld := csv.NewReader(bufio.NewReader(csvFileOld))
	for {
		line, error := readerOld.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		oldList = append(oldList, line[0])
	}
	return oldList
}

func main() {

	oldFile := os.Args[1]
	newFile := os.Args[2]

	oldCompanyList := getNames(oldFile)
	newCompanyList := getNames(newFile)

	fmt.Printf("Old company #%d.\n", len(oldCompanyList))
	fmt.Printf("New company #%d.\n", len(newCompanyList))

	// finalList := generateFinalList(newCompanyList, oldCompanyList)

	// Generate complete list
	var completeList []string
	csvFile, _ := os.Open(oldFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var company []Company

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// Only append if old name is in the new list also
		if Contains(newCompanyList, line[0]) {
			completeList = append(completeList, line[0])
			company = append(company, Company{
				Name:      line[0],
				Website:   line[1],
				Twitter:   line[2],
				Instagram: line[3],
				Facebook:  line[4],
				Linkedin:  line[5],
				Youtube:   line[6],
				Github:    line[7],
			})
		}

	}
	// Add remaining new companies in Company list
	for _, c := range newCompanyList {
		if !Contains(completeList, c) {
			// Add remaining new list to company
			company = append(company, Company{
				Name: c,
			})
		}
	}

	companyJSON, _ := json.Marshal(company)
	ioutil.WriteFile("output.json", companyJSON, 0644)
	fmt.Println(string(companyJSON))

}
