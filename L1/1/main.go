package main

import "fmt"

/*  Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// первый случай
type HumanFirst struct {
	name                string
	loveProgramLenguage string
}

type Programmer struct {
	WorkingLanguage string
}

type ActionFirst struct {
	HumanFirst
	Programmer
	AttitudeSports bool
}

func (h *HumanFirst) learnName() {
	fmt.Printf("Your name %s\n", h.name)
}

func (p *Programmer) learnWorkingLanguage() {
	fmt.Printf("You working lenguage %s\n", p.WorkingLanguage)
}

//-------------------------------------------------------------

// второй случай

type HumanSecond struct {
	name                string
	loveProgramLenguage string
}

type HumanSecondInterface interface {
	getProgrammingLenguage()
	getName()
}

type ActionSecond struct {
	HumanSecondInterface
}

func (h *HumanSecond) getProgrammingLenguage() {
	fmt.Printf("Your programming lenguage %s\n", h.loveProgramLenguage)
}

func (h *HumanSecond) getName() {
	fmt.Printf("Your name %s\n", h.name)
}

//---------------------------------------------------------------

func main() {
	artem := ActionFirst{
		HumanFirst: HumanFirst{
			name:                "artem",
			loveProgramLenguage: "GO",
		}, AttitudeSports: true,
	}

	artem.learnWorkingLanguage() // отобразит пустое поле
	artem.learnName()

	ivan := ActionFirst{
		HumanFirst: HumanFirst{
			name:                "ivan",
			loveProgramLenguage: "GO",
		}, Programmer: Programmer{
			WorkingLanguage: "GO",
		}, AttitudeSports: false,
	}

	ivan.learnName()
	ivan.learnWorkingLanguage()

	nikola := ActionSecond{
		HumanSecondInterface: &HumanSecond{
			name:                "nikola",
			loveProgramLenguage: "GO",
		},
	}

	nikola.getName()
	nikola.getProgrammingLenguage()

}
