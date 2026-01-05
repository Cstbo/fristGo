package main

import (
	"day6game/game" // 这里的 day6game 要和 go.mod 里的 module 名一致
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p1 := &game.Player{
		Name:         "Cstbo",
		Level:        1,
		HP:           100,
		Attack:       20,
		Exp:          0,
		NextLevelExp: 100,
	}

	monsters := []*game.Monster{
		{Name: "史莱姆", HP: 60, Attack: 8},
		{Name: "小野狼", HP: 90, Attack: 12},
		{Name: "山贼头目", HP: 120, Attack: 15},
	}

	for i, m := range monsters {
		fmt.Printf("\n==== 第 %d 关：遇到怪物 [%s] ====\n", i+1, m.Name)
		game.Battle(p1, m)

		if p1.IsDead() {
			fmt.Println("玩家阵亡，游戏结束。")
			return
		}

		reward := (i + 1) * 80
		fmt.Println("战斗胜利！")
		p1.GainExp(reward)
		p1.PrintStatus()
	}

	fmt.Println("\n恭喜你，通关啦！")
}
