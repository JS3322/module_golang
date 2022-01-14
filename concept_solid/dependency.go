package main

// dependency inversion principle

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	BeAttacked()
}

func Attack(attacker *Attackable, defender *BeAttacked) {
	attacker.Attack(defender)
}
