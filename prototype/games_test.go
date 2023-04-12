package prototype_test

import (
	"sdp/prototype"
	"testing"
)

func TestCloneWithoutPrototype(t *testing.T) {
	character1 := prototype.NewCharacter("raden", 1, 20, 20, prototype.Role{
		Name:  "warrior",
		Level: prototype.Beginner,
	})

	character2 := prototype.NewCharacter("raden", 1, 20, 20, prototype.Role{
		Name:  "warrior",
		Level: prototype.Beginner,
	})

	t.Log("=== Character 1 ===")
	PrintCharacterInfo(t, character1)

	t.Log("=== Character 2 ===")
	PrintCharacterInfo(t, character2)
}

func TestWithPrototype(t *testing.T) {
	character1 := prototype.NewCharacter("raden", 1, 20, 20, prototype.Role{
		Name:  "warrior",
		Level: prototype.Beginner,
	})

	character2 := character1.Clone()

	t.Log("=== Character 1 ===")
	PrintCharacterInfo(t, character1)

	t.Log("=== Character 2 ===")
	PrintCharacterInfo(t, character2)
}

func PrintCharacterInfo(t *testing.T, character prototype.Character) {
	t.Log("Name:", character.Name)
	t.Log("Level:", character.Level)
	t.Log("Hp:", character.Hp)
	t.Log("Mp:", character.Mp)
	t.Log("Role:", character.Role)
}
