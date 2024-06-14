package posts

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	postspb "posts/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var server Server

var mock sqlmock.Sqlmock

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	postspb.RegisterPostsServiceServer(s, &server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreatePost(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := postspb.NewPostsServiceClient(conn)

	server.db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer server.db.Close()

	username := "admin"
	postText := "some text"

	mock.ExpectExec("INSERT INTO posts (.+) VALUES (.+)").WithArgs(username, sqlmock.AnyArg(), sqlmock.AnyArg(), postText).WillReturnResult(sqlmock.NewResult(1, 1))

	resp, _ := client.CreatePost(ctx, &postspb.CreatePostReq{Username: username, Text: postText})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if resp.PostId == "" {
		t.Errorf("empty post_id")
	}
}

func TestGetPost(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := postspb.NewPostsServiceClient(conn)

	server.db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer server.db.Close()

	username := "admin"
	postId := "228"
	postText := "some text"

	posts_rows := sqlmock.NewRows([]string{"post_text"}).AddRow(postText)
	mock.ExpectQuery("SELECT post_text FROM posts WHERE post_id = ?").WithArgs(postId).WillReturnRows(posts_rows)

	resp, _ := client.GetPost(ctx, &postspb.GetPostReq{Username: username, PostId: postId})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if resp.Text != postText {
		t.Errorf("resp.Text != postText: %v != %v", resp.Text, postText)
	}
}
