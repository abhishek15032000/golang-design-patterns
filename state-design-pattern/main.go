package main

import "fmt"

// The State pattern allows an object to change its behavior when its internal state changes.
// Think of it like a TV:
// - When the TV is in the "Off" state, it does nothing and the screen is blank.
// - When the TV is in the "On" state, it plays a show.
// Instead of writing massive `if-else` blocks everywhere (if state == ON do X, else do Y),
// this pattern splits each state's behavior into its own dedicated struct (like On and Off).

type state interface {
	speak_up()
}

type On struct{}

func (o *On) speak_up() {
	fmt.Println("kapil sharma show")
}

type Off struct{}

func (o *Off) speak_up() {
	fmt.Println("tv is off")
}

type Tv struct {
	current_position_of_tv state
}

func (t *Tv) getState() state {
	return t.current_position_of_tv
}

func (t *Tv) setState(x state) {
	t.current_position_of_tv = x
}

func (t *Tv) press_power_button() {
	t.current_position_of_tv.speak_up()
}

func main() {
	tv := &Tv{
		current_position_of_tv: &Off{},
	}
	tv.press_power_button()

	tv.setState(&On{})
	tv.press_power_button()
}
