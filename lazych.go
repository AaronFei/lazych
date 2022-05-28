package lazych

type channel_t[T any] struct {
	C chan T
}

var channels = map[string]interface{}{}

func GetChannel[T any](name string) chan T {
	if _, ok := channels[name]; ok {
		return channels[name].(channel_t[T]).C
	}
	ch := make(chan T)
	channelInfo := channel_t[T]{
		C: ch,
	}
	channels[name] = channelInfo
	return channelInfo.C
}
