  pace := make(chan bool, jobs)
  // load the synchronizing channel with max number of concurrent fetchers

  for i := 0; i < jobs; i++ {
    pace <- true
  }