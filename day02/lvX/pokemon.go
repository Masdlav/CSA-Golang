//type pokemon struct {
//	Name string
//	Lv   int
//	Exp  int
//	Hp   int
//	Ce   int
//}
//
//type Attacker interface {
//	Attack()
//}

package main

import "fmt"

// Person 定义一个游戏人物的结构体
type Person struct {
	Name      string
	Level     int
	Exp       int
	HP        int
	AttackPts int
}

// Attacker 定义一个角色攻击接口
type Attacker interface {
	Attack(target *Person)
}

// Warrior 实现 Attacker 接口的结构体
type Warrior struct {
	Person
}

// Attack 实现 Attack 方法
func (w *Warrior) Attack(target *Person) {
	target.HP -= w.AttackPts
}

func main() {
	// 创建两个角色
	p1 := &Person{Name: "Tom", Level: 10, Exp: 1000, HP: 100, AttackPts: 20}
	p2 := &Person{Name: "Jerry", Level: 8, Exp: 800, HP: 80, AttackPts: 30}

	// 让 p1 攻击 p2
	var attacker Attacker
	attacker = &Warrior{*p1}
	attacker.Attack(p2)

	// 打印 p2 的血量变化情况
	fmt.Printf("%s's HP: %d -> %d\n", p2.Name, p2.HP+p1.AttackPts, p2.HP)
}
