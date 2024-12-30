// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"bytes"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_FundManagerAddRewardPool(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("FundManagerAddRewardPool"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(FundManagerAddRewardPool)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(FundManagerAddRewardPool)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}
