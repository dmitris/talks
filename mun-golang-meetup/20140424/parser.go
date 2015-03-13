  for {
    tokenType := parser.Next()
    if tokenType == html.ErrorToken {
      break
    }
    token := parser.Token()
    urlFull = nil
    // fmt.Printf("DMDEBUG parser loop tokenType: %v, token: %v\n", tokenType, token)
    switch tokenType {
    case html.SelfClosingTagToken:
      if token.DataAtom == atom.Base {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            if attr.Val != "" {
              base_href = attr.Val
              baseUrl, err = url.Parse(base_href)