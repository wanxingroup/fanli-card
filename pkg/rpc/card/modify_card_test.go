package card

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
)

func TestModifyCard(t *testing.T) {
	tests := []struct {
		input protos.ModifyCardRequest
		want  uint64
	}{
		{
			input: protos.ModifyCardRequest{
				Card: &protos.ModifyCardInformation{
					CardId:          1,
					Name:            "aa",
					Description:     "aa",
					BackgroundImage: "aa",
					Sort:            0,
				},
			},
			want: 1,
		},
	}

	var cardId uint64
	for _, test := range tests {
		cardId, _ = ModifyCard(&test.input)
		assert.Equal(t, test.want, cardId, test)
	}
}
