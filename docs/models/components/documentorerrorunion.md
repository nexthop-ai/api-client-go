# DocumentOrErrorUnion


## Supported Types

### Document

```go
documentOrErrorUnion := components.CreateDocumentOrErrorUnionDocument(components.Document{/* values here */})
```

### DocumentOrError

```go
documentOrErrorUnion := components.CreateDocumentOrErrorUnionDocumentOrError(components.DocumentOrError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentOrErrorUnion.Type {
	case components.DocumentOrErrorUnionTypeDocument:
		// documentOrErrorUnion.Document is populated
	case components.DocumentOrErrorUnionTypeDocumentOrError:
		// documentOrErrorUnion.DocumentOrError is populated
}
```
