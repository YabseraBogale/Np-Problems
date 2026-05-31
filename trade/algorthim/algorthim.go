package algorthim

type Tick struct {
	Volume float64
	Price  float64
}

func VWAP(t []Tick) float64 {
	var sum_volume float64
	var sum float64
	for _, i := range t {
		sum += i.Price * i.Volume
		sum_volume += i.Volume
	}
	return sum / sum_volume
}
