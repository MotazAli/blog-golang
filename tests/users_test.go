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

// var router *gin.Engine
// var dbTest *mongo.Client
// func TestMain(m *testing.M){

// 	router = gin.Default()
// 	dbTest =configs.ConnectDBTest()
// 	usersController := controllers.CreateUsersController(dbTest)
// 	router.GET("/api/v1/users",usersController.GetAllUsers())
// 	router.GET("/api/v1/users/:id",usersController.GetUserById())
// 	router.POST("/api/v1/users",usersController.CreateUser())
// 	router.PUT("/api/v1/users/:id",usersController.UpdateUserById())
// 	router.DELETE("/api/v1/users/:id",usersController.DeleteUserById())
// 	rc := m.Run()
// 	os.Exit(rc)
// }

// func GGetResponseActualData [T models.UserTest|[]models.UserTest] (response *httptest.ResponseRecorder) T{
// 	expectedResponse := responses.Response{}
// 	var result  T
// 	json.NewDecoder(response.Body).Decode(&expectedResponse)
// 	data := expectedResponse.Result["data"]
// 	dataBytes,_ := json.Marshal(&data)
// 	json.Unmarshal(dataBytes,&result)
// 	return result
// }


func MakePostUserRequest(user models.User) (models.UserTest,*httptest.ResponseRecorder){
	// userCreateRequest := models.User{
	// 	Name: "Motaz Ali Test",
	// 	Email: "Motaz_teat@gmail.com",
	// 	Password: "123456",
	// }
	bodyCreateRequest,_ := json.Marshal(&user)
	//expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.UserLight{} }}
	//mockResponse ,_ :=json.Marshal(&expectedResponse)
	
	//router := GetRouter()
	//usersController := controllers.CreateUsersController(GetDatabase())
	//defer configs.DropCollection(dbTest,"users")
	//router.POST("/api/v1/users",usersController.CreateUser())

	request, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)

	// expectedCreateResponse := responses.Response{}
	// userCreateResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedCreateResponse)
	// data := expectedCreateResponse.Result["data"]

	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userCreateResponse)

	userCreateResponse := utilities.GGetResponseActualData[models.UserTest](response)
	//t.Error(userCreateResponse.Id)
	return userCreateResponse,response
}

func MakePutUserRequestByUserId(id string,user models.User) (models.UserTest,*httptest.ResponseRecorder){
	bodyUpdateRequest,_ := json.Marshal(&user)
	request, _ := http.NewRequest("PUT","/api/v1/users/"+id,bytes.NewBuffer(bodyUpdateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)

	// expectedUpdateResponse := responses.Response{}
	// userUpdateResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedUpdateResponse)
	// data := expectedUpdateResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userUpdateResponse)
	userUpdateResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userUpdateResponse,response
}


func MakeDeleteUserRequestById(id string)(models.UserTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("DELETE","/api/v1/users/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)

	// expectedDeleteResponse := responses.Response{}
	// userDeleteResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedDeleteResponse)
	// data := expectedDeleteResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userDeleteResponse)
	userDeleteResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userDeleteResponse,response
}


func MakeGetUserRequestByUserId(id string) (models.UserTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/users/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)

	// expectedResponse := responses.Response{}
	// userSelectResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedResponse)
	// data := expectedResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userSelectResponse)
	userSelectResponse := utilities.GGetResponseActualData[models.UserTest](response)
	return userSelectResponse,response
}



func TestGetAllUsersEmptyUsers(t *testing.T) {
	expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.UserLight{} }}
	mockResponse ,_ :=json.Marshal(&expectedResponse)
	
	
	//usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	
	// router.GET("/api/v1/users",usersController.GetAllUsers())
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
	// bodyCreateRequest,_ := json.Marshal(&userCreateRequest)
	// //expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.UserLight{} }}
	// //mockResponse ,_ :=json.Marshal(&expectedResponse)
	
	// //router := GetRouter()
	// //usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	// //router.POST("/api/v1/users",usersController.CreateUser())

	// request1, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	// response1 := httptest.NewRecorder()
	// router.ServeHTTP(response1,request1)
	MakePostUserRequest(userCreateRequest)
	
	//router.GET("/api/v1/users",usersController.GetAllUsers())
	request, _ := http.NewRequest("GET","/api/v1/users",nil)
	response := httptest.NewRecorder()

	RouterEngine.ServeHTTP(response,request)


	expectedResponse := responses.Response{}
	usersResponse := []models.UserTest{}
	json.NewDecoder(response.Body).Decode(&expectedResponse)
	data := expectedResponse.Result["data"]
	dataBytes,_ := json.Marshal(&data)
	json.Unmarshal(dataBytes,&usersResponse)

	//responseData , _ := ioutil.ReadAll(response.Body)
	assert.Equal(t,http.StatusOK,response.Code)
	assert.Equal(t,1,len(usersResponse))
	

}

func TestGetUserById(t *testing.T) {
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	
	// bodyCreateRequest,_ := json.Marshal(&userCreateRequest)
	
	// //router := GetRouter()
	// //usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	// //router.POST("/api/v1/users",usersController.CreateUser())

	// request1, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	// response1 := httptest.NewRecorder()
	// router.ServeHTTP(response1,request1)


	// expectedCreateResponse := responses.Response{}
	// userCreateResponse := models.UserTest{}
	// json.NewDecoder(response1.Body).Decode(&expectedCreateResponse)
	// data1 := expectedCreateResponse.Result["data"]

	// dataBytes1,_ := json.Marshal(&data1)
	// json.Unmarshal(dataBytes1,&userCreateResponse)
	// //t.Error(userResponse.Id)


	// request, _ := http.NewRequest("GET","/api/v1/users/"+userCreateResponse.Id,nil)
	// response := httptest.NewRecorder()
	// router.ServeHTTP(response,request)

	// expectedResponse := responses.Response{}
	// userSelectResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedResponse)
	// data := expectedResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userSelectResponse)

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
	// bodyRequest,_ := json.Marshal(&userRequest)
	// //expectedResponse := responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": models.User{} }}
	// //mockResponse ,_ :=json.Marshal(expectedResponse)

	// //router := GetRouter()
	// //usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	// //router.POST("/api/v1/users",usersController.CreateUser())

	// request, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyRequest))
	// response := httptest.NewRecorder()
	// router.ServeHTTP(response,request)


	// expectedResponse := responses.Response{}
	// userResponse := models.UserTest{}
	// //responseData , _ := ioutil.ReadAll(response.Body)
	// json.NewDecoder(response.Body).Decode(&expectedResponse)
	// //json.Unmarshal(responseData,&expectedResponse)
	// data := expectedResponse.Result["data"]
	// //mapstructure.Decode(data, &userResponse)
	// //t.Error(data)
	// //structValue := reflect.ValueOf(data).Elem()
    // //structFieldValue := structValue.FieldByName("Id")

	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userResponse)
	// //t.Error(userResponse.Id)


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
	//bodyCreateRequest,_ := json.Marshal(&userCreateRequest)
	//bodyUpdateRequest,_ := json.Marshal(&userUpdateRequest)
	//expectedResponse := responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": models.User{} }}
	//mockResponse ,_ :=json.Marshal(expectedResponse)

	//router := GetRouter()
	//usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	// //router.POST("/api/v1/users",usersController.CreateUser())

	// request1, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	// response1 := httptest.NewRecorder()
	// router.ServeHTTP(response1,request1)


	// expectedCreateResponse := responses.Response{}
	// userResponse := models.UserTest{}
	// json.NewDecoder(response1.Body).Decode(&expectedCreateResponse)
	// data1 := expectedCreateResponse.Result["data"]

	// dataBytes1,_ := json.Marshal(&data1)
	// json.Unmarshal(dataBytes1,&userResponse)
	// //t.Error(userResponse.Id)
	userCreateResponse , _ := MakePostUserRequest(userCreateRequest)
	userUpdateResponse, response := MakePutUserRequestByUserId(userCreateResponse.Id,userUpdateRequest)
	// request, _ := http.NewRequest("PUT","/api/v1/users/"+userCreateResponse.Id,bytes.NewBuffer(bodyUpdateRequest))
	// response := httptest.NewRecorder()
	// router.ServeHTTP(response,request)

	// expectedUpdateResponse := responses.Response{}
	// userUpdateResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedUpdateResponse)
	// data := expectedUpdateResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userUpdateResponse)

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
	
	// bodyCreateRequest,_ := json.Marshal(&userCreateRequest)
	// //expectedResponse := responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": models.User{} }}
	// //mockResponse ,_ :=json.Marshal(expectedResponse)

	// //router := GetRouter()
	// //usersController := controllers.CreateUsersController(GetDatabase())
	defer configs.DropCollection(DbTest,"users")
	// //router.POST("/api/v1/users",usersController.CreateUser())

	// request1, _ := http.NewRequest("POST","/api/v1/users",bytes.NewBuffer(bodyCreateRequest))
	// response1 := httptest.NewRecorder()
	// router.ServeHTTP(response1,request1)


	// expectedCreateResponse := responses.Response{}
	// userCreateResponse := models.UserTest{}
	// json.NewDecoder(response1.Body).Decode(&expectedCreateResponse)
	// data1 := expectedCreateResponse.Result["data"]

	// dataBytes1,_ := json.Marshal(&data1)
	// json.Unmarshal(dataBytes1,&userCreateResponse)
	// //t.Error(userResponse.Id)
	userCreateResponse , _ := MakePostUserRequest(userCreateRequest)
	userDeleteResponse, response := MakeDeleteUserRequestById(userCreateResponse.Id)
	// request, _ := http.NewRequest("DELETE","/api/v1/users/"+userCreateResponse.Id,nil)
	// response := httptest.NewRecorder()
	// router.ServeHTTP(response,request)

	// expectedDeleteResponse := responses.Response{}
	// userDeleteResponse := models.UserTest{}
	// json.NewDecoder(response.Body).Decode(&expectedDeleteResponse)
	// data := expectedDeleteResponse.Result["data"]
	// dataBytes,_ := json.Marshal(&data)
	// json.Unmarshal(dataBytes,&userDeleteResponse)

	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,userDeleteResponse.Id)
	assert.Equal(t,userCreateResponse.Id,userDeleteResponse.Id)
	assert.Equal(t,userCreateResponse.Name,userDeleteResponse.Name)
	assert.Equal(t,userCreateResponse.Email,userDeleteResponse.Email)
	assert.Equal(t,userCreateResponse.Password,userDeleteResponse.Password)
	assert.NotEmpty(t,userDeleteResponse.CreatedAt)
	assert.NotEmpty(t,userDeleteResponse.UpdatedAt)
}