package npcs
import "math"
import "fmt"

func (self Location) String() string {
	return fmt.Sprintf("(%f, %f, %f)", self.X, self.Y, self.Z)
}

func (self Location) euclideanDistance(target Location) float64 {
	deltaX := self.X - target.X
	deltaY := self.Y - target.Y
	return math.Sqrt(deltaX * deltaX + deltaY * deltaY)
}

func (chacater Character) DistanceTo(target Character) float64 {
	return chacater.Location.euclideanDistance(target.Location)
}

func Wielder(self Weapon) string {
	return "Wielding...\n" + self.Weild()
}

func (self Sword) Weild() string {
	return "You're wielding a sword!"
} 

func (self Gun) Weild() string {
	return "You're wielding a gun!"
} 