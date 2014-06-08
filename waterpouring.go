package waterpouring

import (
	"fmt"
	"container/list"
	"strconv"
)

type Problem struct {
	Capacities []int
	Goal int
}

type state struct {
	cont []int
	how  string
	prev *state
	key  string
}

func newState(cont []int, how string, prev *state) state {
	s := state{cont: cont, how: how, prev: prev}
	b := make([]byte, 0)
	for _, c := range cont {
		b = strconv.AppendInt(b, int64(c), 10)
		b = append(b, ',')
	}
	s.key = string(b)
	return s
}

func (p Problem) next(s state) []state {
	var states []state
	for i, ci := range s.cont {
		for j, cj := range s.cont {
			if i == j {
				if ci < p.Capacities[i] {
					c1 := make([]int, len(s.cont))
					copy(c1, s.cont)
					c1[i] = p.Capacities[i]
					states = append(states, newState(c1, fmt.Sprintf("fill %d", i), &s))
				}
				if ci > 0 {
					c1 := make([]int, len(s.cont))
					copy(c1, s.cont)
					c1[i] = 0
					states = append(states, newState(c1, fmt.Sprintf("spill %d", i), &s))
				}
			} else if ci > 0 {
				vacant := p.Capacities[j] - cj
				if vacant > 0 {
					var amount int
					if s.cont[i] > vacant {
						amount = vacant
					} else {
						amount = ci
					}
					c1 := make([]int, len(s.cont))
					copy(c1, s.cont)
					c1[i] -= amount
					c1[j] += amount
					states = append(states, newState(c1, fmt.Sprintf("pour %d %d", i, j), &s))
				}
			}
		}
	}
	return states
}

func (p Problem) final(s state) bool {
	for _, c := range s.cont {
		if c == p.Goal {
			return true
		}
	}
	return false
}

func (p Problem) Solve() []string {
	q := list.New()
	initial := newState(make([]int, len(p.Capacities)), "start", nil)
	seen := make(map[string]struct{})
	q.PushBack(initial)
	seen[initial.key] = struct{}{}
	for q.Len() != 0 {
		top := q.Remove(q.Front()).(state)
		if p.final(top) {
			steps := make([]string, 0)
			for t := &top; t != nil; t = t.prev {
				steps = append(steps, t.how)
			}
			for i, j := 0, len(steps)-1; i < j; i, j = i+1, j-1 {
				steps[i], steps[j] = steps[j], steps[i]
			}
			return steps[1:]
		}
		n := p.next(top)
		for _, s := range n {
			if _, present := seen[s.key]; !present {
				q.PushBack(s)
				seen[s.key] = struct{}{}
			}
		}
	}
	return nil
}

