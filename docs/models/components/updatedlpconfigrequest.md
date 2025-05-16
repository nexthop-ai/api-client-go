# UpdateDlpConfigRequest


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Config`                                                                        | [*components.DlpConfig](../../models/components/dlpconfig.md)                   | :heavy_minus_sign:                                                              | Detailed configuration of what documents and sensitive content will be scanned. |
| `Frequency`                                                                     | **string*                                                                       | :heavy_minus_sign:                                                              | Only "ONCE" is supported for reports.                                           |