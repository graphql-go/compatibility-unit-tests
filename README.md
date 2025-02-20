# Compatibility Unit Tests
Validation compatibility library for comparing GraphQL implementations unit tests results.

## Quick Start

Running the library:

```bash
$ ./bin/start.sh

```

```bash
(•) https://github.com/graphql-go/graphql

(press q to quit)
Enumerating objects: 5158, done.
Counting objects: 100% (96/96), done.
Compressing objects: 100% (58/58), done.
Total 5158 (delta 52), reused 52 (delta 34), pack-reused 5062 (from 3)
2025/02/10 13:31:56 result: &{}
```

## Validation Compatibility 

The library validates compatibility by checking the end results of the unit tests from the choosen graphql implementation
against the `graphql-js` implementation, the following checks are done:

The strategy works in the following steps:

1- Pulls the tests names of the JavaScript implementation, Eg:
```
describe('Execute: Handles execution with a complex schema', () => {
  it('executes using a schema', () => {
    // ...
  }
}
```

The name is converted to Go test name in the following syntax:
```
ExecuteHandlesExecutionWithAComplexSchema_ExecutesUsingASchema
```

### Details

1- Extractor: Pulls the JavaScript names from the unit tests.


