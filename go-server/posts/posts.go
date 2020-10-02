package posts

import (
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetPosts(ctx context.Context, in *Empty) (*PostList, error) {
	return &PostList{
		Posts: []*Post{
			&Post{
				Id:    1,
				Title: "Titulo 1",
				Text:  "bla bla bla",
			},
			&Post{
				Id:    2,
				Title: "Titulo 2",
				Text:  "bla bla bla",
			},
		},
	}, nil
}
