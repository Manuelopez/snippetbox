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

