# DocumentSpecUnion


## Supported Types

### DocumentSpec1

```go
documentSpecUnion := components.CreateDocumentSpecUnionDocumentSpec1(components.DocumentSpec1{/* values here */})
```

### DocumentSpec2

```go
documentSpecUnion := components.CreateDocumentSpecUnionDocumentSpec2(components.DocumentSpec2{/* values here */})
```

### DocumentSpec3

```go
documentSpecUnion := components.CreateDocumentSpecUnionDocumentSpec3(components.DocumentSpec3{/* values here */})
```

### DocumentSpec4

```go
documentSpecUnion := components.CreateDocumentSpecUnionDocumentSpec4(components.DocumentSpec4{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch documentSpecUnion.Type {
	case components.DocumentSpecUnionTypeDocumentSpec1:
		// documentSpecUnion.DocumentSpec1 is populated
	case components.DocumentSpecUnionTypeDocumentSpec2:
		// documentSpecUnion.DocumentSpec2 is populated
	case components.DocumentSpecUnionTypeDocumentSpec3:
		// documentSpecUnion.DocumentSpec3 is populated
	case components.DocumentSpecUnionTypeDocumentSpec4:
		// documentSpecUnion.DocumentSpec4 is populated
}
```
