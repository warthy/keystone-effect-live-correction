package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

func main() {
	//// parse args
	//if len(os.Args) < 2 {
	//	fmt.Println("How to run:\n\t keystone-live-correction [USB Port ID]")
	//	return
	//}
	deviceID := 1 //strconv.Atoi(os.Args[1])
	//
	//// open video feed from USB port
	//USBFeed, err := gocv.VideoCaptureDevice(deviceID)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer USBFeed.Close()

	// open display fullscreen window
	window := gocv.NewWindow("keystone correction")
	//window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
	defer window.Close()

	// prepare image matrix
	buf := gocv.IMRead("wallpaper.jpeg", gocv.IMReadUnchanged)
	defer buf.Close()

	fmt.Printf("start reading from usb port: %v\n", deviceID)
	for i := 0; i < 100; i++ {
		//if ok := USBFeed.Read(&buf); !ok {
		//	fmt.Printf("cannot read device %d\n", deviceID)
		//	return
		//}
		if buf.Empty() {
			continue
		}

		srcPoints := gocv.NewPointVectorFromPoints([]image.Point{
			{0, 0},                   // top left corner
			{buf.Cols(), 0},          // top right corner
			{0, buf.Rows()},          // bottom left corner
			{buf.Cols(), buf.Rows()}, // bottom right corner
		})
		dstPoints := gocv.NewPointVectorFromPoints([]image.Point{
			{100, 100},               // top left corner
			{buf.Cols(), 0},          // top right corner
			{100, buf.Rows() - 100},  // bottom left corner
			{buf.Cols(), buf.Rows()}, // bottom right corner
		})
		h := gocv.GetPerspectiveTransform(srcPoints, dstPoints)

		imgWrapped := gocv.NewMat()
		gocv.WarpPerspective(buf, &imgWrapped, h, image.Point{X: buf.Cols(), Y: buf.Rows()})

		window.IMShow(imgWrapped)
		window.WaitKey(1)
	}
}
