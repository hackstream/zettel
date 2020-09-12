- `zettel init` -> index.md, config.toml
- `zettel new -s $SLUG_NAME` -> slug.md
- `zettel build` -> return []("s")
  os.Filewalk
  read everything into a DS
  - Link struct []
  - frontmatter struct
  map of slugs with metadata extracted from yaml frontmatter as the value
  regex replace the `[[]]` with a link
  
```go
type Post struct {
    Body string
    Metadata struct
    links []Link
}

type Link struct {
    Slug string
    Title string
}
```

1st Pass: Get the links from the post and fill the struct.
2nd Pass: Replace the body with the links (regex replace).
3rd Pass: Convert body from MD to HTML (some parser lib)
4th Pass: Populate the graph with the edges. Iterate on links, and add edges.

---

-> Convert all []Posts/Graphs to HTML templates

