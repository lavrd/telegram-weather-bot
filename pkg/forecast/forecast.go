package forecast

type Data struct {
}

type Forecast interface {
	GetNow() (*Data, error)
}
