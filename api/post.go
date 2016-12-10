package api

import (
	"bytes"
	"bufio"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"github.com/nfnt/resize"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"

	"github.com/microservices-today/aws-sushi/conf"
	"github.com/microservices-today/aws-sushi/json"
)

type iconf struct {
	image                                                image.Image
	machine, format, density, ui, hash, color, fid, path string
	width, height                                        uint
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  // allow cross domain AJAX requests
	if r.Method != "POST" {
		w.Write(json.Message("ERROR", "Not supported Method"))
		return
	}
	f, _, err := r.FormFile("image")
	if err != nil {
		w.Write(json.Message("ERROR", "Can't Find Image"))
		return
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	h := sha1.New()
	var result json.Result
	var ic iconf
	ic.machine = conf.Image.Machine
	ic.image, _, err = image.Decode(io.TeeReader(buf, h))
	ic.hash = fmt.Sprintf("%x", h.Sum(nil))
	if err != nil {
		w.Write(json.Message("ERROR", "Unable to decode your image! Type=" + conf.InputType + " error:" + err.Error()))
		return
	}
	setColor(&ic)
	for _, format := range conf.Image.Format {
		for _, screen := range conf.Image.Screen {
			ic.format = format
			ic.ui = screen.Ui
			ic.density = screen.Density
			ic.width = screen.Width
			if ic.fid, err = imgToBuf(&ic); err != nil {
				w.Write(json.Message("ERROR", "Unable to run imgToBuf"))
				return
			}
			fid := json.Fid{fmt.Sprintf("%s_%s", screen.Density, screen.Ui), ic.fid}
			result.Image = append(result.Image, fid)
		}
	}
	w.Write(json.Message("OK", &result))
}

func imgToBuf(ic *iconf) (string, error) {
	img := resize.Resize(ic.width, 0, ic.image, resize.NearestNeighbor)
	ic.height = uint(img.Bounds().Size().Y)
	ic.fid = fmt.Sprintf("%s-%s-%s-%s-%s-%s-%d-%d", ic.machine, ic.format, ic.density, ic.ui, ic.hash, ic.color, ic.width, ic.height)
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: conf.Quality})
	if err != nil {
		fmt.Println("ERROR: Unable to Encode into jpeg")
		return "", err
	}
	if err = putToS3(ic.fid, buf.Bytes()); err != nil {
		fmt.Println("ERROR: Unable to s3.Put. err = " + err.Error())
		return "", err
	}
	return ic.fid, err
}

func setColor(ic *iconf) {
	img1x1 := resize.Resize(1, 1, ic.image, resize.NearestNeighbor)
	red, green, blue, _ := img1x1.At(0, 0).RGBA()
	ic.color = fmt.Sprintf("%X%X%X", red >> 8, green >> 8, blue >> 8) // removing 1 byte 9A16->9A
}

func putToS3(path string, data []byte) error {
	sushiAuth := aws.Auth{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
	}
	saeast := aws.APSoutheast
	connection := s3.New(sushiAuth, saeast)
	bucket := connection.Bucket(conf.S3Bucket)
	err := bucket.Put(path, data, conf.Mime, s3.ACL("public-read"), s3.Options{})
	if err != nil {
		fmt.Println("ERROR: Unable to s3.Put. err = " + err.Error())
		return err
	}
	return err
}