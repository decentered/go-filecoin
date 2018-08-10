package api

import (
	"context"

	cid "gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"

	"github.com/filecoin-project/go-filecoin/actor/builtin/paymentbroker"
	"github.com/filecoin-project/go-filecoin/types"
)

type Paych interface {
	Create(ctx context.Context, fromAddr, target types.Address, eol *types.BlockHeight, amount *types.AttoFIL) (*cid.Cid, error)
	Ls(ctx context.Context, fromAddr, payerAddr types.Address) (map[string]*paymentbroker.PaymentChannel, error)
	Voucher(ctx context.Context, fromAddr types.Address, channel *types.ChannelID, amount *types.AttoFIL) (string, error)
	Redeem(ctx context.Context, fromAddr types.Address, voucherRaw string) (*cid.Cid, error)
	Reclaim(ctx context.Context, fromAddr types.Address, channel *types.ChannelID) (*cid.Cid, error)
	Close(ctx context.Context, fromAddr types.Address, voucherRaw string) (*cid.Cid, error)
	Extend(ctx context.Context, fromAddr types.Address, channel *types.ChannelID, eol *types.BlockHeight, amount *types.AttoFIL) (*cid.Cid, error)
}