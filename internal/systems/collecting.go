package systems

import (
	"numberjungle/internal/myecs"
	"numberjungle/pkg/object"
	"numberjungle/pkg/world"
)

func CollectingSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsPlayer) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		coords, okC := result.Components[myecs.Coords].(world.Coords)
		if okO && okC && obj != nil && !obj.Gone {
			for _, result1 := range myecs.Manager.Query(myecs.IsCollectible) {
				obj1, ok1O := result1.Components[myecs.Object].(*object.Object)
				coords1, ok1C := result1.Components[myecs.Coords].(world.Coords)
				//collect := result1.Components[myecs.Collect]
				if ok1O && ok1C && obj1 != nil && coords1 == coords {
					//if coin, ok := collect.(*data.Coin); ok {
					//
					//}
					result1.Entity.RemoveComponent(myecs.Collect)
					result1.Entity.AddComponent(myecs.Temp, myecs.Has{})
				}
			}
		}
	}
}