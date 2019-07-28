package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Unit struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string
	Format     string
	Units      []string
	Quantities []int
}

func AdjustUnitQty(unit Unit, qty int) (string, float64) {
	var closest_idx int
	for idx, elem := range unit.Quantities {
		if elem <= qty {
			closest_idx = idx
		} else {
			break
		}
	}
	return unit.Units[closest_idx], float64(qty) / float64(unit.Quantities[closest_idx])
}

func (unit Unit) DisplayUnit(obj string, qty int) string {
	str := unit.Format
	unit_name, newqty := AdjustUnitQty(unit, qty)
	str = strings.Replace(str, "$OBJ", obj, 1)
	str = strings.Replace(str, "$QTY", fmt.Sprintf("%.3g", newqty), 1)
	str = strings.Replace(str, "$UNIT", unit_name, 1)

	return str
}
