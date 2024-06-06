package stat

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	statpb "stat/proto"
)

const (
	dbConnStr = "tcp://clickhouse:9000?username=default&password="
)

type Server struct {
	statpb.UnimplementedStatServiceServer
	db *sql.DB
}

func NewServer() (server *Server, err error) {
	server = &Server{}
	server.db, err = sql.Open("clickhouse", dbConnStr)
	return
}

func (s *Server) GetPostStats(ctx context.Context, req *statpb.GetPostStatsReq) (*statpb.GetPostStatsResp, error) {
	var likes_cnt, views_cnt uint64
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(DISTINCT user_id) FROM likes WHERE post_id = $1", req.PostId).Scan(&likes_cnt)
	if err != nil {
		return &statpb.GetPostStatsResp{}, err
	}

	err = s.db.QueryRowContext(ctx, "SELECT COUNT(DISTINCT user_id) FROM views WHERE post_id = $1", req.PostId).Scan(&views_cnt)
	if err != nil {
		return &statpb.GetPostStatsResp{}, err
	}

	fmt.Fprintf(os.Stderr, "likes: %v, views: %v\n", likes_cnt, views_cnt)
	return &statpb.GetPostStatsResp{Likes: likes_cnt, Views: views_cnt}, nil
}

func (s *Server) GetTopPosts(ctx context.Context, req *statpb.GetTopPostsReq) (*statpb.GetTopPostsResp, error) {
	var query string
	if req.SortBy == statpb.Metric_Likes {
		query = "SELECT post_id, COUNT(DISTINCT user_id) AS cnt FROM likes GROUP BY post_id ORDER BY cnt DESC LIMIT 5"
	} else {
		query = "SELECT post_id, COUNT(DISTINCT user_id) AS cnt FROM views GROUP BY post_id ORDER BY cnt DESC LIMIT 5"
	}

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return &statpb.GetTopPostsResp{}, err
	}
	defer rows.Close()

	var post_id string
	var cnt uint64
	post_ids := make([]string, 0)
	cnts := make([]uint64, 0)
	for rows.Next() {
		err = rows.Scan(&post_id, &cnt)
		if err != nil {
			return &statpb.GetTopPostsResp{}, err
		}

		post_ids = append(post_ids, post_id)
		cnts = append(cnts, cnt)
	}

	return &statpb.GetTopPostsResp{PostId: post_ids, Value: cnts}, nil
}

func (s *Server) GetTopUsers(ctx context.Context, req *statpb.GetTopUsersReq) (*statpb.GetTopUsersResp, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT author_id, COUNT(DISTINCT user_id) AS cnt FROM likes GROUP BY author_id ORDER BY cnt DESC LIMIT 3")
	if err != nil {
		return &statpb.GetTopUsersResp{}, err
	}
	defer rows.Close()

	var authorId uint64
	var cnt uint64
	authorIds := make([]uint64, 0)
	cnts := make([]uint64, 0)
	for rows.Next() {
		err = rows.Scan(&authorId, &cnt)
		if err != nil {
			return &statpb.GetTopUsersResp{}, err
		}

		authorIds = append(authorIds, authorId)
		cnts = append(cnts, cnt)
	}

	return &statpb.GetTopUsersResp{AuthorId: authorIds, Likes: cnts}, nil
}
