package card

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
)

func TestSetCardStatus(t *testing.T) {
	tests := []struct {
		input protos.SetCardStatusRequest
		want  uint64
	}{
		{
			input: protos.SetCardStatusRequest{
				CardId: 1,
				Status: protos.CardStatus_Unused,
			},
			want: 1,
		},
	}

	var cardId uint64
	for _, test := range tests {
		cardId, _ = SetCardStatus(&test.input)
		assert.Equal(t, test.want, cardId, test)
	}
}
