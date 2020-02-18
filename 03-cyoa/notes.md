
## Notes
 - Using a pointer `&item` means that the changes will carry over and be persistent(?)
 - Use comma, ok idiom on map lookup 
 - Everything is an interface, you want to use interfaces everywhere
 - If you're not exporting a type, dont' return that type. Return whatever interface you expect that type to be. Because it's not going to be clear to end-users what methods are available
 - JSON-to-Go tool: https://mholt.github.io/json-to-go/ (This tool instantly converts JSON into a Go type definition. )

### Functional Options
    - Allows you to pass options to functions without having to pass in a whole bunch of things
    - type HandlerOption func(h *handler)
    - takes a pointer to a handler so it can modify it
    - Allows you to create functions that return a HandlerOption with anything passed in
    - Allows you to set options
        - ```func WithDatabase(username, password string) HandlerOption { ... }`
     - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
        -   Many functions make light work
    - Ok so for functional options, you have a `...opts` of type `Options` which returns a signature `func (h *handler)` 
    - This return signature is what holds the pointer to whatever it is you are passing in (IE your handler in this case). This allows you to modify the thing you are passing in (as options normally help you do )
    - So your functional option can take in any of its own parameters that would help that option do its intended work, but it will always return `func (h *handler)` which would modify that handler instance with the params you passed in.
    -  ..
    -  Name your functional options With<Option> IE WithDatabase() or WithPathFunc() - not a standard just a guideline


#### IE:
    ```go

    type HandlerOption func (h *handler)

    // Use any arbitrary external params you want
    func WithTemplate(t *template.Template) HandlerOption {
        // return HandlerOption(&h)
        // modify shit in the function
        return func(h *handler) {
            h.t = t
        }
    }

    type handler struct {
        t *template.Template
    }

    func NewHandler(t *template.Template, ...opts HandlerOption) {
        h := handler{s, tpl}
        for _, option in range opts {
            //`option` is of type `HandlerOption(&h)` AKA of type `func (h *handler)`
            option(&h)
        }
    }

    func main() {
        tmpl := "First template"
        x := NewHandler(tmpl, WithTemplate("Second Template"))
    }
    ```

#### Compared to using an Options struct:
 - Allows you to Provide options to functions without being forced to pass in a whole bunch of things.
 - Create a struct IE "HandlerOpts" and that would hold all of your options
 - don't do this because it's not too clear.. 


