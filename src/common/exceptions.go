package common

import "fmt"

type BizError struct {
	Code int
	Msg  string
}

func (b *BizError) Error() string {
	return fmt.Sprintf("Bizerror: code=%d, msg=%s", b.Code, b.Msg)
}

var RecordNotExist = BizError{Code: 1, Msg: "record not exist"}
var WrongPassword = BizError{Code: 2, Msg: "wrong password"}
var AuthorNameRepeatError = BizError{Code: 1000, Msg: "author create error: name repeat"}
