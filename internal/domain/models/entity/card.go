package entity

type Rating uint8

const (
	EASY Rating = iota
	MEDIUM 
	HARD
	FORGET
)

type Card struct {
	ID string
	DeckID string
	Front string
	Back string
	Rating Rating
	NextReview int64
	Interval int64
	Repetitions int64
	Ease float64
	ReviewCount int64
}