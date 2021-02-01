just exposing Go's [html.UnescapeString](https://golang.org/pkg/html/#UnescapeString)

UnescapeString unescapes entities like `"&lt;"` to become "<". It unescapes a larger range of entities than EscapeString escapes. For example, `"&aacute;"` unescapes to "á", as does `"&#225;"` and `"&#xE1;"`. UnescapeString(EscapeString(s)) == s always holds, but the converse isn't always true.

install with

`go install github.com/dgparker/unescape`

run with

`unescape "&lt;entity1" "entity2&#xE1;"`
