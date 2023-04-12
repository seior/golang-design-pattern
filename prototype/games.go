package prototype

type Level string

const (
	Beginner     Level = "beginner"
	Intermediate Level = "intermediate"
	Advanced     Level = "advanced"
)

type Role struct {
	Name  string
	Level Level
}

type Character struct {
	Name  string
	Level uint32
	Hp    uint32
	Mp    uint32
	Role  Role
}

func NewCharacter(name string, level uint32, hp uint32, mp uint32, role Role) Character {
	return Character{
		Name:  name,
		Level: level,
		Hp:    hp,
		Mp:    mp,
		Role:  role,
	}
}

func (character *Character) Clone() Character {
	return Character{
		Name:  character.Name,
		Level: character.Level,
		Hp:    character.Hp,
		Mp:    character.Mp,
		Role:  character.Role,
	}
}
