package entity

type Deck struct {
	ID string
	Name string
	OwnerID string
	Theme string
	IsPublic bool
	Tags []string
	CardIDs []string
	TargetLanguage string
	SourceLanguage string
}