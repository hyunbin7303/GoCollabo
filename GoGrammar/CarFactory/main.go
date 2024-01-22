package main

import "fmt"

// source from https://www.youtube.com/watch?v=Gmlh0NrvzP0

type SpecialPosition struct {
	Position
}
type Attack struct {
	damage int
	name   string
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}
func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
	*Attack
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func (e *Enemy) Move(x, y float64) {
	e.x += x
	e.y += y
}

func main() {
	fmt.Println("Test")

	bossMonster := NewEnemy()
	bossMonster.Move(1.1, 10.4)
	fmt.Println("Monster Location: ", bossMonster.Position)

	bossMonster.MoveSpecial(1.1, 10.4)
	fmt.Println("Move Special : ", bossMonster.Position)

	// bossMonster.Attack(30)

	player := NewPlayer()
	player.Move(1.1, 10.4)
	fmt.Println(player.Position)

}
