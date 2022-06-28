package memento

import "fmt"

//用于保存程序内部状态到外部，又不希望暴露内部状态的情形

type Memento interface{}


type gameMemento struct {
	hp, mp int
}

type Game struct {
	hp, mp int			//小写  不暴露
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}