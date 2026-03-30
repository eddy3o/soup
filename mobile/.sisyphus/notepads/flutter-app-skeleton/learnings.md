## 2026-03-30
- shadcn_ui in this project supports `ShadAlert`/`ShadAlert.destructive` directly for reusable error messaging.
- A spinner-specific widget was not exposed in the available package API; `ShadProgress` is the closest shadcn_ui-native loading indicator.
- Keep shared widgets free of Material UI primitives when the feature explicitly requires shadcn-only components.
