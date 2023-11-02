package animal

var (
	AllAnimals = []Animal{
		// gc boom
		&Dog{},
		// goroutine boom
		&Wolf{},
		// mutex boom
		&Cat{},
		// cpu boom
		&Tiger{},
		// heap boom
		&Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()
	Eat()
	Drink()
	Shit()
	Pee()
}

type Muridae interface {
	Animal
	Hole()
	Steal()
}

type Felidae interface {
	Animal
	Climb()
	Sneak()
}

type Canidae interface {
	Animal
	Run()
	Howl()
}
