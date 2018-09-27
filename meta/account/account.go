package account

import (
	"github.com/linkchain/common/serialize"
	"github.com/linkchain/meta"
)

type IAccountID interface {
	GetString() string
}

type IAccount interface {
	//acount management
	ChangeAmount(meta.IAmount) meta.IAmount

	GetAmount() meta.IAmount

	GetAccountID() IAccountID

	GetNounce() uint32

	SetNounce(nounce uint32) error

	CheckNounce(nounce uint32) bool

	//serialize
	serialize.ISerialize
}