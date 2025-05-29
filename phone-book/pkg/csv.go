package pkg

import (
	csvEnc "encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func openAndCheckFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	return os.Open(filePath)

}

func ReadDataFromCsFile(filePath string, comma rune) ([][]string, error) {

	f, err := openAndCheckFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csvEnc.NewReader(f)
	csvReader.Comma = comma

	return csvReader.ReadAll()
}

func WriteDataToCsvFile(filePath string, comma rune, data [][]string) error {

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWr := csvEnc.NewWriter(f)
	csvWr.Comma = comma

	return csvWr.WriteAll(data)

}

func WriteDataToJsonFile(filepath string, data any) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(f, string(b))
	if err != nil {
		return err
	}

	return nil

}

func ReadDataFromJsonFile(filePath string, dataSlice any) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(dataSlice)
	if err != nil {
		panic(err)
	}

}
