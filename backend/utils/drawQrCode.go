package utils

import (
	"bytes"
	"image/png"

	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
)

func DrawQrCode(url, name string) (*bytes.Buffer, error) {
	qrCodeData, err := qrcode.Encode(url+name, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	qrCodeImg, err := png.Decode(bytes.NewReader(qrCodeData))
	if err != nil {
		return nil, err
	}

	const imgW, imgH = 256, 360
	dc := gg.NewContext(imgW, imgH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace("./fonts/Hack Regular Nerd Font Complete.ttf", 16); err != nil {
		return nil, err
	}

	var y float64
	sw, _ := dc.MeasureString(name)
	x := (imgW - sw) / 2
	y = 40
	dc.DrawStringAnchored(name, x, y, 0, 0)

	scanMeText := "Scan Me"
	sw, _ = dc.MeasureString(scanMeText)
	x = (imgW - sw) / 2
	y = 100
	dc.DrawStringAnchored(scanMeText, x, y, 0, 0)

	dc.DrawImage(qrCodeImg, 0, 104)

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, dc.Image())
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
