package posts

import (
	"context"
	"database/sql"
	postspb "posts/proto"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	dbConnStr = "user=posts password=super_secret_pass dbname=posts_db host=posts_db port=5432 sslmode=disable"
)

type Server struct {
	postspb.UnimplementedPostsServiceServer
	db *sql.DB
}

func NewServer() (server *Server, err error) {
	server = &Server{}
	server.db, err = sql.Open("postgres", dbConnStr)
	return
}

func (s *Server) CreatePost(ctx context.Context, req *postspb.CreatePostReq) (*postspb.CreatePostResp, error) {
	post_id := uuid.New().String() + "_" + req.Username
	_, err := s.db.ExecContext(ctx, "INSERT INTO posts (username, post_id, created_at, post_text) VALUES ($1, $2, $3, $4)", req.Username, post_id, time.Now(), req.Text)
	if err != nil {
		return &postspb.CreatePostResp{}, status.Errorf(codes.Internal, "failed to insert to posts relation while creating post for user %s", req.Username)
	}

	return &postspb.CreatePostResp{PostId: post_id}, nil
}

func (s *Server) UpdatePost(ctx context.Context, req *postspb.UpdatePostReq) (*postspb.UpdatePostResp, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return &postspb.UpdatePostResp{}, status.Errorf(codes.Internal, "failed to begin tx while updating post %s from user %s", req.PostId, req.Username)
	}
	defer tx.Rollback()

	count := 0
	tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE username = $1 AND post_id = $2", req.Username, req.PostId).Scan(&count)
	if count == 0 {
		return &postspb.UpdatePostResp{}, status.Errorf(codes.NotFound, "failed to found post %s from user %s", req.PostId, req.Username)
	}

	_, err = tx.ExecContext(ctx, "UPDATE posts SET post_text = $3 WHERE username = $1 AND post_id = $2", req.Username, req.PostId, req.Text)
	if err != nil {
		return &postspb.UpdatePostResp{}, status.Errorf(codes.Internal, "failed to update posts relation by post %s from user %s", req.PostId, req.Username)
	}

	err = tx.Commit()
	if err != nil {
		return &postspb.UpdatePostResp{}, status.Errorf(codes.Internal, "failed to commit tx updating post %s from user %s", req.PostId, req.Username)
	}

	return &postspb.UpdatePostResp{}, nil
}

func (s *Server) DeletePost(ctx context.Context, req *postspb.DeletePostReq) (*postspb.DeletePostResp, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return &postspb.DeletePostResp{}, status.Errorf(codes.Internal, "failed to begin tx while deleting post %s from user %s", req.PostId, req.Username)
	}
	defer tx.Rollback()

	count := 0
	tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM posts WHERE username = $1 AND post_id = $2", req.Username, req.PostId).Scan(&count)
	if count == 0 {
		return &postspb.DeletePostResp{}, status.Errorf(codes.NotFound, "failed to found post %s from user %s", req.PostId, req.Username)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM posts WHERE username = $1 AND post_id = $2", req.Username, req.PostId)
	if err != nil {
		return &postspb.DeletePostResp{}, status.Errorf(codes.Internal, "failed to delete from posts relation while deleting post %s from user %s", req.PostId, req.Username)
	}

	err = tx.Commit()
	if err != nil {
		return &postspb.DeletePostResp{}, status.Errorf(codes.Internal, "failed to commit tx deleting post %s from user %s", req.PostId, req.Username)
	}

	return &postspb.DeletePostResp{}, nil
}

func (s *Server) GetPost(ctx context.Context, req *postspb.GetPostReq) (*postspb.GetPostResp, error) {
	var post_text string
	err := s.db.QueryRowContext(ctx, "SELECT post_text FROM posts WHERE username = $1 AND post_id = $2", req.Username, req.PostId).Scan(&post_text)
	if err != nil {
		if err == sql.ErrNoRows {
			return &postspb.GetPostResp{}, status.Errorf(codes.NotFound, "there is no post %s for user %s", req.PostId, req.Username)
		} else {
			return &postspb.GetPostResp{}, status.Errorf(codes.Internal, "failed while select from posts relation while getting post %s for user %s", req.PostId, req.Username)
		}
	}

	return &postspb.GetPostResp{Text: post_text}, nil
}

func (s *Server) GetAllPosts(ctx context.Context, req *postspb.GetAllPostsReq) (*postspb.GetAllPostsResp, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT post_id, post_text FROM posts WHERE username = $1 ORDER BY created_at LIMIT $2 OFFSET $3", req.Username, req.Count, req.From)
	if err != nil {
		return &postspb.GetAllPostsResp{}, status.Errorf(codes.Internal, "failed while select from posts relation while getting posts for user %s", req.Username)
	}
	defer rows.Close()

	var post_id, post_text string
	post_ids := make([]string, 0)
	post_texts := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&post_id, &post_text)
		if err != nil {
			return &postspb.GetAllPostsResp{}, status.Errorf(codes.Internal, "failed while scannings from posts relation while getting posts for user %s", req.Username)
		}

		post_ids = append(post_ids, post_id)
		post_texts = append(post_texts, post_text)
	}

	return &postspb.GetAllPostsResp{PostIds: post_ids, Texts: post_texts}, nil
}
