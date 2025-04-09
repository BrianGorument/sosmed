package interactions

type IInteractionService interface {
	CreateCommentService(req InteractRequest, user UserData) (*InteractResponse, error)
}

type IInteractionRepository interface {
	InsertInteraction(input Post_Interactions) (*InteractResponse, error)
}
