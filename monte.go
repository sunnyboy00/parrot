package main

type PlayoutResult struct {
	score float64
	hand  Hand
}

type MonteCarloPlayer struct {
	State    GameState
	tryCount int
}

func calcScore(sim Simulator) float64 {
	return float64(sim.Score() + sim.GetMaxTile()*200)
}

func (p *MonteCarloPlayer) SetState(s GameState) {
	p.State = s
}

func (p *MonteCarloPlayer) Playout(firstHand Hand, res chan PlayoutResult, gen func() uint32) {
	avg := 0.0
	for cnt := 0; cnt < p.tryCount; cnt++ {
		var sim Simulator = &Kanna{p.State.Grid, p.State.Score, p.State.Over, gen}
		if !sim.Move(firstHand) {
			avg -= 100
			break
		}
		sim.AddRandomCell()

		for sim.Move(intToHand(int(gen() % 4))) {
			if !sim.AddRandomCell() {
				break
			}
		}
		avg += calcScore(sim)
	}
	res <- PlayoutResult{avg / float64(p.tryCount), firstHand}
}

func (p *MonteCarloPlayer) NextHand(gen func() uint32) Hand {
	res := make(chan PlayoutResult)
	go p.Playout(Up, res, gen)
	go p.Playout(Right, res, gen)
	go p.Playout(Down, res, gen)
	go p.Playout(Left, res, gen)

	bestAvg, ret := 0.0, Up
	for i := 0; i < 4; i++ {
		r := <-res
		if bestAvg < r.score {
			bestAvg = r.score
			ret = r.hand
		}
	}

	// fmt.Printf("after best move %v:\n", ret)
	// var sim Midori = Midori{encode2(&p.State.Grid), 0, p.State.Over, r}
	// sim.Move(ret)
	// fmt.Printf("%v\n", sim.GetState())

	return ret
}
