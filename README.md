# gatekeeper
Simple validation using code generation. No reflection welcome.

Gatekeeper generates code that validates http request form values.

Validated values are assigned to corresponding struct fields without the use of reflection.

see example/gatekeeper.yaml for an example configuration file

# features
- Validation for 
    - ints
        - min and max
    - strings
        - min length and max length
    - emails
    - options (must be one of { "January", "February", etc }
    - time.Time
        - specify time format (DateOnly, DateTime, etc)
    - UUIDS
    - Regex Matching

    

# roadmap
## features
- Phone Numbers
- File types
- Additional primitives and other custom types (int64, uuid.UUID, etc)
- Middleware provider
- Custom object-level validators
- Fine-grained error handling
- File watching & caching
- Generate tests for sample data

## internals
- Run generation of all objects in parallel using a workgroup


