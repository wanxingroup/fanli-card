package card

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
)

func TestCreateCard(t *testing.T) {
	tests := []struct {
		input protos.CreateCardRequest
		want  uint64
	}{
		{
			input: protos.CreateCardRequest{
				Name:        "name",
				Description: "description",
				Sort:        1,
			},
			want: 1,
		},
	}

	var cardId uint64
	for _, test := range tests {
		cardId, _ = createCard(&test.input)
		assert.Greater(t, cardId, test.want, test)
	}
}
