package project

type Project struct {
	ID             string
	Name           string
	Color          string
	Order          int
	CommentCount   int
	IsShared       bool
	IsFavorite     bool
	IsInboxProject bool
	IsTeamInbox    bool
	URL            string
}
