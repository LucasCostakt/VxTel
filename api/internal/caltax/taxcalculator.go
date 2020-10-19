package caltax

import (
	"fmt"

	store "github.com/Vortx-Test/api/internal/store"
)

//Calcula a taxas que serão pagas com e sem o plano fale mais
func CalculeteTax(time int, cost float64, plano string, c *store.ConvertNumbers) error {
	switch plano {
	case "30":

		if time-30 > 0 {
			t := time - 30
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}

	case "60":
		if time-60 > 0 {
			t := time - 60
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}
	case "120":
		if time-120 > 0 {
			t := time - 120
			desc := float64(t) * cost
			c.Value = fmt.Sprintf("%.2f", desc*0.10+desc)
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		} else {
			c.Value = "0"
			c.ValueFull = fmt.Sprintf("%.2f", float64(time)*cost)
			return nil
		}
	default:
		return fmt.Errorf("erro")
	}

	return fmt.Errorf("erro")
}
