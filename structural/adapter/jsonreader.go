package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//constants JSON_FILE_PATH
const (
	JSONFilePath = "bank_details.json"
)

//ReadBankDetailsFromJSON function
func ReadBankDetailsFromJSON() BankDetails {
	bankDetailJSONFile, err := os.Open(JSONFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer bankDetailJSONFile.Close()
	ByteValue, _ := ioutil.ReadAll(bankDetailJSONFile)
	var bankDetails BankDetails
	json.Unmarshal(ByteValue, &bankDetails)

	return bankDetails
}
