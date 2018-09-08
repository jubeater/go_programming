package main

import (
	"log"
	"net/http"
	"image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "strconv"
    "strings"
)
var palette = []color.Color{color.White, color.RGBA{0x00,0xff,0x00,0xff}}

const (
    whiteIndex = 0 // first color in palette
    greenIndex = 1 // next color in palette
)
func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if strings.Compare(k, "cycles") == 0 {
				f, err := strconv.ParseFloat(v[0], 64) 
				if err != nil{
					log.Print(err)
				}
				lissajous(w, f)
			}
		}
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
// handler echoes the Path component of the request URL r.
   func lissajous(out io.Writer, in float64) {
       const (
           res     = 0.001 // angular resolution
           size    = 100   // image canvas covers [-size..+size]
           nframes = 64    // number of animation frames
           delay   = 8     // delay between frames in 10ms units
       )
       var cycles float64     // number of complete x oscillator revolutions
       cycles = in
       freq := rand.Float64() * 3.0 // relative frequency of y oscillator
       anim := gif.GIF{LoopCount: nframes}
       phase := 0.0 // phase difference
       for i := 0; i < nframes; i++ {
           rect := image.Rect(0, 0, 2*size+1, 2*size+1)
           img := image.NewPaletted(rect, palette)
           for t := 0.0; t < cycles*2*math.Pi; t += res {
               x := math.Sin(t)
               y := math.Sin(t*freq + phase)
               img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                   greenIndex)
           }
           phase += 0.1
           anim.Delay = append(anim.Delay, delay)
           anim.Image = append(anim.Image, img)
       }
       gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
   }