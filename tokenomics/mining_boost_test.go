// SPDX-License-Identifier: ice License 1.0

package tokenomics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMiningBoostPaymentAddress(t *testing.T) {
	assert.Equal(t, "0x000000000000000000000000000000000000dead", generateMiningBoostPaymentAddress(0))
	assert.Equal(t, "0x000000000000000000000000000000000001dead", generateMiningBoostPaymentAddress(1))
	assert.Equal(t, "0x000000000000000000000000000000000002dead", generateMiningBoostPaymentAddress(2))
	assert.Equal(t, "0x000000000000000000000000000000000010dead", generateMiningBoostPaymentAddress(10))
	assert.Equal(t, "0x000000000000000000000000000001234567dead", generateMiningBoostPaymentAddress(1234567))
	assert.Equal(t, "0x000000000000000000000000000010000000dead", generateMiningBoostPaymentAddress(10_000_000))
	assert.Equal(t, "0x000000000000000000000000000011111111dead", generateMiningBoostPaymentAddress(11_111_111))
}
