package leetcode

import (
	"testing"
)

/**
 * 1603. 设计停车系统
 */

func Test_leetcode_1603(t *testing.T) {
	parkingSystem := Constructor1603(1, 1, 0)
	parkingSystem.AddCar(1) // 返回 true ，因为有 1 个空的大车位
	parkingSystem.AddCar(2) // 返回 true ，因为有 1 个空的中车位
	parkingSystem.AddCar(3) // 返回 false ，因为没有空的小车位
	parkingSystem.AddCar(1) // 返回 false ，因为没有空的大车位，唯一一个大车位已经被占据了
}

type ParkingSystem struct {
	big, medium, small          int
	bigMax, mediumMax, smallMax int
}

func Constructor1603(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{0, 0, 0, big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big < this.bigMax {
			this.big++
			return true
		}
	case 2:
		if this.medium < this.mediumMax {
			this.medium++
			return true
		}
	case 3:
		if this.small < this.smallMax {
			this.small++
			return true
		}
	}
	return false
}
