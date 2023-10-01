package model

// import "github.com/ethereum/go-ethereum/common"

type UserAnswer struct {
	RespondentAddress string
	TxId              string
	QuestionID        int64
	Content           string
}

// QuestionType  express type of answer of question
type AnswerType int

const (
	Choice AnswerType = iota // 0
	Text                     // 1
)
