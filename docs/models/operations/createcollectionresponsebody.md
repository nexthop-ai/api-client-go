# CreatecollectionResponseBody

OK


## Supported Types

### ResponseBody1

```go
createcollectionResponseBody := operations.CreateCreatecollectionResponseBodyResponseBody1(operations.ResponseBody1{/* values here */})
```

### ResponseBody2

```go
createcollectionResponseBody := operations.CreateCreatecollectionResponseBodyResponseBody2(operations.ResponseBody2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createcollectionResponseBody.Type {
	case operations.CreatecollectionResponseBodyTypeResponseBody1:
		// createcollectionResponseBody.ResponseBody1 is populated
	case operations.CreatecollectionResponseBodyTypeResponseBody2:
		// createcollectionResponseBody.ResponseBody2 is populated
}
```
