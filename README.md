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

1- Testing the default reference implementation star-wars schema.
