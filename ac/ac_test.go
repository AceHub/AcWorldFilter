package ac

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go-wordfilter/common"
	"log"
	"testing"
)

func TestAC(t *testing.T) {
	ac := NewAc()
	ac.LoadWords(common.GetWords())
	result := ac.Search("hello av java 毛泽东 sm 气枪 测试, 支付宝 ")
	log.Println("len:", len(result))
	for _, item := range result {
		fmt.Printf("%+v \n", item)
	}
}

func TestB(t *testing.T) {
	ac := NewAc()
	ac.LoadWords([]*common.SensitiveWords{
		&common.SensitiveWords{
			Word: "三四",
			Rank: 1,
		},
		&common.SensitiveWords{
			Word: "一二三四",
			Rank: 1,
		},
		&common.SensitiveWords{
			Word: "一二三",
			Rank: 1,
		},
		&common.SensitiveWords{
			Word: "四五六七",
			Rank: 1,
		},
	})
	result := ac.Search("一二三四五六七")
	for _, item := range result {
		fmt.Printf("%+v \n", item)
	}
}

func TestAc_Replace(t *testing.T) {
	ac := NewAc()
	ac.LoadWords(common.GetWords())
	result := ac.Replace("hello av java 毛泽东 sm 气枪 测试, 支付宝 ", 0)
	fmt.Printf("%+v\n", result.NewContent)
}

func TestRun(t *testing.T) {
	println("run")
	ac := NewAc()
	ac.AddWord("ab", 1)
	ac.AddWord("av", 1)
	ac.AddWord("avx", 1)
	ac.AddWord("老胡", 1)
	ac.Make()
	testWord := func(text string, need_result int, msg string) {
		out := ac.Search(text)
		if len(out) != need_result {
			t.Error("error", msg, "find count:", len(out), "expect count", need_result)
		} else {
			t.Log("pass ", msg)
		}
	}
	out := ac.Search("abc 123 你好 3333 ab")
	log.Println("search result out:", len(out))
	for _, item := range out {
		log.Println("item")
		spew.Dump(item)
	}
	out2 := ac.Search("abc 123 老胡 你好 3333 ab")
	log.Println("search result out:", len(out2))
	for _, item := range out2 {
		log.Println("item")
		spew.Dump(item)
	}

	testWord("abcd", 0, "连续字母用例")
	testWord("av", 1, "两头边界")
	testWord("avx", 1, "两头边界")
	testWord("avi", 0, "两头边界 反例")
	testWord("你好av你好have 你好", 1, "中文")
	testWord("哈哈老胡空白符你好", 1, "中文")
	testWord("老胡", 1, "中文")

}

func TestReload(t *testing.T) {
	ac := NewAc()
	ac.LoadWords(common.GetWords())
	result := ac.Search("hello av hava 毛泽东 sm 气枪 测试, 支付宝 ")
	log.Println("len:", len(result))
	for _, item := range result {
		fmt.Println(item)
	}
}
