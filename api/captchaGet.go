package api

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	util "github.com/AVENTER-UG/util"
	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Function to create captcha
func (s *Service) apiV0CaptchaGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png; charset=utf-8")
	w.Header().Set("Api-Service", "v0")

	captcha := randString(8)
	img := image.NewRGBA(image.Rect(0, 0, 140, 50))
	addLabel(img, 20, 30, captcha)

	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		logrus.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	if err := png.Encode(tmpFile, img); err != nil {
		tmpFile.Close()
		logrus.Error(err)
	}

	if err := tmpFile.Close(); err != nil {
		logrus.Error(err)
	}

	content, _ := ioutil.ReadFile(tmpFile.Name())

	// create a session token, save it together with the captcha in the caching server
	sessionToken, err := util.GenUUID()
	if err != nil {
		logrus.Error("Could not create session token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logrus.Debug("Add SessionToken: ", sessionToken, captcha)
	err = s.Cache.Set(sessionToken, captcha, 0).Err()
	if err != nil {
		logrus.Error("Could not store session token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("sessionToken", sessionToken)
	w.Write(content)
}

// add label to temp image file
func addLabel(img *image.RGBA, x, y int, label string) {
	oFont, err := ioutil.ReadFile("font/Inter-Regular.ttf")

	if err != nil {
		logrus.Error("Could not load font: ", err)
		return
	}

	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	col := color.RGBA{200, 100, 0, 255}

	fontTT, _ := truetype.Parse(oFont)

	fnt := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(col),
		//Face: basicfont.Face7x13,
		Face: truetype.NewFace(fontTT, &truetype.Options{
			Size:    16,
			DPI:     92,
			Hinting: font.HintingFull,
		}),
		Dot: point,
	}
	fnt.DrawString(label)
}

// create random string
func randString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
