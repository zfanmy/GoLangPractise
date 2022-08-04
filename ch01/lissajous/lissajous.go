package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var paletee = []color.Color{
	//color.White,
	//color.Black,
	//color.RGBA{255, 255, 255, 255}, // Red, Green, Blue, Alpha (Alpha channel is how opaque each pixel is) Alpha 为不透明度
	color.RGBA{0, 0, 0, 255}, // 练习1.5 使用RGBA颜色
	color.RGBA{0, 255, 0, 255},
}

const (
	//whiteIndex = 0
	blackIndex = 0
	greenIndex = 1
)

// Paint 绘图
func Paint() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 //角度分辨率
		size    = 100   // 画像画布包含[-size, size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 5.0 // 练习1.6 改变参数获取纺锤图
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletee)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
