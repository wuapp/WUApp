package wua

type Widget interface {
	Register()
}

func Settings(settings WindowSettings) {

}

func Run() (err error) {
	err = create()
	defer exit()

	return
}

func RegisterWidgets(widgets ...Widget) {
	for _, w := range widgets {
		w.Register()
	}
}
