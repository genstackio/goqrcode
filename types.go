package goqrcode

type Config struct {
	Data            string  `json:"data"`
	Scale           float64 `json:"scale"`
	BlockSize       int     `json:"blockSize"`
	BlockStyle      string  `json:"blockStyle"`
	EmptyBlockStyle string  `json:"emptyBlockStyle"`
	Offset          int     `json:"offset"`
}
