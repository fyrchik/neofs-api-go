package signature

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/stretchr/testify/require"
)

func TestBalanceResponse(t *testing.T) {
	dec := new(accounting.Decimal)
	dec.SetValue(100)

	body := new(accounting.BalanceResponseBody)
	body.SetBalance(dec)

	meta := new(service.ResponseMetaHeader)
	meta.SetTTL(1)

	req := new(accounting.BalanceResponse)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	// verify unsigned request
	require.Error(t, VerifyServiceMessage(req))

	key, err := crypto.LoadPrivateKey("Kwk6k2eC3L3QuPvD8aiaNyoSXgQ2YL1bwS5CP1oKoA9waeAze97s")
	require.NoError(t, err)

	// sign request
	require.NoError(t, SignServiceMessage(key, req))

	// verification must pass
	require.NoError(t, VerifyServiceMessage(req))

	// add level to meta header matryoshka
	meta = new(service.ResponseMetaHeader)
	meta.SetOrigin(req.GetMetaHeader())
	req.SetMetaHeader(meta)

	// sign request
	require.NoError(t, SignServiceMessage(key, req))

	// verification must pass
	require.NoError(t, VerifyServiceMessage(req))

	// corrupt body
	dec.SetValue(dec.GetValue() + 1)

	// verification must fail
	require.Error(t, VerifyServiceMessage(req))

	// restore body
	dec.SetValue(dec.GetValue() - 1)

	// corrupt meta header
	meta.SetTTL(meta.GetTTL() + 1)

	// verification must fail
	require.Error(t, VerifyServiceMessage(req))

	// restore meta header
	meta.SetTTL(meta.GetTTL() - 1)

	// corrupt origin verification header
	req.GetVerificationHeader().SetOrigin(nil)

	// verification must fail
	require.Error(t, VerifyServiceMessage(req))
}