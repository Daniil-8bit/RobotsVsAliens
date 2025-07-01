package main

type AreaConfiguration struct {
	name        string
	linesAmount int
}

type Level struct {
	AreaConfiguration
	levelNumber int
	levelName   string
	GameField
}

func (l *Level) levelInit(ac AreaConfiguration, number int, name string, gf GameField) {

	l.levelName = name
	l.levelNumber = number
	l.AreaConfiguration = ac
	l.GameField = gf
}
