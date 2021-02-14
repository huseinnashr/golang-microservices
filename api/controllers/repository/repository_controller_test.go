package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/nvdhunter/golang-microservices/clients/restclient"
	"github.com/nvdhunter/golang-microservices/domain/repository"
	"github.com/nvdhunter/golang-microservices/utils/errors"
	"github.com/nvdhunter/golang-microservices/utils/tests"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := tests.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.GetStatus())
	assert.EqualValues(t, "invalid json body", apiErr.GetMessage())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := tests.GetMockedContext(request, response)

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#list-repositories-for-the-authenticated-user"}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.GetStatus())
	assert.EqualValues(t, "Requires authentication", apiErr.GetMessage())
}

func TestCreateRepoNoError(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := tests.GetMockedContext(request, response)

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repository.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
