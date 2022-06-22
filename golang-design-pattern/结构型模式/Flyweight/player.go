package main

type player struct {
	dress     dress
	playType  string
	lat, long int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		dress:    dress,
		playType: playerType,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
