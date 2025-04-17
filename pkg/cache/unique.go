package cache

import (
	"context"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/sysatom/framework/pkg/flog"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/utils"
)

func Unique(ctx context.Context, id string, latest []any) ([]types.KV, error) {
	result := make([]types.KV, 0)
	uniqueKey := fmt.Sprintf("unique:%s", id)

	for _, item := range latest {
		val, err := kvHash(item)
		if err != nil {
			return nil, fmt.Errorf("failed to hash kv: %w", err)
		}
		if len(val) == 0 {
			continue
		}
		b, err := DB.SAdd(ctx, uniqueKey, val).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to set unique key: %w", err)
		}
		if b == 1 {
			kv, ok := item.(map[string]any)
			if !ok {
				continue
			}
			result = append(result, kv)
			flog.Info("[unique] key: %s added: %s", id, val)
		}
	}

	return result, nil
}

func kvHash(item any) (string, error) {
	b, err := json.ConfigCompatibleWithStandardLibrary.Marshal(item)
	if err != nil {
		return "", fmt.Errorf("failed to marshal kv: %w", err)
	}
	return utils.SHA1(utils.BytesToString(b)), nil
}

func UniqueString(ctx context.Context, id string, latest string) (bool, error) {
	uniqueKey := fmt.Sprintf("unique:%s", id)
	b, err := DB.SAdd(ctx, uniqueKey, latest).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set unique key: %w", err)
	}
	if b == 1 {
		return true, nil
	}

	return false, nil
}
