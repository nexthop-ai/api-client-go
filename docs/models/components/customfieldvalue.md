# CustomFieldValue


## Supported Types

### CustomFieldValueStr

```go
customFieldValue := components.CreateCustomFieldValueCustomFieldValueStr(components.CustomFieldValueStr{/* values here */})
```

### CustomFieldValueHyperlink

```go
customFieldValue := components.CreateCustomFieldValueCustomFieldValueHyperlink(components.CustomFieldValueHyperlink{/* values here */})
```

### CustomFieldValuePerson

```go
customFieldValue := components.CreateCustomFieldValueCustomFieldValuePerson(components.CustomFieldValuePerson{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldValue.Type {
	case components.CustomFieldValueTypeCustomFieldValueStr:
		// customFieldValue.CustomFieldValueStr is populated
	case components.CustomFieldValueTypeCustomFieldValueHyperlink:
		// customFieldValue.CustomFieldValueHyperlink is populated
	case components.CustomFieldValueTypeCustomFieldValuePerson:
		// customFieldValue.CustomFieldValuePerson is populated
}
```
