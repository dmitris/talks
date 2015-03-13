pagelinks, err := scanr.GetLinks(r, result.link, inParams.domains, inParams.pathPrefix)
  if err != nil {
    fmt.Fprintf(os.Stderr,
      "Error in GetLinks for URL %s: %s\n",
      result.link,
      result.err)
    lmap.mu.Lock()
    lmap.links[link] = linkDoneFailure
    lmap.mu.Unlock()
    return result.err
  }
  lmap.mu.Lock()
  lmap.links[link] = linkDoneSuccess
  lmap.mu.Unlock()

  for _, linkNew := range pagelinks {
    // verify that we an parse the URL before adding it
    _, err = url.Parse(linkNew)