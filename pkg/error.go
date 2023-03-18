package pkg

import (
	"encoding/json"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
)

func ErrorHandler(error error) []byte {
	errStruct := &dto.Error{}
	errStruct.Error = error.Error()

	body, err := json.Marshal(&errStruct)

	if err != nil {
		return nil
	}

	return body
}
