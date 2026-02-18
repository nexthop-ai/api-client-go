# GetShortcutRequestUnion


## Supported Types

### UserGeneratedContentID

```go
getShortcutRequestUnion := components.CreateGetShortcutRequestUnionUserGeneratedContentID(components.UserGeneratedContentID{/* values here */})
```

### GetShortcutRequest

```go
getShortcutRequestUnion := components.CreateGetShortcutRequestUnionGetShortcutRequest(components.GetShortcutRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getShortcutRequestUnion.Type {
	case components.GetShortcutRequestUnionTypeUserGeneratedContentID:
		// getShortcutRequestUnion.UserGeneratedContentID is populated
	case components.GetShortcutRequestUnionTypeGetShortcutRequest:
		// getShortcutRequestUnion.GetShortcutRequest is populated
}
```
