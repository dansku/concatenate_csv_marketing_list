package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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

func main() {

	oldFile := os.Args[1]
	newFile := os.Args[2]

	csvFile, _ := os.Open(oldFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var company []Company
	var nameList []string // Temporary list with only names

	i := 0
	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		nameList = append(nameList, line[0])

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

		i++
	}

	// Parse second file, and add just in case it is not alrady there
	csvFile, _ = os.Open(newFile)
	reader = csv.NewReader(bufio.NewReader(csvFile))
	j := 0

	var newList []string

	added := 0
	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// Make a list of only new items
		newList = append(newList, line[0])

		// If company doesn't exist in old list, add it.
		if Contains(nameList, line[0]) == false {
			company = append(company, Company{
				Name: line[0],
			})
			nameList = append(nameList, line[0])
			fmt.Printf("[%s] was added to the list!\n", line[0])
			added++
		}

		j++
	}

	// Check if old is present in new, otherwise remove it
	csvFile, _ = os.Open(oldFile)
	reader = csv.NewReader(bufio.NewReader(csvFile))
	removed := 0
	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// Check if old is in new list
		if !Contains(newList, line[0]) {
			fmt.Printf("[%s] was removed from the list!\n", line[0])
			removed++
		}
	}

	fmt.Printf("Total %d - %d added and %d removed.\n", len(nameList), added, removed)

	// companyJSON, _ := json.Marshal(company)
	// ioutil.WriteFile("output.json", companyJSON, 0644)
	// fmt.Println(string(companyJSON))

	// diff := j - i
	// fmt.Printf("There are %d companies in old and %d in new. Diff %d\n", i, j, diff)
}
