package main

import "fmt"

type Gun interface {
	Name() string
	Damage() int
	AmmoCapacity() int
	NoiseLevel() int
}

type BaseGun struct {
	name         string
	damage       int
	ammoCapacity int
	noiseLevel   int
}

func (g *BaseGun) Name() string      { return g.name }
func (g *BaseGun) Damage() int       { return g.damage }
func (g *BaseGun) AmmoCapacity() int { return g.ammoCapacity }
func (g *BaseGun) NoiseLevel() int   { return g.noiseLevel }

type ExtendedMagazineGunDecorator struct {
	Gun
}

func (g *ExtendedMagazineGunDecorator) AmmoCapacity() int {
	return g.Gun.AmmoCapacity() + 10
}

type SilencerGunDecorator struct {
	Gun
}

func (g *SilencerGunDecorator) NoiseLevel() int {
	suppressedNoise := g.Gun.NoiseLevel() - 5
	if suppressedNoise < 1 {
		return 1
	}
	return suppressedNoise
}

func main() {
	fmt.Println("Creating Basic carbine")
	var carbine Gun = &BaseGun{name: "AK-102", damage: 40, ammoCapacity: 30, noiseLevel: 9}
	printGunDetails(carbine)

	fmt.Println()

	fmt.Println("Applying Extended Magazine")
	carbine = &ExtendedMagazineGunDecorator{Gun: carbine}
	printGunDetails(carbine)

	fmt.Println()

	fmt.Println("Applying Silencer")
	carbine = &SilencerGunDecorator{Gun: carbine}
	printGunDetails(carbine)
}

func printGunDetails(g Gun) {
	fmt.Printf("Name: %s, Damage: %d, Ammo Capacity: %d, Noise Level: %d\n",
		g.Name(), g.Damage(), g.AmmoCapacity(), g.NoiseLevel())
}
