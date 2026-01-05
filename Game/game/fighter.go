package game

import "fmt"

// Fighter 接口：所有能参与战斗的对象都要满足它
type Fighter interface {
	GetName() string       // 名字，用来打印
	AttackValue() int      // 这次攻击能打多少伤害
	TakeDamage(damage int) // 挨打，扣血
	IsDead() bool          // 是否死亡
	PrintStatus()          // 打印当前状态
}
 
// 让 attacker 攻击 defender，打印过程
func attackOnce(attacker, defender Fighter) {
	damage := attacker.AttackValue()
	fmt.Printf("%s 攻击了 %s，造成 %d 点伤害！\n",
		attacker.GetName(), defender.GetName(), damage)

	defender.TakeDamage(damage)
	defender.PrintStatus()
	fmt.Println()
}

// 进行一场战斗，直到一方死亡
func Battle(a, b Fighter) {
	fmt.Println("战斗开始！")
	a.PrintStatus()
	b.PrintStatus()
	fmt.Println()

	round := 1
	for {
		fmt.Printf("===== 回合 %d =====\n", round)

		// A 先手
		attackOnce(a, b)
		if b.IsDead() {
			fmt.Printf("%s 倒下了，%s 获胜！\n", b.GetName(), a.GetName())
			break
		}

		// B 反击
		attackOnce(b, a)
		if a.IsDead() {
			fmt.Printf("%s 倒下了，%s 获胜！\n", a.GetName(), b.GetName())
			break
		}

		round++
	}
}
