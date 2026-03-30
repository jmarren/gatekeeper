# gatekeeper
Simple validation using code generation. No reflection welcome.

Gatekeeper generates code that validates http request form values.

Validated values are assigned to corresponding struct fields without the use of reflection.

see example/gatekeeper.yaml for an example configuration file

# roadmap
## features
x UUIDs 
x Dates 
- Phone Numbers
x Regex
- File types
- Additional primitives and other custom types (int64, uuid.UUID, etc)
- Middleware provider
x Handler generator/decorator
- Custom object-level validators
- Fine-grained error handling
- File watching & caching
- Generate tests for sample data

## internals
- Run generation of all objects in parallel using a workgroup


