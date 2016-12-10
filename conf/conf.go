package conf

import (
	"encoding/json"
	"log"
	"os"
)

// Sushi
const (
	InputType = "jpeg"
	SushiId = 1
	SushiStore = "var/sushi/store"
	SushiPort = 9090
	CacheMaxAge = 30 * 24 * 60 * 60 // 30 days
	Mime = "image/jpeg"
	Lossless = true
	Quality = 95
)

// Aws
var (
	AccessKey = ""
	SecretKey = ""
	S3Bucket = ""
)

var Image ImageConf

type ImageConf struct {
	Machine string    `json:"machine"`
	Format  []string  `json:"format"`
	Hash    string    `json:"hash"`
	Color   string    `json:"color"`
	Screen  []Density `json:"screen"`
}

type Density struct {
	Density string `json:"density"`
	Ui      string `json:"ui"`
	Width   uint   `json:"width"`
	Height  uint   `json:"height"`
}

func initImageConf() {
	file, _ := os.Open("/etc/sushi/sushi.conf")
	decoder := json.NewDecoder(file)
	Image = ImageConf{}
	err := decoder.Decode(&Image)
	if err != nil {
		log.Fatal("Was not able to decode conf file: ", err)
	}
}

func initEnvron() {
	AccessKey = os.Getenv("AWS_ACCESS_KEY")
	if len(AccessKey) == 0 {
		log.Fatal("Set AWS_ACCESS_KEY (e.g. export AWS_ACCESS_KEY=XXX)")
	}
	SecretKey = os.Getenv("AWS_SECRET_KEY")
	if len(SecretKey) == 0 {
		log.Fatal("Set AWS_SECRET_KEY (e.g. export AWS_SECRET_KEY=YYY)")
	}
	S3Bucket = os.Getenv("S3_BUCKET")
	if len(S3Bucket) == 0 {
		log.Fatal("Set S3_BUCKET (e.g. export S3_BUCKET=example.sushi.com)")
	}
}

func init() {
	initEnvron()
	initImageConf()
}