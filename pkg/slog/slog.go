package slog

type Slog interface {
	WriteLog()
}

func Writer(sl Slog) {
	sl.WriteLog()
}
