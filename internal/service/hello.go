package service

import (
	"context"
	"fmt"
	"github.com/sysatom/framework/ent"
	"time"
)

type HelloService struct {
	store *ent.Client
}

func NewHelloService(store *ent.Client) *HelloService {
	return &HelloService{
		store: store,
	}
}

func (s *HelloService) Hello(ctx context.Context) error {
	a1, err := s.store.MerchantAccount.Create().
		SetUsername(time.Now().String()).
		SetPassword("123456").
		SetEmail("j@j.com").
		SetIsMainAccount(true).
		SetPhone("1234567890").
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	m1, err := s.store.Merchant.Create().
		SetMerchantName(time.Now().String()).
		SetCity("joey").
		SetAddress("joey").
		SetContactPerson("joey").
		SetContactPhone("joey").
		SetCountry("joey").
		SetProvince("joey").
		SetDistrict("joey").
		AddAccounts(a1).
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(m1)

	list, err := s.store.Merchant.Query().QueryAccounts().All(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(list)

	return nil
}
