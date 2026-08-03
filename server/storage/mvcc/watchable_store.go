// ... existing code ...

func (s *watchableStore) syncWatchersLoop() {
	for {
		s.mu.Lock()
		// Ensure we check the current revision before moving watchers to synced
		curRev := s.rev
		for w := range s.unsynced {
			if w.minRev <= curRev {
				// Move to synced only if we are sure we can catch up
				s.synced.add(w)
				delete(s.unsynced, w)
			}
		}
		s.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

// ... existing code ...