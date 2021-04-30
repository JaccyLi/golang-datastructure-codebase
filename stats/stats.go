package stats

import (
	"fmt"
	"suosuoli-golangds/array"
	"time"

	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
)

func StartStats() {
	viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeWesteros), viewer.WithAddr("localhost:8087"))
	mgr := statsview.New()

	// Start() runs a HTTP server at `localhost:18066` by default.
	go mgr.Start()

	// Stop() will shutdown the http server gracefully
	// mgr.Stop()

	// busy working....
	newarr := array.NewArrayList()
	for i := 1; i < 400; i++ {
		time.Sleep(time.Millisecond * 100)
		newarr.AppendElement(20 + i)
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("size", newarr.GetSize())
	for i := 0; i < 400; i++ {
		time.Sleep(time.Millisecond * 100)
		newarr.DeleteElement(int64(i))
	}
	fmt.Println("size", newarr.GetSize())

}
