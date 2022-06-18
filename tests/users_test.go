package tests

import (
	"blog/configs"
	"blog/utilities"
	"blog/models"
	"blog/responses"
	"bytes"
	"encoding/json"



	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	
	"github.com/stretchr/testify/assert"
)



func MakePostUserRequest(user models.User) (models.UserTest,*httptest.ResponseRecorder){
	bodyCreateRequest,_ := json.Marshal(&user)
	request, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	userCreateResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userCreateResponse,response
}

func MakePutUserRequestByUserId(id string,user models.User) (models.UserTest,*httptest.ResponseRecorder){
	bodyUpdateRequest,_ := json.Marshal(&user)
	request, _ := http.NewRequest("PUT","/api/v1/users/"+id,bytes.NewBuffer(bodyUpdateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	userUpdateResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userUpdateResponse,response
}


func MakeDeleteUserRequestById(id string)(models.UserTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("DELETE","/api/v1/users/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	userDeleteResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userDeleteResponse,response
}


func MakeGetUserRequestByUserId(id string) (models.UserTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/users/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	userSelectResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userSelectResponse,response
}



func TestGetAllUsersEmptyUsers(t *testing.T) {
	expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.UserLight{} }}
	mockResponse ,_ :=json.Marshal(&expectedResponse)

	defer configs.DropCollection(DbTest,"users")
	
	request, _ := http.NewRequest("GET","/api/v1/users",nil)
	response := httptest.NewRecorder()

	RouterEngine.ServeHTTP(response,request)

	responseData , _ := ioutil.ReadAll(response.Body)
	assert.Equal(t,string(mockResponse),string(responseData))
	assert.Equal(t,http.StatusOK,response.Code)

}

func TestGetAllUsersNotEmptyUsers(t *testing.T) {
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}

	defer configs.DropCollection(DbTest,"users")

	MakePostUserRequest(userCreateRequest)
	
	request, _ := http.NewRequest("GET","/api/v1/users",nil)
	response := httptest.NewRecorder()

	RouterEngine.ServeHTTP(response,request)


	expectedResponse := responses.Response{}
	usersResponse := []models.UserTest{}
	json.NewDecoder(response.Body).Decode(&expectedResponse)
	data := expectedResponse.Result["data"]
	dataBytes,_ := json.Marshal(&data)
	json.Unmarshal(dataBytes,&usersResponse)

	assert.Equal(t,http.StatusOK,response.Code)
	assert.Equal(t,1,len(usersResponse))
	

}

func TestGetUserById(t *testing.T) {
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	

	defer configs.DropCollection(DbTest,"users")
	
	userCreateResponse,_ := MakePostUserRequest(userCreateRequest)
	userSelectResponse,response := MakeGetUserRequestByUserId(userCreateResponse.Id)

	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,userSelectResponse.Id)
	assert.Equal(t,userCreateResponse.Id,userSelectResponse.Id)
	assert.Equal(t,userCreateResponse.Name,userSelectResponse.Name)
	assert.Equal(t,userCreateResponse.Email,userSelectResponse.Email)
	assert.Equal(t,userCreateResponse.Password,userSelectResponse.Password)
	assert.NotEmpty(t,userSelectResponse.CreatedAt)
	assert.NotEmpty(t,userSelectResponse.UpdatedAt)
}


func TestCreateUser(t *testing.T) {
	userRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}

	defer configs.DropCollection(DbTest,"users")

	userResponse , response := MakePostUserRequest(userRequest)

	assert.Equal(t,http.StatusCreated,response.Code)
	assert.NotEmpty(t,userResponse.Id)
	assert.Equal(t,userRequest.Name,userResponse.Name)
	assert.Equal(t,userRequest.Email,userResponse.Email)
	assert.Equal(t,userRequest.Password,userResponse.Password)
	assert.NotEmpty(t,userResponse.CreatedAt)
	assert.NotEmpty(t,userResponse.UpdatedAt)
}


func TestUpdateUser(t *testing.T) {
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userUpdateRequest := models.User{
		Name: "Motaz Ali Test Edit",
		Email: "Motaz_teat_edit@gmail.com",
		Password: "111222",
	}
	
	defer configs.DropCollection(DbTest,"users")
	
	userCreateResponse , _ := MakePostUserRequest(userCreateRequest)
	userUpdateResponse, response := MakePutUserRequestByUserId(userCreateResponse.Id,userUpdateRequest)
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,userUpdateResponse.Id)
	assert.Equal(t,userCreateResponse.Id,userUpdateResponse.Id)
	assert.Equal(t,userUpdateRequest.Name,userUpdateResponse.Name)
	assert.Equal(t,userUpdateRequest.Email,userUpdateResponse.Email)
	assert.Equal(t,userUpdateRequest.Password,userUpdateResponse.Password)
	assert.NotEmpty(t,userUpdateResponse.CreatedAt)
	assert.NotEmpty(t,userUpdateResponse.UpdatedAt)
}



func TestDeleteUserById(t *testing.T) {
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	
	
	defer configs.DropCollection(DbTest,"users")
	
	userCreateResponse , _ := MakePostUserRequest(userCreateRequest)
	userDeleteResponse, response := MakeDeleteUserRequestById(userCreateResponse.Id)
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,userDeleteResponse.Id)
	assert.Equal(t,userCreateResponse.Id,userDeleteResponse.Id)
	assert.Equal(t,userCreateResponse.Name,userDeleteResponse.Name)
	assert.Equal(t,userCreateResponse.Email,userDeleteResponse.Email)
	assert.Equal(t,userCreateResponse.Password,userDeleteResponse.Password)
	assert.NotEmpty(t,userDeleteResponse.CreatedAt)
	assert.NotEmpty(t,userDeleteResponse.UpdatedAt)
}