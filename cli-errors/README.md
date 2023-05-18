Inconsistent errors - reproduction steps:


* run the cli commands to generate go types and bindings
```
thema lineage gen gobindings --pkgname main
thema lineage gen gotypes --pkgname main
```
There is no error output.

Now run the code (`go run .` should suffice). You should see this error:
```
schema 0.1 is not backwards compatible with schema 0.0:
required field is optional in subsumed value: header
value not an instance
```
