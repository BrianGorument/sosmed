package posts

type IPostService interface {
	CreatePosting(req CreatePostRequest, users UserData) (*PostResponse, error)
}

type IPostRepository interface {
	InsertPosting(input Post_Content) (*PostResponse, error)
}
