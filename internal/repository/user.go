package repository

import (
	"context"
	"github.com/sysatom/framework/ent"
	"github.com/sysatom/framework/ent/merchantaccount"
)

type UserRepository struct {
	store *ent.Client
}

func NewUserRepository(store *ent.Client) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (s *UserRepository) GetMerchantAccount(ctx context.Context, username string) (user *ent.MerchantAccount, err error) {
	return s.store.MerchantAccount.Query().Where(merchantaccount.UsernameEQ(username), merchantaccount.DeletedAtIsNil()).Only(ctx)
}
