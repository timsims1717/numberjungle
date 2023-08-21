package data

import "github.com/bytearena/ecs"

var Player *player

type player struct{
	Move *Movement
	Char *ecs.Entity
}

func NewPlayer() {
	Player = &player{}
}

var Expression *expression

type expression struct{
	Coins []*Coin
}

func NewExpression() {
	Expression = &expression{}
}