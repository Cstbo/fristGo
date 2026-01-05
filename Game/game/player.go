package game

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name         string
	Level        int
	HP           int
	Attack       int
	Exp          int
	NextLevelExp int
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) AttackValue() int {
	base := p.Attack
	delta := rand.Intn(7) - 3 // -3 ~ +3
	dmg := base + delta
	if dmg < 1 {
		dmg = 1
	}
	// 10% 暴击
	if rand.Intn(100) < 10 {
		fmt.Printf("玩家[%s] 触发暴击！\n", p.Name)
		dmg *= 2
	}
	return dmg
}

func (p *Player) TakeDamage(damage int) {
	p.HP -= damage
	if p.HP < 0 {
		p.HP = 0
	}
}

func (p *Player) IsDead() bool {
	return p.HP <= 0
}

func (p Player) PrintStatus() {
	fmt.Printf("玩家[%s] 等级[%d] HP[%d] 攻击力[%d] EXP[%d/%d]\n",
		p.Name, p.Level, p.HP, p.Attack, p.Exp, p.NextLevelExp)
}

func (p *Player) LevelUp() {
	fmt.Printf("恭喜玩家[%s]升级！\n", p.Name)
	p.Level++
	p.HP += 20
	p.Attack += 5
}

func (p *Player) GainExp(amount int) {
	if amount <= 0 {
		return
	}

	fmt.Printf("玩家[%s] 获得 %d 点经验！\n", p.Name, amount)
	p.Exp += amount

	// 可能一次拿很多经验，所以用 for 处理连升多级的情况
	for p.Exp >= p.NextLevelExp {
		p.Exp -= p.NextLevelExp
		p.LevelUp()

		// 简单规则：每升一级，下一级需求 +50
		p.NextLevelExp += 50
	}
}
