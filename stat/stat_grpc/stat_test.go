package stat

import (
	"context"
	"log"
	"net"
	"slices"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	statpb "stat/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var server Server

var mock sqlmock.Sqlmock

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	statpb.RegisterStatServiceServer(s, &server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetPostStats(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := statpb.NewStatServiceClient(conn)

	server.db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer server.db.Close()

	postId := "228"
	var likes uint64 = 8
	var views uint64 = 1

	likes_rows := sqlmock.NewRows([]string{"count"}).AddRow(likes)
	views_rows := sqlmock.NewRows([]string{"count"}).AddRow(views)
	mock.ExpectQuery("SELECT(.*)").WithArgs(postId).WillReturnRows(likes_rows)
	mock.ExpectQuery("SELECT(.*)").WithArgs(postId).WillReturnRows(views_rows)

	resp, err := client.GetPostStats(ctx, &statpb.GetPostStatsReq{PostId: postId})

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if likes != resp.Likes {
		t.Errorf("likes != resp.Likes: %v != %v", likes, resp.Likes)
	}

	if views != resp.Views {
		t.Errorf("views != resp.Views: %v != %v", views, resp.Views)
	}
}

func TestGetTopUsers(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := statpb.NewStatServiceClient(conn)

	server.db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer server.db.Close()

	rows := sqlmock.NewRows([]string{"author", "cnt"}).
		AddRow("egg", 77777777).
		AddRow("pivo", 52).
		AddRow("vasek", 3)
	mock.ExpectQuery("SELECT(.*)").WithArgs().WillReturnRows(rows)

	resp, err := client.GetTopUsers(ctx, &statpb.GetTopUsersReq{})

	if err != nil {
		t.Fatal(err)
	}

	usernames := []string{"egg", "pivo", "vasek"}
	likes := []uint64{77777777, 52, 3}

	if !slices.Equal(usernames, resp.Username) {
		t.Errorf("usernames is not equal")
	}

	if !slices.Equal(likes, resp.Likes) {
		t.Errorf("likes is not equal")
	}
}
