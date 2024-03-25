package main

type s struct {
	On    bool
	Ammo  int
	Power int
}

func (v *s) Shoot() bool {
	if v.On == true && v.Ammo > 0 {
		v.Ammo--
		return true
	}
	return false
}

func (v *s) RideBike() bool {
	if v.On == true && v.Power > 0 {
		v.Power--
		return true
	}
	return false
}
