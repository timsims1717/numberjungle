package systems

import (
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/vars"
	"numberjungle/pkg/object"
	"numberjungle/pkg/timing"
	"numberjungle/pkg/world"
)

var (
	threshold = 24.
	speed     = 150.
)

func MovementSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasMovement) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		coords, okC := result.Components[myecs.Coords].(world.Coords)
		movement, okM := result.Components[myecs.Moving].(*data.Movement)
		if okO && okC && okM && obj != nil && movement != nil {
			// determine target coords and pos
			tarCo := coords
			if !movement.Return {
				switch movement.Direction {
				case data.Up:
					tarCo.Y++
				case data.Down:
					tarCo.Y--
				case data.Right:
					tarCo.X++
				case data.Left:
					tarCo.X--
				}
			}
			target := world.MapToWorld(tarCo)
			// check if past threshold
			moveBack := false
			if !movement.Return && !LegalMove(tarCo) {
				switch movement.Direction {
				case data.Up:
					moveBack = target.Y - threshold <= obj.Pos.Y
				case data.Down:
					moveBack = target.Y + threshold >= obj.Pos.Y
				case data.Left:
					moveBack = target.X + threshold >= obj.Pos.X
				case data.Right:
					moveBack = target.X - threshold <= obj.Pos.X
				}
			}
			if moveBack {
				movement.Return = true
				continue
			}
			// move toward target
			if target.X > obj.Pos.X {
				obj.Pos.X += speed * timing.DT
			} else if target.X < obj.Pos.X {
				obj.Pos.X -= speed * timing.DT
			}
			if target.Y > obj.Pos.Y {
				obj.Pos.Y += speed * timing.DT
			} else if target.Y < obj.Pos.Y {
				obj.Pos.Y -= speed * timing.DT
			}
			// check if done
			done := false
			if movement.Return {
				switch movement.Direction {
				case data.Up:
					done = obj.Pos.Y <= target.Y
				case data.Down:
					done = obj.Pos.Y >= target.Y
				case data.Left:
					done = obj.Pos.X >= target.X
				case data.Right:
					done = obj.Pos.X <= target.X
				}
			} else {
				switch movement.Direction {
				case data.Up:
					done = obj.Pos.Y >= target.Y
				case data.Down:
					done = obj.Pos.Y <= target.Y
				case data.Left:
					done = obj.Pos.X <= target.X
				case data.Right:
					done = obj.Pos.X >= target.X
				}
			}
			if done {
				obj.Pos.Y = target.Y
				obj.Pos.X = target.X
				result.Entity.RemoveComponent(myecs.Moving)
				result.Entity.AddComponent(myecs.Coords, tarCo)
			}
		}
	}
}

func LegalMove(coords world.Coords) bool {
	if coords.X < 0 || coords.Y < 0 || coords.X >= vars.Width || coords.Y >= vars.Height {
		return false
	}
	for _, result := range myecs.Manager.Query(myecs.HasCoords) {
		if co, ok := result.Components[myecs.Coords].(world.Coords); ok {
			if co == coords {
				if result.Entity.HasComponent(myecs.Occupy) {
					return false
				}
			}
		}
	}
	return true
}