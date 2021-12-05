package gameSystem

import (
	"math/rand"
	"time"
)

func createAndShowAllHeroes(heroes *[]Hero) string {
	CreateRandomHeroes(heroes)
	var str string
	for _, v := range *heroes {
		str += v.PrintListHeroes() + "\n"
	}
	return str
}
func Run32() (string, *Hero) {
	heroes := make([]Hero, 32)
	roundHeroes := make([]Hero, 32)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	//CreateRandomHeroes(&heroes)
	resultStr = createAndShowAllHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr += "IN THIS FIGHT, " + heroes[0].GetName() + " WON!!!"
	return resultStr, &heroes[0]
}

func Run64() (string, *Hero) {
	heroes := make([]Hero, 64)
	roundHeroes := make([]Hero, 64)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	//CreateRandomHeroes(&heroes)
	resultStr = createAndShowAllHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr = resultStr + "IN THIS FIGHT, " + heroes[0].GetName() + " WON!!!"
	return resultStr, &heroes[0]
}
func Run128() (string, *Hero) {
	heroes := make([]Hero, 128)
	roundHeroes := make([]Hero, 128)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	//CreateRandomHeroes(&heroes)
	resultStr = createAndShowAllHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr = resultStr + "IN THIS FIGHT, " + heroes[0].GetName() + " WON!!!"
	return resultStr, &heroes[0]
}

func ToFight(i int, heroes *[]Hero, downstream, downstream2 chan Hero) {
	first := 2 * i
	if first < len(*heroes) && (len(*heroes) != 1 || len(*heroes) != 0) {
		downstream <- (*heroes)[first]
	} else {
		return
	}
	second := 2*i + 1
	if second < len(*heroes) && (len(*heroes) != 1 || len(*heroes) != 0) {
		downstream2 <- (*heroes)[second]
	} else {
		return
	}
}

func MakeFight(resultStr *string, upstream, upstream2, downstream chan Hero) {
	for v := range upstream {
		for p := range upstream2 {
			*resultStr = *resultStr + "This battle between " + v.GetName() + " and " + p.GetName() + "\n"
			for {
				if v.amountStamina() > 20 {
					*resultStr += v.SecondSkill(p) + "\n"
				} else if v.amountStamina() > 10 {
					*resultStr += v.FirstSkill(p) + "\n"
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						*resultStr = *resultStr + v.GetName() + " Attack\n"
						*resultStr += v.Attack(p) + "\n"
					} else {
						*resultStr = *resultStr + v.GetName() + " Defend\n"
						*resultStr += v.Defend() + "\n"
					}
				}
				if p.Health() <= 0 {
					*resultStr = *resultStr + "In this battle win " + v.GetName() + "\n"
					*resultStr = *resultStr + "In this battle lose " + p.GetName() + "\n"
					downstream <- v
					return
				}
				if p.amountStamina() > 20 {
					*resultStr += p.SecondSkill(v) + "\n"
				} else if p.amountStamina() > 10 {
					*resultStr += p.FirstSkill(v) + "\n"
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						*resultStr = *resultStr + p.GetName() + " Attack\n"
						*resultStr += p.Attack(v) + "\n"
					} else {
						*resultStr = *resultStr + p.GetName() + " Defend\n"
						*resultStr += p.Defend() + "\n"
					}
				}
				if v.Health() <= 0 {
					*resultStr = *resultStr + " In this battle win " + p.GetName() + "\n"
					*resultStr = *resultStr + " In this battle lose " + v.GetName() + "\n"
					downstream <- p
					return
				}
			}
		}
	}
}
