package lazych

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type channel_t[T any] struct {
	C chan T
}

type rrChannel_t[T any] struct {
	Request  chan T
	Response chan T
}

var (
	channels      = map[string]interface{}{}
	channelsMutex = sync.Mutex{}

	rrChannels      = map[string]interface{}{}
	rrChannelsMutex = sync.Mutex{}
)

// Create channel
func CreateCh[T any](name string, size int) error {
	channelsMutex.Lock()
	defer channelsMutex.Unlock()

	if _, ok := channels[name]; ok {
		return fmt.Errorf("%s exist", name)
	} else {
		channels[name] = channel_t[T]{
			C: make(chan T, size),
		}
		return nil
	}
}

// Get the channel[T] with the name.
// Return error if the [T] is not same as the recorded [T] or channel's name does not exist
func GetCh[T any](name string) (chan T, error) {
	channelsMutex.Lock()
	defer channelsMutex.Unlock()

	if _, ok := channels[name]; ok {
		t := reflect.TypeOf(channels[name])
		t2 := reflect.TypeOf(channel_t[T]{})
		if t.Name() == t2.Name() {
			return channels[name].(channel_t[T]).C, nil
		} else {
			startIndex := strings.Index(t.Name(), "[") + 1
			return nil, fmt.Errorf("Required T is \"%s\" but get \"%s\"", t.Name()[startIndex:len(t.Name())-1], t2.Name()[startIndex:len(t2.Name())-1])
		}
	} else {
		return nil, fmt.Errorf("%s does not exist", name)
	}
}

// Must get channel
func MustGetCh[T any](name string) chan T {
	if c, err := GetCh[T](name); err != nil {
		panic(err)
	} else {
		return c
	}
}

// Delete the channel[T] with the name.
// If the channel does not exist, will create the channel[T] with size.
// Return error if the [T] is not same as the recorded [T] or the name does not exist.
func DelCh[T any](name string) error {
	channelsMutex.Lock()
	defer channelsMutex.Unlock()

	if _, ok := channels[name]; ok {
		t := reflect.TypeOf(channels[name])
		t2 := reflect.TypeOf(channel_t[T]{})
		if t.Name() == t2.Name() {
			close(channels[name].(channel_t[T]).C)
			delete(channels, name)
			return nil
		} else {
			startIndex := len("channel_t[")
			return fmt.Errorf("Required T is \"%s\" but get \"%s\"", t.Name()[startIndex:len(t.Name())-1], t2.Name()[startIndex:len(t2.Name())-1])
		}
	} else {
		return fmt.Errorf("Channel %s does not exist", name)
	}
}

// Check the channel is exist or not.
func IsChExist(name string) bool {
	channelsMutex.Lock()
	defer channelsMutex.Unlock()

	_, ok := channels[name]
	return ok
}

// Create RR channel
func CreateRRCh[T any](name string, size int) error {
	rrChannelsMutex.Lock()
	defer rrChannelsMutex.Unlock()

	if _, ok := rrChannels[name]; ok {
		return fmt.Errorf("%s exist", name)
	} else {
		rrChannels[name] = rrChannel_t[T]{
			Request:  make(chan T, size),
			Response: make(chan T, size),
		}
		return nil
	}
}

// Get the rrChannel[T] with the name.
// If the RR channel does not exist, will create the rrChannel[T] with size.
// Return error if the [T] is not same as the recorded [T].
func GetRRCh[T any](name string) (chan T, chan T, error) {
	rrChannelsMutex.Lock()
	defer rrChannelsMutex.Unlock()

	if _, ok := rrChannels[name]; ok {
		t := reflect.TypeOf(rrChannels[name])
		t2 := reflect.TypeOf(rrChannel_t[T]{})
		if t.Name() == t2.Name() {
			return rrChannels[name].(rrChannel_t[T]).Request, rrChannels[name].(rrChannel_t[T]).Response, nil
		} else {
			startIndex := strings.Index(t.Name(), "[") + 1
			return nil, nil, fmt.Errorf("Required T is \"%s\" but get \"%s\"", t.Name()[startIndex:len(t.Name())-1], t2.Name()[startIndex:len(t2.Name())-1])
		}
	} else {
		return nil, nil, fmt.Errorf("%s does not exist", name)
	}
}

// Must get RR channel
func MustGetRRCh[T any](name string) (chan T, chan T) {
	if rq, rp, err := GetRRCh[T](name); err != nil {
		panic(err)
	} else {
		return rq, rp
	}
}

// Delete the rrChannel[T] with the name.
// If the RR channel does not exist, will create the rrChannel[T] with size.
// Return error if the [T] is not same as the recorded [T] or the name does not exist.
func DelRRCh[T any](name string) error {
	rrChannelsMutex.Lock()
	defer rrChannelsMutex.Unlock()

	if _, ok := rrChannels[name]; ok {
		t := reflect.TypeOf(rrChannels[name])
		t2 := reflect.TypeOf(rrChannel_t[T]{})
		if t.Name() == t2.Name() {
			close(rrChannels[name].(rrChannel_t[T]).Request)
			close(rrChannels[name].(rrChannel_t[T]).Response)
			delete(rrChannels, name)
			return nil
		} else {
			startIndex := len("rrChannel_t[")
			return fmt.Errorf("Required T is \"%s\" but get \"%s\"", t.Name()[startIndex:len(t.Name())-1], t2.Name()[startIndex:len(t2.Name())-1])
		}
	} else {
		return fmt.Errorf("RR Channel %s does not exist", name)
	}
}

// Check the RR channel is exist or not.
func IsRRChExist(name string) bool {
	rrChannelsMutex.Lock()
	defer rrChannelsMutex.Unlock()

	_, ok := rrChannels[name]
	return ok
}
