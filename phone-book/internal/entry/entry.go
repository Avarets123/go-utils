package entry

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"phone-book/internal/index"
	"phone-book/pkg"
)

type Entry struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

type PhoneBooks struct {
	Entries []Entry
	index   index.Index
}

func (p *PhoneBooks) SaveDataToFile() error {

	// var saveData [][]string

	// for _, v := range p.Entries {
	// 	entrySlice := []string{
	// 		v.Name, v.Surname, v.Phone,
	// 	}
	// 	saveData = append(saveData, entrySlice)

	// }

	return pkg.WriteDataToJsonFile("data.json", p.Entries)

	// return pkg.WriteDataToCsvFile("./data.csv", '|', saveData)

}

func (p *PhoneBooks) IdxList() map[string]int {
	return p.index.Idx
}

func InitPhoneBook() PhoneBooks {

	// data, err := pkg.ReadDataFromCsFile("./data.csv", '|')
	// if err != nil {
	// 	panic(err)
	// }

	entries := []Entry{}

	// e2 := []Entry{}
	pkg.ReadDataFromJsonFile("data.json", &entries)

	// fmt.Printf("%+v", e2)

	// for _, v := range data {
	// 	entries = append(entries, Entry{
	// 		Name:    v[0],
	// 		Surname: v[1],
	// 		Phone:   v[2],
	// 	})
	// }

	pb := PhoneBooks{
		Entries: entries,
	}

	pb.index = *index.New(entries, "Phone")

	return pb

}

func (pb *PhoneBooks) List() []Entry {
	return pb.Entries
}

func (pb *PhoneBooks) Search(phone string) *Entry {
	idx, has := pb.index.Exists(phone)
	if !has {
		return nil
	}

	return &pb.Entries[idx]

}

func (pb *PhoneBooks) Insert(name, surname, phone string) error {
	_, has := pb.index.Exists(phone)
	if has {
		return fmt.Errorf("Phone exists!")
	}

	pb.Entries = append(pb.Entries, Entry{
		Name:    name,
		Surname: name,
		Phone:   phone,
	})

	pb.index.Add(len(pb.Entries)-1, phone)

	return pb.SaveDataToFile()

}

func (pb *PhoneBooks) Delete(phone string) error {
	idx, has := pb.index.Exists(phone)
	if !has {
		return fmt.Errorf("Phone not found!")

	}

	pb.Entries = append(pb.Entries[:idx], pb.Entries[idx+1:]...)

	pb.index.Delete(phone)

	return pb.SaveDataToFile()

}

func GenerateRandomEntries(dataCount int) []Entry {
	resEntries := make([]Entry, dataCount)

	for i := 0; i < dataCount; i++ {
		newEntry := Entry{
			Name:    getRandomString(6),
			Surname: getRandomString(6),
			Phone:   fmt.Sprintf("+7928%d", getRandomInt(1000000, 9999999)),
		}

		fmt.Printf("%+v \n", newEntry)

		resEntries = append(resEntries, newEntry)

		i++

	}

	return resEntries

}

func getRandomString(strLen int) string {
	b := make([]byte, strLen)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)

}

func getRandomInt(min, max int) int {
	return mrand.Intn(max-min) + min
}
