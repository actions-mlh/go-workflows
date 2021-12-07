Goal: Is to have a parser that can easily integrate with other applications that need a golang representation of of a yaml file.

Architectural decisions:
    1) yaml file is a collection of key, value (How can we easily reference each key:value ?) 
        a) Graph representation of the yaml file: by binding each value to a node and checking for unmarshalling errors i.e, (language specific type errors), then appending it to a branch
        b) One example, is cyclic dependencies 
    2) levels of errors 
        a) i.e, language specific errors, unmarshalling errors at the parsing level
        b) linting errors, specific to the application you are building
    3) how can we be application agnostic
        a) differnet applications can use this, such as the linter, syntax/expressions checker, lsp
    4) code gen (things that can be added)
        a) instead of creating a json schema first, think of parsing rules then create json schema, because it was difficult to generate with public json schema

go run cmd/parser/main.go -i yaml/dirty/circular_dependencies.yaml -d

Goal: Is to have a language server that communicates 
