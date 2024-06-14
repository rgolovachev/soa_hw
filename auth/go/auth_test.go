package auth

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	mock_postspb "mock_posts"
	postspb "posts/proto"
)

var mock sqlmock.Sqlmock

type MockTokenCreator struct {
	mockUUID uuid.UUID
}

type MockTimestampCreator struct {
	mockTime time.Time
}

func (c *MockTokenCreator) New() uuid.UUID {
	return c.mockUUID
}

func (c *MockTimestampCreator) Now() time.Time {
	return c.mockTime
}

func TestRegisterUser(t *testing.T) {
	var err error

	authDB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer authDB.Close()

	mockUUID, _ := uuid.FromBytes([]byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4})
	tokenCreator = &MockTokenCreator{mockUUID: mockUUID}
	mockTs := time.Now()
	tsCreator = &MockTimestampCreator{mockTime: mockTs}

	username := "admin"
	password := "admin"
	md5Sum := md5.Sum([]byte(password))

	mock.ExpectBegin()

	users_rows := sqlmock.NewRows([]string{"password_hash"}).AddRow(md5Sum[:])
	user_sessions_rows := sqlmock.NewRows([]string{"token", "created"})
	mock.ExpectQuery("SELECT password_hash FROM users WHERE username = ?").WithArgs(username).WillReturnRows(users_rows)
	mock.ExpectQuery("SELECT token, created FROM user_sessions WHERE username = ?").WithArgs(username).WillReturnRows(user_sessions_rows)
	mock.ExpectExec("INSERT INTO user_sessions (.+) VALUES (.+)").WithArgs(username, mockUUID.String(), mockTs).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	reqBody := LoginBody{Username: username, Password: password}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginPost)
	handler.ServeHTTP(rr, req)

	log.Println(rr.Body)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status: expected is %v, current is %v", http.StatusOK, status)
	}
}

func TestGetPost(t *testing.T) {
	var err error

	authDB, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer authDB.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	grpc_posts_client = mock_postspb.NewMockPostsServiceClient(ctrl)
	username := "admin"
	cookie := "pechenka"
	postId := "228"
	text := "tttext"

	req, err := http.NewRequest("GET", "/get_post/"+postId, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Cookie", cookie)

	userSessionsRows := sqlmock.NewRows([]string{"username", "created"}).AddRow(username, time.Now())
	mock.ExpectQuery("SELECT username, created FROM user_sessions WHERE token = ?").WithArgs(cookie).WillReturnRows(userSessionsRows)

	mockGrpcClient, _ := grpc_posts_client.(*mock_postspb.MockPostsServiceClient)
	mockGrpcClient.EXPECT().GetPost(gomock.Any(), &postspb.GetPostReq{Username: username, PostId: postId}).Return(&postspb.GetPostResp{Text: text}, nil)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/get_post/{post_id}", GetPostPostIdGet)
	r.ServeHTTP(rr, req)

	responseBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status: expected is %v, current is %v", http.StatusOK, status)
	}

	if strings.TrimSpace(string(responseBody)) != text {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			string(responseBody), text)
	}
}
