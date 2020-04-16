package prepared

import (
	"bytes"
	"testing"

	"github.com/modern-go/parse"
	"github.com/stretchr/testify/require"
)

func TestDecodePreparedQuery(t *testing.T) {
	var testCase = []struct {
		raw       []byte
		expect    *QueryBody
		shouldErr bool
	}{
		{
			raw: []byte{
				0x3c, 0x00, 0x00, 0x00, 0x16, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20,
				0x2a, 0x20, 0x46, 0x52, 0x4f, 0x4d, 0x20, 0x64, 0x65, 0x70, 0x61, 0x72,
				0x74, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x57, 0x48, 0x45, 0x52, 0x45, 0x20,
				0x28, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x3f, 0x20, 0x41, 0x4e, 0x44, 0x20,
				0x61, 0x67, 0x65, 0x3e, 0x3f, 0x20, 0x41, 0x4e, 0x44, 0x20, 0x61, 0x67,
				0x65, 0x3c, 0x3f, 0x29,
			},
			expect: &QueryBody{
				RawQuery: "SELECT * FROM department WHERE (name=? AND age>? AND age<?)",
			},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		src, err := parse.NewSource(bytes.NewReader(tc.raw), 10)
		should.NoError(err)
		actual, err := DecodePreparedQuery(src)
		if tc.shouldErr {
			should.Error(err, "case #%d fail", idx)
			should.Nil(actual, "case #%d fail", idx)
		} else {
			should.NoError(err, "case #%d fail", idx)
			should.Equal(tc.expect, actual, "case #%d fail", idx)
		}
	}
}