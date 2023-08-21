package myecs

import "github.com/bytearena/ecs"

var (
	Manager = ecs.NewManager()

	// Components
	Object = Manager.NewComponent()
	Parent = Manager.NewComponent()
	Temp   = Manager.NewComponent()
	Update = Manager.NewComponent()

	Drawable = Manager.NewComponent()
	Animated = Manager.NewComponent()

	Clickable = Manager.NewComponent()

	Moving = Manager.NewComponent()
	Coords = Manager.NewComponent()

	Player  = Manager.NewComponent()
	Occupy  = Manager.NewComponent()
	Collect = Manager.NewComponent()

	// Tags
	IsObject   = ecs.BuildTag(Object)
	IsTemp     = ecs.BuildTag(Temp, Object)
	HasParent  = ecs.BuildTag(Object, Parent)
	IsDrawable = ecs.BuildTag(Object, Drawable)
	HasAnim    = ecs.BuildTag(Animated)
	HasUpdate  = ecs.BuildTag(Update)
	IsClickable = ecs.BuildTag(Object, Clickable)

	IsPlayer      = ecs.BuildTag(Object, Coords, Player)
	HasCoords     = ecs.BuildTag(Object, Coords)
	HasMovement   = ecs.BuildTag(Object, Coords, Moving)
	IsCollectible = ecs.BuildTag(Object, Coords, Collect)
)

type ClearFlag bool

type Has struct{}