package handlers

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestVinHandler(t *testing.T) {
	dataTable := map[string]string {
		"KMHDH4AH9FU230302": `{
			"vin": "KMHDH4AH9FU230302",
			"make": "Hyundai",
			"model": "Elantra",
			"year": "2015",
			"color": "",
			"weight": "3880",
			"type": ""
		}`,
	}

	for key, expectedResult := range dataTable {
		// Given
		url := fmt.Sprintf("http://localhost:9999/vin/%s", key)

		// When
		request, _ := http.NewRequest("GET", url, nil)
		response := httptest.NewRecorder()
		GetApp().router.ServeHTTP(response, request)

		// Then
		assert.Equal(t, 200, response.Code)
		responseString := string(response.Body.Bytes())
		assert.Equal(t,  _trimAll(expectedResult), _trimAll(responseString))
	}
}

func _trimAll(str string) (string) {
	replacedTabs := strings.Replace(str, "\t", "", -1)
	replacedSpaces := strings.Replace(replacedTabs, " ", "", -1)
	replacedN := strings.Replace(replacedSpaces, "\n", "", -1)
	return replacedN
}