package integration

import (
	"banner-service/internal/rmq"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRabbit(t *testing.T) {
	t.Run("тест отправки сообщения", func(t *testing.T) {
		ctx := context.Background()

		r, err := rmq.NewRabbit(
			ctx,
			"amqp://guest:guest@127.0.0.1:5672",
			"banner-test",
			"banner-test",
			"banners-stat-test",
		)
		if err != nil {
			assert.Fail(t, err.Error())
		}

		err = r.SendStat(rmq.StatMessage{BannerId: 1, SlotId: 1, GroupId: 1, Shows: 10, Hits: 5})
		assert.Nil(t, err)
	})
}
