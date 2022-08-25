package bandit

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMultiArmBandit(t *testing.T) {
	t.Run("каждая ручка должна быть дернута хотя бы один раз", func(t *testing.T) {
		rand.Seed(time.Now().Unix())

		count := 10
		arms := make(Arms, 0, count)
		for i := 0; i < count; i++ {
			arms = append(arms, Arm{0, uint(rand.Intn(100))})
		}
		expectedResult := make([]uint, 0, count)
		for i := 0; i < count; i++ {
			expectedResult = append(expectedResult, uint(i))
		}

		result := make([]uint, 0, count)
		for i := 0; i < count; i++ {
			result = append(result, MultiArmBandit(arms))
			arms[i].Count = 1
		}

		require.Equal(t, expectedResult, result)
	})

	t.Run("работа механизма многорукого бандита", func(t *testing.T) {
		visits := 2000

		arms := Arms{
			Arm{0, 1},
			Arm{0, 1},
			Arm{0, 1},
		}

		hits := make(map[uint]int)
		for i := 0; i < visits; i++ {
			armIdx := MultiArmBandit(arms)

			hits[armIdx]++
			arms[armIdx].Count++
			switch armIdx {
			case 0:
				arms[0].Reward += 3
			case 1:
				arms[1].Reward += 2
			default:
				arms[armIdx].Reward++
			}
		}

		require.GreaterOrEqual(t, float64(hits[0])/float64(visits)*100, float64(70), "кол-во просмотров должно быть больше 85%")
		require.GreaterOrEqual(t, float64(hits[0])/float64(visits)*100, float64(10), "кол-во просмотров должно быть меньше 10%")
	})
}
