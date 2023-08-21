package systems

import (
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/pkg/object"
	"numberjungle/pkg/util"
)

func ClickableSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsClickable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO {
			if util.PointInside(data.TheInput.World, obj.Rect, obj.Mat) &&
				data.TheInput.Get("click").JustPressed() {
				click := result.Components[myecs.Clickable]
				if fn, ok := click.(*data.FrameFunc); ok {
					if fn.Func() {
						myecs.Manager.DisposeEntity(result)
					}
				}
			}
		}
	}
}
