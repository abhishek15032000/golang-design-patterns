// The Adapter pattern acts as a bridge between two incompatible interfaces.
// Just like a physical power adapter lets you plug a European charger into a US outlet,
// a software adapter wraps an incompatible system so it can be used
// by existing code without making any breaking changes.

package main

import "fmt"

type Mobile interface {
	charging_mobile()
}

type Apple struct{}

func (a *Apple) charging_mobile() {
	fmt.Println("we are charging an apple mobile")
}

type Client struct{}

func (c *Client) just_put_it_on_charge(mob Mobile) {
	mob.charging_mobile()
}

// if we were to extend this for android mobile, 1 thing we could do is do the same android struct and everything, but how many times would
// we do it, there is repeatedness in this thing.

// we can come up with an adapter, in which you can pass all the extra systems you wanna support.

type Android struct{}

func (a *Android) charging_mobile() {
	fmt.Println("we are charging an android mobile")
}

type Adapter struct {
	ax *Android
}

func (a *Adapter) charging_mobile() {
	a.ax.charging_mobile()
}

func main() {
	a := &Apple{}
	c := &Client{}
	c.just_put_it_on_charge(a)

	ax := &Android{}
	adapter := &Adapter{
		ax: ax,
	}
	c.just_put_it_on_charge(adapter)
}
