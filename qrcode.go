package goqrcode

import (
	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"image/color"
	"io"
	"math"
)

var DefaultFilledBlockStyle = "fill:black"
var DefaultNotFilledBlockStyle = "fill:white"

//goland:noinspection GoUnusedExportedFunction
func GenerateAndStreamQrCode(w io.Writer, config Config) {
	qrCode, _ := qr.Encode(config.Data, qr.H, qr.Auto)
	s := svg.New(w)
	blockSize := config.BlockSize
	if math.Abs(config.Scale) > 0.00001 {
		blockSize = int(float64(blockSize) * config.Scale)
	}
	width := (qrCode.Bounds().Max.X * blockSize) + (blockSize * (config.Offset * 2))
	firstX := blockSize * config.Offset
	firstY := blockSize * config.Offset

	s.Start(width, width)

	currY := firstY

	filledBlockStyle, notFilledBlockStyle := buildBlockStylesFromConfig(config)

	for x := 0; x < width; x++ {
		currX := firstX
		for y := 0; y < width; y++ {
			switch qrCode.At(x, y) {
			case color.Black:
				s.Rect(currX, currY, blockSize, blockSize, filledBlockStyle)
			case color.White:
				s.Rect(currX, currY, blockSize, blockSize, notFilledBlockStyle)
			}
			currX += blockSize
		}
		currY += blockSize
	}

	s.End()
}

func buildBlockStylesFromConfig(config Config) (string, string) {
	filledBlockStyle := config.BlockStyle
	if len(filledBlockStyle) == 0 {
		filledBlockStyle = DefaultFilledBlockStyle
	}
	notFilledBlockStyle := config.EmptyBlockStyle
	if len(notFilledBlockStyle) == 0 {
		notFilledBlockStyle = DefaultNotFilledBlockStyle
	}

	return filledBlockStyle, notFilledBlockStyle
}
