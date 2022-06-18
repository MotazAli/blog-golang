package tests

import (
	//"blog/configs"
	"blog/controllers"
	"blog/models"
	"blog/responses"
	"bytes"
	"encoding/json"

	// "fmt"
	// "log"
	// "reflect"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.UserLight{} }}
	mockResponse ,_ :=json.Marshal(&expectedResponse)
	router := GetRouter()

	usersController := controllers.CreateUsersController(GetDatabase())
	router.GET("/api/v1/users",usersController.GetAllUsers())

	request, _ := http.NewRequest("GET","/api/v1/users",nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response,request)

	responseData , _ := ioutil.ReadAll(response.Body)
	assert.Equal(t,string(mockResponse),string(responseData))
	assert.Equal(t,http.StatusOK,response.Code)

}


func TestCreateUser(t *testing.T) {
	userRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	bodyRequest,_ := json.Marshal(&userRequest)
	//expectedResponse := responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": models.User{} }}
	//mockResponse ,_ :=json.Marshal(expectedResponse)

	router := GetRouter()
	usersController := controllers.CreateUsersController(GetDatabase())
	//defer configs.DropCollection(GetDatabase(),"users")
	router.POST("/api/v1/users",usersController.CreateUser())

	request, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyRequest))
	response := httptest.NewRecorder()
	router.ServeHTTP(response,request)


	expectedResponse := responses.Response{}
	userResponse := models.UserTest{}
	//responseData , _ := ioutil.ReadAll(response.Body)
	json.NewDecoder(response.Body).Decode(&expectedResponse)
	//json.Unmarshal(responseData,&expectedResponse)
	data := expectedResponse.Result["data"]
	//mapstructure.Decode(data, &userResponse)
	//t.Error(data)
	//structValue := reflect.ValueOf(data).Elem()
    //structFieldValue := structValue.FieldByName("Id")

	dataBytes,_ := json.Marshal(&data)
	json.Unmarshal(dataBytes,&userResponse)
	//t.Error(userResponse.Id)
	assert.Equal(t,http.StatusCreated,response.Code)
	assert.NotEmpty(t,userResponse.Id)
	assert.Equal(t,userRequest.Name,userResponse.Name)
	assert.Equal(t,userRequest.Email,userResponse.Email)
	assert.Equal(t,userRequest.Password,userResponse.Password)
	assert.NotEmpty(t,userResponse.CreatedAt)
	assert.NotEmpty(t,userResponse.UpdatedAt)
	

}
