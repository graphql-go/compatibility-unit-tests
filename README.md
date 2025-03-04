# Compatibility Unit Tests

Golang CLI library for **validating compatibility** of any GraphQL implementation against the GraphQL reference implementation: [graphql-js](https://github.com/graphql/graphql-js).


Current implementation supports the following GraphQL implementations:
- [https://github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)

## Quick Start

Running the library:

```bash
$ ./bin/start.sh

```

![Screenshot from 2025-03-04 14-00-43](https://github.com/user-attachments/assets/d0040404-3d70-4b4e-9051-f6f520178075)



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

1- `extractor-js`: Pulls the tests names from the **JavaScript GraphQL reference implementation**.

2- `extractor-js`: Formats and saves the tests names into a file named: **unit-tests.txt**.

3- `compatibility-unit-tests/extractor`: Pulls the tests names from **a GraphQL implementation, eg. Golang**.

4- `compatibility-unit-tests/validator`: Compares the tests names from the **JavaScript GraphQL reference implementation** using **unit-tests.txt** against the **GraphQL implementation, eg. Golang**.


### Further Work

There are multiple **GraphQL implementations** in different languages such in Golang: 
- https://github.com/99designs/gqlgen
- https://github.com/graph-gophers/graphql-go

Current work covers compatibility validation in terms of checking the tests names of the GraphQL reference implementation against the choosen implementation.

This could be extended into including the **body and result** of each test and improving the comparison and making a robust and deeper checks.

