package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) (count uint64) {
	count = k.GetPostCount(ctx)
	post.Id = count
	byteKey := []byte(types.PostKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, post.Id)
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(byteKey, appendedValue)
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.PostCountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.PostCountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)

}
