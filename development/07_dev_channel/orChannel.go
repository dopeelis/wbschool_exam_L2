package main

func or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})

	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				done <- nil
				return
			case <-done:
				return
			}
		}(ch)
	}
	return done
}
