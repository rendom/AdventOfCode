package main

type Queue []Node

func (s *Queue) shift() Node {
	old := *s
	last := old[0]
	*s = old[1:]
	return last
}

func (s *Queue) pop() Node {
	old := *s
	n := len(old)
	last := old[n-1]
	*s = old[:n-1]
	return last
}

func (s *Queue) append(p Node) {
	*s = append(*s, p)
}
