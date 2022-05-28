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

var (
	channels = map[string]interface{}{}
	mutex    = sync.Mutex{}
)

// Get the channel[T] with the name.
// If the channel does not exist, will create the channel[T] with size.
// Return error if the [T] is not same as the recorded [T].
func GetCh[T any](name string, size int) (chan T, error) {
	mutex.Lock()
	defer mutex.Unlock()

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
		ch := make(chan T, size)
		channelInfo := channel_t[T]{
			C: ch,
		}
		channels[name] = channelInfo
		return channelInfo.C, nil
	}
}

// Delete the channel[T] with the name.
// If the channel does not exist, will create the channel[T] with size.
// Return error if the [T] is not same as the recorded [T] or the name does not exist.
func DelCh[T any](name string) error {
	mutex.Lock()
	defer mutex.Unlock()

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
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := channels[name]
	return ok
}
