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


	// Tags
	IsObject   = ecs.BuildTag(Object)
	IsTemp     = ecs.BuildTag(Temp, Object)
	HasParent  = ecs.BuildTag(Object, Parent)
	IsDrawable = ecs.BuildTag(Object, Drawable)
	HasAnim    = ecs.BuildTag(Animated)
	HasUpdate  = ecs.BuildTag(Update)
)

type ClearFlag bool