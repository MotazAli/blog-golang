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



func MakePostCommentRequest(comment models.CommentCreateRequest) (models.CommentTest,*httptest.ResponseRecorder){
	
	bodyCreateRequest,_ := json.Marshal(&comment)
	request, _ := http.NewRequest("POST","/api/v1/comments",bytes.NewBuffer(bodyCreateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	commentCreateResponse := utilities.GGetResponseActualData[models.CommentTest](response)
	return commentCreateResponse,response
}

func MakePutCommentRequestByCommentId(id string,comment models.CommentUpdateRequest) (models.CommentTest,*httptest.ResponseRecorder){
	bodyUpdateRequest,_ := json.Marshal(&comment)
	request, _ := http.NewRequest("PUT","/api/v1/comments/"+id,bytes.NewBuffer(bodyUpdateRequest))
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	commentUpdateResponse := utilities.GGetResponseActualData[models.CommentTest](response)
	return commentUpdateResponse,response
}


func MakeDeleteCommentRequestById(id string)(models.CommentTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("DELETE","/api/v1/comments/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	commentDeleteResponse := utilities.GGetResponseActualData[models.CommentTest](response)
	return commentDeleteResponse,response
}


func MakeGetCommentRequestByCommentId(id string) (models.CommentTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/comments/"+id,nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	commentSelectResponse := utilities.GGetResponseActualData[models.CommentTest](response)
	return commentSelectResponse,response
}


func MakeGetAllCommentRquest()([]models.CommentTest,*httptest.ResponseRecorder){
	request, _ := http.NewRequest("GET","/api/v1/comments",nil)
	response := httptest.NewRecorder()
	RouterEngine.ServeHTTP(response,request)
	commentsSelectResponse := utilities.GGetResponseActualData[[]models.CommentTest](response)
	return commentsSelectResponse,response
}



func TestGetAllComments_EmptyComments(t *testing.T) {
	expectedResponse := responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": []models.CommentLight{} }}
	mockResponse ,_ :=json.Marshal(&expectedResponse)
	defer configs.DropCollection(DbTest,"comments")
	defer configs.DropCollection(DbTest,"posts")
	defer configs.DropCollection(DbTest,"users")
	
	
	request, _ := http.NewRequest("GET","/api/v1/comments",nil)
	response := httptest.NewRecorder()

	RouterEngine.ServeHTTP(response,request)

	responseData , _ := ioutil.ReadAll(response.Body)
	assert.Equal(t,string(mockResponse),string(responseData))
	assert.Equal(t,http.StatusOK,response.Code)

}

func TestGetAllComments_NotEmptyComments(t *testing.T) {
	defer configs.DropCollection(DbTest,"comments")
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

	postResponse,_ :=MakePostForPostRequest(postCreateRequest)
	 
	commentCreateRequest := models.CommentCreateRequest{
		Body: "Motaz comment body",
		UserId: userReponse.Id,
		PostId: postResponse.Id,
	}

	MakePostCommentRequest(commentCreateRequest)
	commentsResponse, response := MakeGetAllCommentRquest()
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.Equal(t,1,len(commentsResponse))
	

}

func TestGetCommentById(t *testing.T) {
	defer configs.DropCollection(DbTest,"comments")
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

	 
	commentCreateRequest := models.CommentCreateRequest{
		Body: "Motaz comment body",
		UserId: userReponse.Id,
		PostId: postCreateResponse.Id,
	}
	commentCreateResponse , _ :=MakePostCommentRequest(commentCreateRequest)


	commentSelectResponse,response := MakeGetCommentRequestByCommentId(commentCreateResponse.Id)

	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,commentSelectResponse.Id)
	assert.Equal(t,commentCreateResponse.Id,commentSelectResponse.Id)
	assert.Equal(t,commentCreateResponse.Body,commentSelectResponse.Body)
	assert.NotEmpty(t,commentSelectResponse.CreatedAt)
	assert.NotEmpty(t,commentSelectResponse.UpdatedAt)
}


func TestCreateComment(t *testing.T) {
	defer configs.DropCollection(DbTest,"comments")
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
	commentCreateRequest := models.CommentCreateRequest{
		Body: "Motaz comment body",
		UserId: userReponse.Id,
		PostId: postCreateResponse.Id,
	}
	commentCreateResponse ,response :=MakePostCommentRequest(commentCreateRequest)



	assert.Equal(t,http.StatusCreated,response.Code)
	assert.NotEmpty(t,commentCreateResponse.Id)
	assert.Equal(t,commentCreateRequest.Body,commentCreateResponse.Body)
	assert.NotEmpty(t,commentCreateResponse.CreatedAt)
	assert.NotEmpty(t,commentCreateResponse.UpdatedAt)
}


func TestUpdateComments(t *testing.T) {
	defer configs.DropCollection(DbTest,"comments")
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

	commentCreateRequest := models.CommentCreateRequest{
		Body: "Motaz comment body",
		UserId: userReponse.Id,
		PostId: postCreateResponse.Id,
	}
	commentCreateResponse ,_ :=MakePostCommentRequest(commentCreateRequest)


	commentUpdateRequest := models.CommentUpdateRequest{
		Body: "Motaz edit comment body",
	}

	commentUpdateResponse, response := MakePutCommentRequestByCommentId(commentCreateResponse.Id,commentUpdateRequest)


	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,commentUpdateResponse.Id)
	assert.Equal(t,commentCreateResponse.Id,commentUpdateResponse.Id)
	assert.Equal(t,commentUpdateRequest.Body,commentUpdateResponse.Body)
	assert.NotEmpty(t,commentUpdateResponse.CreatedAt)
	assert.NotEmpty(t,commentUpdateResponse.UpdatedAt)
}



func TestDeleteCommentById(t *testing.T) {
	defer configs.DropCollection(DbTest,"comments")
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


	commentCreateRequest := models.CommentCreateRequest{
		Body: "Motaz comment body",
		UserId: userReponse.Id,
		PostId: postCreateResponse.Id,
	}
	commentCreateResponse ,_ :=MakePostCommentRequest(commentCreateRequest)


	commentDeleteResponse, response := MakeDeleteCommentRequestById(commentCreateResponse.Id)
	
	assert.Equal(t,http.StatusOK,response.Code)
	assert.NotEmpty(t,commentDeleteResponse.Id)
	assert.Equal(t,commentCreateResponse.Id,commentDeleteResponse.Id)
	assert.Equal(t,commentCreateResponse.Body,commentDeleteResponse.Body)
	assert.NotEmpty(t,commentDeleteResponse.CreatedAt)
	assert.NotEmpty(t,commentDeleteResponse.UpdatedAt)
}