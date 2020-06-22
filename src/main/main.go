package main

import (
	. "./lib"
	. "./lib/imageio"
	. "./lib/morphing"
	"strconv"
)

//func avg(ts []int64) int64 {
//	total := 0.0
//	n := len(ts)
//	for _, t := range ts {
//		total += float64(t) / float64(n)
//	}
//
//	return int64(math.Round(total))
//}
//
//type testFunc func()

func main() {
	img1, _ := Load("../resources/dog.jpg")
	img2, _ := Load("../resources/parrot.jpg")

	krom := Parallel(16, 1000)

	imgs := krom.Morph(img1, img2, []Vertex{
		{X: 1050, Y: 364},
		{X: 132, Y: 580},
		{X: 380, Y: 80},
		{X: 600, Y: 1020},
		{X: 650, Y: 250},
		{X: 720, Y: 600},
		{X: 800, Y: 200},
		{X: 820, Y: 1050},
		{X: 570, Y: 70},
	}, []Vertex{
		{X: 892, Y: 732},
		{X: 156, Y: 890},
		{X: 270, Y: 480},
		{X: 600, Y: 1030},
		{X: 600, Y: 500},
		{X: 720, Y: 600},
		{X: 840, Y: 440},
		{X: 805, Y: 1090},
		{X: 580, Y: 340},
	}, 12)

	for idx, img := range imgs {
		Save(img, "../resources/result"+strconv.Itoa(idx), "png")
	}

	delays := make([]int, len(imgs))
	for i := range delays {
		delays[i] = 20
	}
	result, _ := krom.Convert().ToGif(imgs, delays, 256)

	Save(result, "../resources/result", "gif")
	//img, err := Load("../resources/boat.png")
	//if err != nil {
	//	panic(err)
	//}
	//
	//numRepeats := 100
	//
	//workers := make([]string, 0)
	//
	//testName := []string{"ColorMapping", "Resize2x", "GaussianBlur", "Rotate180"}
	//libName := []string{"kromatique", "bild", "imaging"}
	//
	//var tests []testFunc
	//times := make([][][]string, 4)
	//for i := range times {
	//	times[i] = make([][]string, 3)
	//}
	//
	//maxProcs := runtime.GOMAXPROCS(0)
	//
	//var krom *Krom
	//
	//for numWorkers := 1; numWorkers <= maxProcs; numWorkers *= 2 {
	//	fmt.Println(">> >> Started for", numWorkers, "workers")
	//
	//	runtime.GOMAXPROCS(numWorkers)
	//	workers = append(workers, strconv.Itoa(numWorkers))
	//
	//	for libTested := 0; libTested < 3; libTested++ {
	//		fmt.Println(">> Started for", libName[libTested])
	//
	//		sigma := 0.84
	//		rotationAngle := math.Pi
	//
	//		switch libTested {
	//		case 0:
	//			krom = Parallel(numWorkers, 1000)
	//			mapper := krom.Effect().ColorMapper(strategy.Negative)
	//			scale := krom.Effect().Scale(strategy.CornerPixelsSampling)
	//			filter := krom.Effect().Filter(strategy.Mirror, strategy.GaussianBlurKernel(sigma))
	//			dist := krom.Effect().Distortion(
	//				strategy.None,
	//				strategy.NewRotationLens(
	//					geometry.Pt2D(
	//						float64(img.Bounds().Dx()/2),
	//						float64(img.Bounds().Dy()/2)),
	//					rotationAngle))
	//
	//			tests = []testFunc{
	//				func() {
	//					mapper.Apply(img).Result()
	//				},
	//				func() {
	//					scale.Apply(img, img.Bounds().Dx()*2, img.Bounds().Dy()*2).Result()
	//				},
	//				func() {
	//					filter.Apply(img).Result()
	//				},
	//				func() {
	//					dist.Apply(img).Result()
	//				},
	//			}
	//			break
	//		case 1:
	//			tests = []testFunc{
	//				func() {
	//					bildEffect.Invert(img)
	//				},
	//				func() {
	//					bildTransform.Resize(img, img.Bounds().Dx()*2, img.Bounds().Dy()*2, bildTransform.Linear)
	//				},
	//				func() {
	//					bildBlur.Gaussian(img, 3)
	//				},
	//				func() {
	//					bildTransform.Rotate(img, rotationAngle, nil)
	//				},
	//			}
	//			break
	//
	//		case 2:
	//			tests = []testFunc{
	//				func() {
	//					imaging.Invert(img)
	//				},
	//				func() {
	//					imaging.Resize(img, img.Bounds().Dx()*2, img.Bounds().Dy()*2, imaging.Linear)
	//				},
	//				func() {
	//					imaging.Blur(img, sigma)
	//				},
	//				func() {
	//					imaging.Rotate(img, rotationAngle, color.Transparent)
	//				},
	//			}
	//			break
	//		}
	//
	//
	//		for k := 0; k < 4; k++ {
	//			fmt.Print("Started", testName[k])
	//
	//			timesRun := make([]int64, numRepeats)
	//			for j := 0; j < numRepeats; j++ {
	//				start := time.Now()
	//
	//				tests[k]()
	//
	//				end := time.Now()
	//				t := end.Sub(start).Microseconds()
	//				timesRun[j] = t
	//			}
	//
	//			avgTimesRun := avg(timesRun)
	//			times[k][libTested] = append(times[k][libTested], strconv.Itoa(int(avgTimesRun)))
	//
	//			fmt.Println(" =", avgTimesRun, "μs")
	//		}
	//
	//		if libTested == 0 {
	//			krom.Stop()
	//		}
	//	}
	//
	//	for k := 0; k < 4; k++ {
	//		filename := "../resources/benchmarks/" + testName[k] + ".csv"
	//		file, err := os.Create(filename)
	//		if err != nil {
	//			log.Fatal("Cannot create file", err)
	//		}
	//
	//		writer := csv.NewWriter(file)
	//
	//		if err := writer.Write([]string{"workers", "kromatique", "bild", "imaging"}); err != nil {
	//			log.Fatal("Cannot write to file", err)
	//		}
	//
	//		for i := range workers {
	//			if err := writer.Write([]string{workers[i], times[k][0][i], times[k][1][i], times[k][2][i]}); err != nil {
	//				log.Fatal("Cannot write to file", err)
	//			}
	//		}
	//
	//		writer.Flush()
	//		if err := file.Close(); err != nil {
	//			log.Fatal("Cannot close file", err)
	//		}
	//	}
	//}
}
