# gatekeeper
Simple validation using code generation. No reflection welcome.

Gatekeeper generates code that validates http request form values.

Validated values are assigned to corresponding struct fields without the use of reflection.

To generate a .gatekeeper.go file, run gatekeeper on a .yaml configuration file.

see example/gatekeeper.yaml for an example configuration file

# features
- Validation for 
    - Ints
        - min and max
    - Strings
        - min length and max length
    - Emails
    - Options (must be one of { "January", "February", etc }
    - time.Time
        - specify time format (DateOnly, DateTime, etc)
    - UUIDS
    - Regex Matching
    - Phone Numbers
        - provided as *libphonenumber.PhoneNumber validated against a region
   
- Handler Decorator
    Functionality is provided to convert a handler that takes in the request, responseWriter,  a pointer to the validated struct, and a *gkerror.ValidationErrGroup to an http.HandlerFunc by decorating it to create and validate the struct, then pass it to your function.
    All you need to do is create a function like:
    ```go
        func UserHandler(w http.ResponseWriter, r *http.Request, User *User, errs *gkerror.ValidationErrGroup) {
            // do something with the struct and errors        
            if errs.Any() {
                // an error occurred
                firstNameErrs := errs.ByField("FirstName") 
                // do something with errors on the FirstName field
            }
        }
    ```
    then you can pass it to the generated function decorator to get an http.HandlerFunc:
    ```go
        handler := NewUserHandler(UserHandler) // this returns an http.HandlerFunc that can be assigned to a route
    ```

# roadmap
## features to implement going forward
- File types
- Additional primitives and other custom types (booleans, slices, int64, uuid.UUID, etc)
- Middleware provider
- Custom object-level validators
- Fine-grained error handling
- File watching & caching
- Generate tests for sample data
- Additional string features:
    - hasPrefix
    - hasSuffix
    - includes
    - doesn't include
    - alphaNumeric
- Additional Int features:
    - isPositive
    - isNegative
    - isMultipleOf
- optional fields with default values
- URLs
- IP addresses
- Mac Addresses
- Emojis
- Additional UUID versions uuidv4, uuidv6, etc
- JWTs
- Parsing from request body as JSON
- Panic on Error option w/ recover handler
- Validation escape on first failure


## internals
- Run generation of all objects in parallel using a workgroup


