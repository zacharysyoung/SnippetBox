# Let's Go: SnippetBox

## Ch 2

- Initialized Go module, Git

- From "Header canonicalization", HTTP/1 allows for bumpy-case header keys, and Go wants to follow this convention (CanonicalMIMEHeaderKey()).  To get past this update the header map directly (e.g., `w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}`, otherwise Go would convert to X-Xss-Protection).

    HTTP/2 requires the server respond will all lowercase, and Go enforces this regardless of direct manipulation.
