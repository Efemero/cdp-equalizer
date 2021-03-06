package cdp

import (
	"github.com/efemero/cdp-equalizer/ethutils"

	"fmt"
	"math/big"
)

// Status represents the status of a CDP at a given price
type Status struct {
	ID      int64
	DaiDebt *Float
	EthCol  *Float
	Price   *Float
	Ratio   *Float
	DaiNet  *Float
	EthNet  *Float
}

// Float is for json purposes
type Float big.Float

// MarshalJSON transforms the Float in JSON number
func (f *Float) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", (*big.Float)(f))
	return []byte(s), nil
}

func (f *Float) String() string {
	f64, _ := (*big.Float)(f).Float64()
	return fmt.Sprintf("%s", ethutils.RenderFloat("", f64))
}

func (s *Status) String() string {
	return fmt.Sprintf("%s: %s (%s)", s.Price, s.EthNet, s.DaiNet)
}
