/*
 * User Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package auth

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	postspb "posts/proto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	_ "github.com/lib/pq"
)

var (
	authDB      *sql.DB
	grpc_conn   *grpc.ClientConn
	grpc_client postspb.PostsServiceClient
)

const (
	dbConnStr        = "user=auth password=super_secret_pass dbname=auth_db host=auth_db port=5432 sslmode=disable"
	expiresDeltaSecs = 10000
)

func init() {
	var err error
	authDB, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	grpc_conn, err = grpc.Dial("dns:///posts:11777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	grpc_client = postspb.NewPostsServiceClient(grpc_conn)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	var user_with_pass LoginBody
	if err := json.NewDecoder(r.Body).Decode(&user_with_pass); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to decode JSON")
		return
	}
	if user_with_pass.Username == "" || user_with_pass.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "There is no username or password field in JSON, or one of them (or both) is empty")
		return
	}

	tx, err := authDB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	defer tx.Rollback()

	true_pass_md5_sum := make([]byte, 16)
	err = tx.QueryRow("SELECT password_hash FROM users WHERE username = $1", user_with_pass.Username).Scan(&true_pass_md5_sum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	cur_pass_md5_sum := md5.Sum([]byte(user_with_pass.Password))
	if slices.Compare(true_pass_md5_sum, cur_pass_md5_sum[:]) != 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "There is no such user or password is wrong")
		return
	}

	// check if there is an existing session (and return it if it exists)
	var token string
	var created time.Time
	err = tx.QueryRow("SELECT token, created FROM user_sessions WHERE username = $1", user_with_pass.Username).Scan(&token, &created)
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	if err == nil && time.Since(created).Seconds() < expiresDeltaSecs {
		err = tx.Commit()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}

		w.Header().Add("Set-Cookie", token)
		fmt.Fprintln(w, "User has been logined successfully")
		return
	}

	cur_token := uuid.New().String()
	cur_created := time.Now()

	if err == nil {
		_, err = tx.Exec("UPDATE user_sessions SET token = $2, created = $3 WHERE username = $1", user_with_pass.Username, cur_token, cur_created)
	} else {
		_, err = tx.Exec("INSERT INTO user_sessions (username, token, created) VALUES ($1, $2, $3)", user_with_pass.Username, cur_token, cur_created)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Set-Cookie", cur_token)
	fmt.Fprintln(w, "User has been logined successfully")
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	var user_with_pass RegisterBody
	if err := json.NewDecoder(r.Body).Decode(&user_with_pass); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to decode JSON")
		return
	}
	if user_with_pass.Username == "" || user_with_pass.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "There is no username or password field in JSON, or one of them (or both) is empty")
		return
	}

	tx, err := authDB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	defer tx.Rollback()

	count := 0
	tx.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", user_with_pass.Username).Scan(&count)
	if count > 0 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "User with this name already registered")
		return
	}

	pass_md5_sum := md5.Sum([]byte(user_with_pass.Password))
	_, err = tx.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", user_with_pass.Username, pass_md5_sum[:])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	_, err = tx.Exec("INSERT INTO user_data (username) VALUES ($1)", user_with_pass.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, "User has been registered")
}

func UpdatePut(w http.ResponseWriter, r *http.Request) {
	var user_updates UpdateBody
	if err := json.NewDecoder(r.Body).Decode(&user_updates); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to decode JSON")
		return
	}

	token := r.Header.Get("Cookie")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User must specify Cookie header")
		return
	}

	tx, err := authDB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	defer tx.Rollback()

	var username string
	var created time.Time
	err = tx.QueryRow("SELECT username, created FROM user_sessions WHERE token = $1", token).Scan(&username, &created)
	if err != nil || time.Since(created).Seconds() >= expiresDeltaSecs {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Cookie is invalid")
		return
	}

	// just bruteforce JSON and evaluate UPDATE query
	// UPDATE user_data SET firstname = $2, lastname = $3, ... WHERE username = $1
	set_stmt := ""
	seqnum := 2
	fl := false
	params := make([]interface{}, 0)
	params = append(params, username)
	if user_updates.FirstName != "" {
		set_stmt += fmt.Sprintf("firstname = $%d", seqnum)
		seqnum++
		fl = true
		params = append(params, user_updates.FirstName)
	}
	if user_updates.LastName != "" {
		if !fl {
			set_stmt += fmt.Sprintf("lastname = $%d", seqnum)
		} else {
			set_stmt += fmt.Sprintf(", lastname = $%d", seqnum)
		}
		seqnum++
		fl = true
		params = append(params, user_updates.LastName)
	}
	if user_updates.BirthDate != "" {
		if !fl {
			set_stmt += fmt.Sprintf("birthdate = $%d", seqnum)
		} else {
			set_stmt += fmt.Sprintf(", birthdate = $%d", seqnum)
		}
		seqnum++
		fl = true
		params = append(params, user_updates.BirthDate)
	}
	if user_updates.Email != "" {
		if !fl {
			set_stmt += fmt.Sprintf("email = $%d", seqnum)
		} else {
			set_stmt += fmt.Sprintf(", email = $%d", seqnum)
		}
		seqnum++
		fl = true
		params = append(params, user_updates.Email)
	}
	if user_updates.Telephone != "" {
		if !fl {
			set_stmt += fmt.Sprintf("telephone = $%d", seqnum)
		} else {
			set_stmt += fmt.Sprintf(", telephone = $%d", seqnum)
		}
		seqnum++
		fl = true
		params = append(params, user_updates.Telephone)
	}

	_, err = tx.Exec("UPDATE user_data SET "+set_stmt+" WHERE username = $1", params...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, "Data was updated")
}

func GetUsername(w http.ResponseWriter, r *http.Request) (string, error) {
	token := r.Header.Get("Cookie")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User must specify Cookie header")
		return "", errors.New("")
	}

	var username string
	var created time.Time
	err := authDB.QueryRow("SELECT username, created FROM user_sessions WHERE token = $1", token).Scan(&username, &created)
	if err != nil || time.Since(created).Seconds() >= expiresDeltaSecs {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Cookie is invalid")
		return "", errors.New("")
	}

	return username, nil
}

func CreatePostPost(w http.ResponseWriter, r *http.Request) {
	var body CreatePostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to decode JSON")
		return
	}

	username, err := GetUsername(w, r)
	if err != nil {
		return
	}

	grpc_resp, err := grpc_client.CreatePost(context.Background(), &postspb.CreatePostReq{Username: username, Text: body.Text})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		return
	}

	http_resp := CreatePostResponse{PostId: grpc_resp.PostId}
	http_resp_bytes, err := json.Marshal(http_resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		return
	}

	fmt.Fprintln(w, string(http_resp_bytes))
}

func DeletePostPostIdDelete(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(w, r)
	if err != nil {
		return
	}

	post_id := mux.Vars(r)["post_id"]
	_, err = grpc_client.DeletePost(context.Background(), &postspb.DeletePostReq{Username: username, PostId: post_id})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "There is no such user with this post_id")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		}
		return
	}
	fmt.Fprintln(w, "Post has been deleted succesfully")
}

func GetPostPostIdGet(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(w, r)
	if err != nil {
		return
	}

	post_id := mux.Vars(r)["post_id"]
	grpc_resp, err := grpc_client.GetPost(context.Background(), &postspb.GetPostReq{Username: username, PostId: post_id})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "There is no such user with this post_id")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		}
		return
	}

	fmt.Fprintln(w, grpc_resp.Text)
}

func GetPostsGet(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(w, r)
	if err != nil {
		return
	}

	from_str, count_str := r.Header.Get("From"), r.Header.Get("Count")
	if len(from_str) == 0 || len(count_str) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User must specify From and Count headers properly")
		return
	}

	from, err := strconv.ParseUint(from_str, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User must specify From and Count headers properly")
		return
	}

	count, err := strconv.ParseUint(count_str, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User must specify From and Count headers properly")
		return
	}

	resp, err := grpc_client.GetAllPosts(context.Background(), &postspb.GetAllPostsReq{Username: username, From: from, Count: count})
	if err != nil || len(resp.PostIds) != len(resp.Texts) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		return
	}

	for i := 0; i < len(resp.PostIds); i++ {
		fmt.Fprintf(w, "Post #%d with post_id %s:\n", i, resp.PostIds[i])
		fmt.Fprintf(w, "%s\n\n", resp.Texts[i])
	}
}

func UpdatePostPostIdPut(w http.ResponseWriter, r *http.Request) {
	var body UpdatePostPostIdBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to decode JSON")
		return
	}

	username, err := GetUsername(w, r)
	if err != nil {
		return
	}

	post_id := mux.Vars(r)["post_id"]
	_, err = grpc_client.UpdatePost(context.Background(), &postspb.UpdatePostReq{Username: username, Text: body.Text, PostId: post_id})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "There is no such user with this post_id")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Some internal error (with postgres, for example)")
		}
		return
	}

	fmt.Fprintln(w, "Post has been updated succesfully")
}
