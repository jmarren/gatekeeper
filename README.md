# gatekeeper
Simple validation using code generation. No reflection welcome.

Gatekeeper generates code that validates http request form values.

Validated values are assigned to corresponding struct fields without the use of reflection.

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
        

# roadmap
## features
- File types
- Additional primitives and other custom types (int64, uuid.UUID, etc)
- Middleware provider
- Custom object-level validators
- Fine-grained error handling
- File watching & caching
- Generate tests for sample data

## internals
- Run generation of all objects in parallel using a workgroup


