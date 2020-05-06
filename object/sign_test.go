package object

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestSignVerifyRequests(t *testing.T) {
	sk := test.DecodeKey(0)

	type sigType interface {
		service.SignedDataWithToken
		service.SignKeyPairAccumulator
		service.SignKeyPairSource
		SetToken(*Token)
	}

	items := []struct {
		constructor func() sigType
		bodyCorrupt []func(sigType)
	}{
		{ // PutRequest.PutHeader
			constructor: func() sigType {
				return MakePutRequestHeader(new(Object))
			},
			bodyCorrupt: []func(sigType){
				func(s sigType) {
					obj := s.(*PutRequest).GetR().(*PutRequest_Header).Header.GetObject()
					obj.SystemHeader.PayloadLength++
				},
			},
		},
		{ // PutRequest.Chunk
			constructor: func() sigType {
				return MakePutRequestChunk(make([]byte, 10))
			},
			bodyCorrupt: []func(sigType){
				func(s sigType) {
					h := s.(*PutRequest).GetR().(*PutRequest_Chunk)
					h.Chunk[0]++
				},
			},
		},
	}

	for _, item := range items {
		{ // token corruptions
			v := item.constructor()

			token := new(Token)
			v.SetToken(token)

			require.NoError(t, service.SignDataWithSessionToken(sk, v))

			require.NoError(t, service.VerifyAccumulatedSignaturesWithToken(v))

			token.SetSessionKey(append(token.GetSessionKey(), 1))

			require.Error(t, service.VerifyAccumulatedSignaturesWithToken(v))
		}

		{ // body corruptions
			for _, corruption := range item.bodyCorrupt {
				v := item.constructor()

				require.NoError(t, service.SignDataWithSessionToken(sk, v))

				require.NoError(t, service.VerifyAccumulatedSignaturesWithToken(v))

				corruption(v)

				require.Error(t, service.VerifyAccumulatedSignaturesWithToken(v))
			}
		}
	}
}
