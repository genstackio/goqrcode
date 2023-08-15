package goqrcode

import (
	"github.com/aaronarduino/goqrsvg"
	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"io"
	"math"
)

//goland:noinspection GoUnusedExportedFunction
func GenerateAndStreamQrCode(w io.Writer, config Config) {
	qrCode, _ := qr.Encode(config.Data, qr.H, qr.Auto)
	s := svg.New(w)
	blockSize := 4.0
	if math.Abs(config.Scale) > 0.00001 {
		blockSize *= config.Scale
	}
	qs := goqrsvg.NewQrSVG(qrCode, int(blockSize))
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)
	s.End()
}
