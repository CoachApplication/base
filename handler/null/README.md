# Null handler

This handler provides Null Operations which can be used to fill in missing 
required Operation structs, with objects that introduce no actual 
functionality.

This handler has a builder which can be used to provide any Operation struct.
The builder can be used as an early API builder in the list, which would 
allow any other builders to override it's provided Operation structs.

## Operations

Any Null Operation should be a valid operation, but should return a failed
api.Result when Executed.
