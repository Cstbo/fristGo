package game

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name   string
	HP     int
	Attack int
}

func (m *Monster) GetName() string {
	return m.Name
}

func (m *Monster) AttackValue() int {
	base := m.Attack
	delta := rand.Intn(5) - 3 // -3 ~ +1（你可以自己调）
	dmg := base + delta
	if dmg < 1 {
		dmg = 1
	}
	return dmg
}

func (m *Monster) TakeDamage(damage int) {
	m.HP -= damage
	if m.HP < 0 {
		m.HP = 0
	}
}

func (m *Monster) IsDead() bool {
	return m.HP <= 0
}

func (m Monster) PrintStatus() {
	fmt.Printf("怪物[%s] HP[%d] 攻击力[%d]\n",
		m.Name, m.HP, m.Attack)
}
