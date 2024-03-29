package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"fmt"
)

const cityTextureLocation = "textures/city.png"

// CityAssets are the assets for the system
var CityAssets = []string{cityTextureLocation}

type CityMouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type CityBuildingSystem struct {
	world *ecs.World

	mouseTracker CityMouseTracker
}

// New is the initialisation of the System
func (cb *CityBuildingSystem) New(w *ecs.World) {
	cb.world = w
	fmt.Println("CityBuildingSystem was added to the Scene")

	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cb.mouseTracker.BasicEntity, &cb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed F1")

		city := City{BasicEntity: ecs.NewBasic()}

		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{cb.mouseTracker.MouseComponent.MouseX, cb.mouseTracker.MouseComponent.MouseY},
			Width:    30,
			Height:   62,
		}

		texture, err := common.LoadedSprite("textures/city.png")
		if err != nil {
			panic(err)
		}
		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.5, 0.5},
		}

		// Add to the system
		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}

// Remove is called whenever an Entity is removed from the scene, and thus from this system
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}
