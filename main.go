package keystone_effect_live_correction

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"strconv"
)

func main() {
	// parse args
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\t keystone-live-correction [USB Port ID]")
		return
	}
	deviceID, _ := strconv.Atoi(os.Args[1])


	// open video feed from USB port
	USBFeed, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer USBFeed.Close()


	// open display fullscreen window
	window := gocv.NewWindow("Face Detect")
	window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
	defer window.Close()


	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("start reading from usb port: %v\n", deviceID)
	for {
		if ok := USBFeed.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}
