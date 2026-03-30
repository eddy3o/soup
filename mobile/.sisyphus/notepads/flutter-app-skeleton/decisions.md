## 2026-03-30
- Kept router configuration in `lib/core/router/app_router.dart` as a single `GoRouter` provider for app-wide reuse.
- Used named routes for `/`, `/login`, and `/home` to keep navigation consistent for future screens.
- Added `lib/app.dart` as the single root widget entry point, with `main.dart` reduced to bootstrapping `ProviderScope` + `MyApp`.
- Used Shadcn theme primitives (`ShadThemeData` / `ShadZincColorScheme`) rather than hardcoded colors.
- For category scaffolding, kept the model intentionally minimal (`id`, `name`, `description`) to match the backend contract exactly.
- For user profile scaffolding, kept datasource/repository/usecase layers as TODO-only shells and reused the existing auth-domain `User` entity for the contract.
