package main

// interface segregation principle

type Talkable interface {
	Talk()
}

type Moveable interface {
	Move()
}
