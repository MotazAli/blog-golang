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



func MakePostForPostRequest(post models.PostCreateRequest) (models.PostTest,*httptest.ResponseRecorder){
	
	bodyCreateRequest,_ := json.Marshal(&post)
	request, _ := http.NewRequest("POST","/api/v1/posts",bytes.NewBuffer(bodyCreateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	postCreateResponse := utilities.GGetResponseActualData[models.PostTest](response)
	return postCreateResponse,response
}

func MakePutPostRequestByPostId(id string,post models.PostUpdateRequest) (models.PostTest,*httptest.ResponseRecorder){
	bodyUpdateRequest,_ := json.Marshal(&post)
	request, _ := http.NewRequest("PUT","/api/v1/posts/"+id,bytes.NewBuffer(bodyUpdateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	postUpdateResponse := utilities.GGetResponseActualData[models.PostTest](response)
	return postUpdateResponse,response
}


func MakeDeletePostRequestById(id string)(models.PostTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("DELETE","/api/v1/posts/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	postDeleteResponse := utilities.GGetResponseActualData[models.PostTest](response)
	return postDeleteResponse,response
}


func MakeGetPostRequestByPostId(id string) (models.PostTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/posts/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	postSelectResponse := utilities.GGetResponseActualData[models.PostTest](response)
	return postSelectResponse,response
}


func MakeGetAllPostRquest()([]models.PostTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/posts",nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	postsSelectResponse := utilities.GGetResponseActualData[[]models.PostTest](response)
	return postsSelectResponse,response
}



func TestGetAllPostsEmptyUsers(t *testing.T) {
	expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.PostLight{} }}
	mockResponse ,_ :=json.Marshal(&expectedResponse)
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	
	
	request, _ := http.NewRequest("GET","/api/v1/posts",nil)
	response := httptest.NewRecorder()

	RouterEngine.ServeHTTP(response,request)

	responseData , _ := ioutil.ReadAll(response.Body)
	assert.Equal(t,string(mockResponse),string(responseData))
	assert.Equal(t,http.StatusOK,response.Code)

}

func TestGetAllPostsNotEmptyUsers(t *testing.T) {
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userReponse,_ := MakePostUserRequest(userCreateRequest)

	postCreateRequest := models.PostCreateRequest{
		Title: "Motaz title",
		Body: "Motaz body",
		UserId: userReponse.Id,
	}

	MakePostForPostRequest(postCreateRequest)
	postsResponse,response := MakeGetAllPostRquest()
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.Equal(t,1,len(postsResponse))
	

}

func TestGetPostById(t *testing.T) {
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userReponse,_ := MakePostUserRequest(userCreateRequest)

	postCreateRequest := models.PostCreateRequest{
		Title: "Motaz title",
		Body: "Motaz body",
		UserId: userReponse.Id,
	}

	postCreateResponse,_ :=MakePostForPostRequest(postCreateRequest)

	postSelectResponse,response := MakeGetPostRequestByPostId(postCreateResponse.Id)

	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,postSelectResponse.Id)
	assert.Equal(t,postCreateResponse.Id,postSelectResponse.Id)
	assert.Equal(t,postCreateResponse.Title,postSelectResponse.Title)
	assert.Equal(t,postCreateResponse.Body,postSelectResponse.Body)
	assert.NotEmpty(t,postSelectResponse.CreatedAt)
	assert.NotEmpty(t,postSelectResponse.UpdatedAt)
}


func TestCreatePost(t *testing.T) {
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userReponse,response := MakePostUserRequest(userCreateRequest)

	postCreateRequest := models.PostCreateRequest{
		Title: "Motaz title",
		Body: "Motaz body",
		UserId: userReponse.Id,
	}

	postCreateResponse,_ :=MakePostForPostRequest(postCreateRequest)


	assert.Equal(t,http.StatusCreated,response.Code)
	assert.NotEmpty(t,postCreateResponse.Id)
	assert.Equal(t,postCreateRequest.Title,postCreateResponse.Title)
	assert.Equal(t,postCreateRequest.Body,postCreateResponse.Body)
	assert.NotEmpty(t,postCreateResponse.CreatedAt)
	assert.NotEmpty(t,postCreateResponse.UpdatedAt)
}


func TestUpdatePost(t *testing.T) {
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userReponse,_ := MakePostUserRequest(userCreateRequest)

	postCreateRequest := models.PostCreateRequest{
		Title: "Motaz title",
		Body: "Motaz body",
		UserId: userReponse.Id,
	}

	postCreateResponse,_ :=MakePostForPostRequest(postCreateRequest)

	postUpdateRequest := models.PostUpdateRequest{
		Title: "Motaz edit title",
		Body: "Motaz edit body",
	}

	postUpdateResponse, response := MakePutPostRequestByPostId(postCreateResponse.Id,postUpdateRequest)


	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,postUpdateResponse.Id)
	assert.Equal(t,postCreateResponse.Id,postUpdateResponse.Id)
	assert.Equal(t,postUpdateRequest.Title,postUpdateResponse.Title)
	assert.Equal(t,postUpdateRequest.Body,postUpdateResponse.Body)
	assert.NotEmpty(t,postUpdateResponse.CreatedAt)
	assert.NotEmpty(t,postUpdateResponse.UpdatedAt)
}



func TestDeletePostById(t *testing.T) {
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	userCreateRequest := models.User{
		Name: "Motaz Ali Test",
		Email: "Motaz_teat@gmail.com",
		Password: "123456",
	}
	userReponse,_ := MakePostUserRequest(userCreateRequest)

	postCreateRequest := models.PostCreateRequest{
		Title: "Motaz title",
		Body: "Motaz body",
		UserId: userReponse.Id,
	}

	postCreateResponse,_ :=MakePostForPostRequest(postCreateRequest)

	postDeleteResponse, response := MakeDeletePostRequestById(postCreateResponse.Id)
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,postDeleteResponse.Id)
	assert.Equal(t,postCreateResponse.Id,postDeleteResponse.Id)
	assert.Equal(t,postCreateResponse.Title,postDeleteResponse.Title)
	assert.Equal(t,postCreateResponse.Body,postDeleteResponse.Body)
	assert.NotEmpty(t,postDeleteResponse.CreatedAt)
	assert.NotEmpty(t,postDeleteResponse.UpdatedAt)
}