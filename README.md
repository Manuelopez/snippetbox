# snippetbox

## server Mux 
is like the router do not use the deffault is posible riskier
mux := http.NewServeMux()
mux.HandleFunc("route", function)

### types of routes
"/" is a catch all
"/a/b" fixed route

## interfaces example
fmt.Fprintf is takes in a writer interface. We can pass the w http.ResposnseWriter as the first arguments as it fullfil the writer interface needed (implements the Write function)
## inernal folder
internal carries a special meaning and behaviour in go: any packages which live under this directory can only be imported by code inside the parent of the internal directory Ex: internal can only be imported by code inside our snippetbox project directory


# templating
call the base template which calls the other templates inside 
{{template "name of templ" .}} the dot '.' represents any dynamic data you want to pass to the invoked template

# flags
addr := flag.String("addr", ":4000", "HTTP network address")
    - automatically converts the addr flag to string and sets defaul value :4000
    - “go run ./cmd/web -addr=":9999”
- to see all flag list use -flag when runnig app
    -“go run ./cmd/web -help”

# logging
- log, is the standrard go logger, by defualt prefixes the messages with the local date and time and writes them to the sntandard error stream.
- to redirect the stdout and stderr streams to on-disk files when starting the apliction do
    -“go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log”

# sql 
- mysql driver in golang provides a pool of connections
