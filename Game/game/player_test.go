package game

import "testing"

func TestPlayer_LevelUp(t *testing.T) {
	p := &Player{Name: "Cstbo", Level: 1, HP: 100, Attack: 20}

	p.LevelUp()

	if p.Level != 2 {
		t.Errorf("LevelUp 后等级错误，期望 2，实际 %d", p.Level)
	}
	if p.HP != 120 { // 你现在的逻辑是 +20，就按你的来
		t.Errorf("LevelUp 后 HP 错误，期望 120，实际 %d", p.HP)
	}
	if p.Attack != 25 { // 你现在的逻辑是 +5
		t.Errorf("LevelUp 后攻击力错误，期望 25，实际 %d", p.Attack)
	}
}

func TestPlayer_TakeDamage(t *testing.T) {
	p := &Player{Name: "Cstbo", Level: 1, HP: 50, Attack: 20}

	p.TakeDamage(30)
	if p.HP != 20 {
		t.Errorf("第一次受伤后 HP 错误，期望 20，实际 %d", p.HP)
	}

	p.TakeDamage(100)
	if p.HP != 0 {
		t.Errorf("HP 不应小于 0，期望 0，实际 %d", p.HP)
	}
}
