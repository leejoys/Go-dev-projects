package electronic

//Phone (Brand, Model, Type)
type Phone interface {
	Brand() string
	Model() string
	Type() string
}

//StationPhone (ButtonsCount)
type StationPhone interface {
	ButtonsCount() int
}

//Smartphone (OS)
type Smartphone interface {
	OS() string //название операционной системы
}

//applePhone model, os
type applePhone struct {
	phonemodel string
	phoneos    string
}

//Brand applePhone brand
func (a applePhone) Brand() string {
	return "Apple"
}

//Model applePhone model
func (a applePhone) Model() string {
	return a.phonemodel
}

//Type applePhone type
func (a applePhone) Type() string {
	return "smartphone"
}

//OS applePhone os
func (a applePhone) OS() string {
	return a.phoneos
}

//androidPhone brand, model, os
type androidPhone struct {
	phonebrand string
	phonemodel string
	phoneos    string
}

//Brand androidPhone brand
func (a androidPhone) Brand() string {
	return a.phonebrand
}

//Model androidPhone model
func (a androidPhone) Model() string {
	return a.phonemodel
}

//Type androidPhone type
func (a androidPhone) Type() string {
	return "smartphone"
}

//OS androidPhone os
func (a androidPhone) OS() string {
	return a.phoneos
}

//radioPhone brand, model, buttons
type radioPhone struct {
	phonebrand   string
	phonemodel   string
	phonebuttons int
}

//Brand radioPhone brand
func (a radioPhone) Brand() string {
	return a.phonebrand
}

//Model radioPhone model
func (a radioPhone) Model() string {
	return a.phonemodel
}

//Type radioPhone type
func (a radioPhone) Type() string {
	return "station"
}

//ButtonsCount radioPhone buttons
func (a radioPhone) ButtonsCount() int {
	return a.phonebuttons
}

//ConsrustorApple construct applePhone (_, model, _, os)
func ConsrustorApple(a, b string) applePhone {
	return applePhone{
		phonemodel: a,
		phoneos:    b,
	}
}

//ConsrustorAndroid construct androidPhone (brand, model, _, os)
func ConsrustorAndroid(a, b, c string) androidPhone {
	return androidPhone{
		phonebrand: a,
		phonemodel: b,
		phoneos:    c,
	}
}

//ConsrustorRadio construct radioPhone (brand, model, _, buttons)
func ConsrustorRadio(a, b string, c int) radioPhone {
	return radioPhone{
		phonebrand:   a,
		phonemodel:   b,
		phonebuttons: c,
	}
}
