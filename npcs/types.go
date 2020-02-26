package npcs

type Weapon interface {
	Weild() string
}

type Power struct {
	Attack int
	Defense int 
}

type Location struct {
	X, Y, Z float64
}


type Character struct {
	Name string
	Speed int
	Hp int
	Power Power
	Location Location
}

type Attacker struct {
	Power, Dmgbonus int 
}

type Sword struct {
	Attacker
	TwoHanded bool
}

type Gun struct {
	Attacker 
	BulletSremaining int
}
