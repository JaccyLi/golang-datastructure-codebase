package stats

import (
	"fmt"
	"suosuoli-golangds/array"
	"time"

	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
)

// StartStats is
func StartStats() {
	// Visit your browser at http://localhost:18066/debug/statsview
	// Or debug as always via http://localhost:18066/debug/pprof, http://localhost:18066/debug/pprof/heap, ...
	viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeWesteros), viewer.WithAddr("localhost:18066"))
	mgr := statsview.New()

	// Start() runs a HTTP server at `localhost:18066` by default.
	go mgr.Start()

	// Stop() will shutdown the http server gracefully
	// mgr.Stop()

	// busy working....
	newarr := array.NewArrayList()
	for i := 0; i < 400; i++ {
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
