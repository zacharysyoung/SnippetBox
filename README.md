# Let's Go: SnippetBox

## Ch 2

- Initialized Go module, Git

- From "Header canonicalization", HTTP/1 allows for bumpy-case header keys, and Go wants to follow this convention (CanonicalMIMEHeaderKey()).  To get past this update the header map directly (e.g., `w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}`, otherwise Go would convert to X-Xss-Protection).

    HTTP/2 requires the server respond will all lowercase, and Go enforces this regardless of direct manipulation.

- Handlers satisfy the http.Handler interface, which has the sole method ServeHTTP(ResponseWriter, *Request).

    ```go
    type home struct {}

    func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        ...
    }

    ...

    mux := http.NewServerMux()
    mux.Handle("/", &home{})
    ```

    is equivalent to what we currently have:

    ```go
    mux := http.NewServerMux()
    mux.HandleFunc("/", home)
    ```

    where home is an actual Handler function—internally, mux.HandleFunc just wraps home, like HandlerFunc(home), to generate something (?) with the ServeHTTP method.

## Ch 3

- Create an application to hold info and error loggers (I think it will soon hold DB info).  Refactored all functions to be methods of application so they have access to the loggers, and also just to "bring it all together":

    ```go
    type application struct {
        errorLog: *log.Logger
        infoLog:  *log.Logger
    }
    ```

    then in helpers,

    ```go
    func (app *application) serverError(w http.ResponseWriter, err error) {
        app.errorLog(...)

        httpErr=http.StatusInternalServerError
        http.Error(w, http.StatusText(httpErr), httpErr)
    }
    ```

    so that in handlers,

    ```go
    func (app *application) pathHandler(w http.ResponseWriter, w *http.Request) {
        if err := someInternalCheck(); err != nil {
            app.serverError(w, err)
            return
        }

        ...
    }

## Ch 4

- Go's sql driver doesn't work with the broader concept of NULL often used in db engines.  Be sure to set NOT DEFAULT constraints when defining the table/columns, like we did earlier:

    ```sql
    CREATE TABLE snippets (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(100) NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL,
        expires DATETIME NOT NULL
    );
    ```

    and provide sensible DEFAULT values as necessary.

## Ch 5

- Check git log for "before the template in-memory cache change", on Fri Jul 14 13:57:19 2023 -0700.
