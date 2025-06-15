package validator

import (
	"encoding/json"
	"fmt"
)

type ValidatorError map[string][]string

func (ve *ValidatorError) ToJson() ([]byte, error) {
	return json.Marshal(ve)
}

func (ve *ValidatorError) LogError() {
	fmt.Println(*ve)
}
