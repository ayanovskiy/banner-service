package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("парсинг конфига", func(t *testing.T) {
		actualConfig, err := NewConfig("./../../test/testdata/config/config_test.yaml")

		assert.IsType(t, &Config{}, actualConfig)
		assert.Nil(t, err)

		assert.Equal(t, "sql://localhost:1221", actualConfig.Database.Dsn)
		assert.Equal(t, "localhost", actualConfig.Server.Addr)
		assert.Equal(t, uint(8080), actualConfig.Server.Port)
	})

	t.Run("проверка на расширение файла", func(t *testing.T) {
		_, err := NewConfig("./../../test/testdata/config/config_test.txt")

		assert.Equal(t, ErrUnsupportedType, err)
	})
}
