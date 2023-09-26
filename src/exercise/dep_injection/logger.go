package dep_injection

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (adapter LoggerAdapter) Log(message string) {
	adapter(message)
}
