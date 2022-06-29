package main

import "fmt"

type Hero struct {
	ID       int
	Name     string
	NickName string
	Prev     *Hero
	Next     *Hero
}

func NewHero(id int, name string, nickname string) *Hero {
	return &Hero{
		ID:       id,
		Name:     name,
		NickName: nickname,
	}
}

func (h *Hero) AppendHero(hero *Hero) {
	tmp := h
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = hero
	hero.Prev = h
	fmt.Println("h == ", &h)
	fmt.Println("tmp == ", &tmp)
	fmt.Println("hero == ", &hero)
}

func (h *Hero) ListHeros() {
	tmp := h
	for {
		fmt.Printf("%v --> ", tmp)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	fmt.Println()
}

func main() {
	head := NewHero(0, "", "")
	hero01 := NewHero(1, "宋江", "及时雨")
	head.AppendHero(hero01)
	head.ListHeros()
}
