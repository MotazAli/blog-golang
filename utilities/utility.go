package utilities

import (
	"blog/models"
	"blog/responses"
	"encoding/json"
	"net/http/httptest"

	"github.com/go-playground/validator"
)

var Validate = validator.New()

//that is mean T is the generic type which is slice of Element where the Element is any type
func GRemoveElementByIndex[T ~[]Element , Element any] (s T, index int) T {
    return append(s[:index], s[index+1:]...)
}


func GGetResponseActualData [T models.UserTest|models.PostTest|models.CommentTest|[]models.UserTest|[]models.PostTest|[]models.CommentTest] (response *httptest.ResponseRecorder) T{
	expectedResponse := responses.Response{}
	var result  T
	json.NewDecoder(response.Body).Decode(&expectedResponse)
	data := expectedResponse.Result["data"]
	dataBytes,_ := json.Marshal(&data)
	json.Unmarshal(dataBytes,&result)
	return result
}