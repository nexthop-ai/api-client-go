# ChatFileFailureReason

Reason for failed status.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ChatFileFailureReasonParseFailed

// Open enum: custom values can be created with a direct type cast
custom := components.ChatFileFailureReason("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ChatFileFailureReasonParseFailed`                | PARSE_FAILED                                      |
| `ChatFileFailureReasonAvScanFailed`               | AV_SCAN_FAILED                                    |
| `ChatFileFailureReasonFileTooSmall`               | FILE_TOO_SMALL                                    |
| `ChatFileFailureReasonFileTooLarge`               | FILE_TOO_LARGE                                    |
| `ChatFileFailureReasonFileExtensionUnsupported`   | FILE_EXTENSION_UNSUPPORTED                        |
| `ChatFileFailureReasonFileMetadataValidationFail` | FILE_METADATA_VALIDATION_FAIL                     |
| `ChatFileFailureReasonFileProcessingTimedOut`     | FILE_PROCESSING_TIMED_OUT                         |
| `ChatFileFailureReasonOauthNeeded`                | OAUTH_NEEDED                                      |
| `ChatFileFailureReasonURLFetchFailed`             | URL_FETCH_FAILED                                  |
| `ChatFileFailureReasonEmptyContent`               | EMPTY_CONTENT                                     |
| `ChatFileFailureReasonAuthRequired`               | AUTH_REQUIRED                                     |