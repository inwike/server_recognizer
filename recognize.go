package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"

	"gocv.io/x/gocv"
)

var EMO = []string{"сердитый", "отвращение", "страх", "счастливый", "нейтральный", "грустный", "сюрприз"}

var emoProto = "./opencv/emo_deploy.prototxt"
var emoModel = "./opencv/EmotiW_VGG_S.caffemodel"

var input chan []byte
var emoNet gocv.Net
var reader *bytes.Reader

func initCV() {

	emoNet = gocv.ReadNet(emoModel, emoProto)
	emoNet.SetPreferableBackend(gocv.NetBackendDefault)
	emoNet.SetPreferableTarget(gocv.NetTargetCPU)

}

func definition(src []byte) string {
	pt1 := image.Pt(224, 224)
	scalar := gocv.NewScalar(0, 0, 0, 0)

	reader = bytes.NewReader(src)
	tmp, err := jpeg.Decode(reader)
	reader = nil

	if err != nil {
		log.Println(err)
		return ""
	}

	img, err := gocv.ImageToMatRGB(tmp)
	if err != nil {
		return ""
	}

	blob := gocv.BlobFromImage(img, 1, pt1, scalar, false, false)
	emoNet.SetInput(blob, "")
	emoPreds := emoNet.Forward("")
	_, _, _, emoLoc := gocv.MinMaxLoc(emoPreds)
	blob.Close()
	emoPreds.Close()
	img.Close()

	tmp = nil

	return EMO[emoLoc.X]
}
