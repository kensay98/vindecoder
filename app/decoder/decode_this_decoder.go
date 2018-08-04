package decoder

import (
	"github.com/kensay98/vindecoder/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var log = logger.GetLogger()

const ( // DecodeThis response statuses
	success               = "SUCCESS"
	invalidParameterError = "PARAMERR"
	invalidApiKeyError    = "SECERR"
	checkError            = "CHECKERR"
	notFoundError         = "NOTFOUND"
)

var messageByError = map[string]string{
	invalidParameterError: "Provided api key is invalid.",
	invalidApiKeyError:    "Provided api key is invalid.",
	checkError:            "",
	notFoundError:         "Requested vin is not found",
}

type (
	DecodeThisDecoder struct {
		apiKey string
	}

	decodeThisResponse struct {
		Decode struct {
			Status   string    `json:"status"`
			Vehicles []vehicle `json:"vehicle"`
		} `json:"decode"`
	}

	vehicle struct {
		Make  string `json:"make"`
		Model string `json:"model"`
		Year  string `json:"year"`
		Equip []struct {
			Name  string `json:"name"`
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"Equip"`
	}
)

func (decoder *DecodeThisDecoder) Decode(vin string) (decodedVin DecodedVin, err error) {
	vehicle, err := decoder.sendRequest(vin)

	if err != nil {
		return
	}

	equipMap := map[string]string{}

	for _, e := range vehicle.Equip {
		equipMap[e.Name] = e.Value
	}

	decodedVin.Vin = vin
	decodedVin.Year = vehicle.Year
	decodedVin.Make = vehicle.Make
	decodedVin.Model = vehicle.Model

	if weight, contains := equipMap["Standard GVWR"]; contains {
		decodedVin.Weight = weight
	}

	return
}

func (decoder *DecodeThisDecoder) sendRequest(vin string) (v vehicle, err error) {
	url := "https://www.decodethis.com/webservices/decodes/%s/%s/1.JSON"
	httpResp, err := http.Get(fmt.Sprintf(url, vin, decoder.apiKey))

	if err != nil {
		return
	}

	if httpResp.StatusCode != 200 {
		log.Error("Decode service returns error code " + strconv.Itoa(httpResp.StatusCode))
	}

	response := decodeThisResponse{}
	err = json.NewDecoder(httpResp.Body).Decode(&response)
	if err != nil {
		return
	}

	if response.Decode.Status != success {
		errorMessage, ok := messageByError[response.Decode.Status]
		if ok {
			return v, fmt.Errorf(errorMessage)
		} else {
			return v, fmt.Errorf("unknown error status")
		}
	}

	v = response.Decode.Vehicles[0]
	return
}

func NewDecodeThisDecoder(apiKey string) *DecodeThisDecoder {
	return &DecodeThisDecoder{apiKey: apiKey}
}
