package monster

import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T) {
	// 先创建一个monster实例
	monster := Monster{
		Name:  "红孩儿",
		Age:   201,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("序列化函数测试错误,期望为%v, 实际为%v\n", true, res)
	}
	t.Logf("monster.Store() 测试成功!!")
}

// 测试反序列化
func TestReStore(t *testing.T) {
	// 先创建一个Monster实例,不需要指定字段内容
	var monster Monster
	monster.ReStore()
	res := monster.ReStore()
	if !res {
		t.Fatalf("序列化函数测试错误,期望为%v, 实际为%v\n", true, res)
	}
	fmt.Println("名字是", monster.Name)
	// 进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误, 期望为%v, 实际为%v\n", "红孩儿", monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功!!")
}
