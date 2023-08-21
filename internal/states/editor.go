package states

import (
	"github.com/faiface/pixel/pixelgl"
	"numberjungle/internal/data"
	"numberjungle/internal/systems"
	"numberjungle/internal/vars"
	"numberjungle/pkg/camera"
	"numberjungle/pkg/img"
	"numberjungle/pkg/state"
	"numberjungle/pkg/world"
)

var EditorState = &editorState{}

type editorState struct {
	*state.AbstractState

	Picker   []interface{}
	EditMode bool
}

func (s *editorState) Unload() {
	systems.ClearSystem()
}

func (s *editorState) Load() {
	camera.Main.Pos.X = (vars.Width - 5) * world.TileSize * 0.5
	camera.Main.Pos.Y = (vars.Height) * world.TileSize * 0.5

	initUIBorder()
	initGamePanel()
	initEditorPanel()

	s.EditMode = true
	initPicker()

	data.EditPuzzle = blankPuzzle()
	setPlayerStart(data.EditPuzzle.Start)
}

func (s *editorState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win, camera.Main.Mat)


	systems.TemporarySystem()
	systems.FullTransformSystem()
	systems.AnimationSystem()

	camera.Main.Update(win)
}

func (s *editorState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win)
	img.Draw(win)
}

func (s *editorState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}