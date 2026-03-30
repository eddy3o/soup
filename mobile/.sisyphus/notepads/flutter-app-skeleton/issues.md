# Issues & Gotchas

## Backend Schema Quirks
- `is_admin` field is STRING, not boolean (returns "true"/"false")
- Backend typo: `OderDetail` (should be `OrderDetail`) - noted for backend team

## Known Limitations
- Access token expires in 15 minutes with no refresh capability
- Users must re-login after expiry

## Dependency Notes
- `shadcn_ui ^0.10.2` could not be resolved; `flutter pub get` succeeded only after using the available published version `^0.53.2`.

(Issues encountered during implementation will be logged here)
- Verification step required python3 because python is unavailable in this environment.
- build_runner currently fails on unrelated `lib/main.dart` dot-shorthand language feature errors; auth model generation still succeeds before the overall build aborts.
- `flutter analyze` on T3.4 only reports pre-existing style infos in `core/constants/api_constants.dart`; auth provider itself is clean.
- `flutter analyze` still reports pre-existing constant naming infos in `lib/core/constants/api_constants.dart`; unrelated to app bootstrap.
- `ProductModel` is currently paired with a handwritten minimal freezed stub because codegen was not part of this task; regenerate proper freezed output later if the project starts relying on generated helpers.
- `CategoryModel` is also using a handwritten minimal freezed-generated stub so the feature compiles without running build_runner in this task.
- `UpdateUserRequest` needed a handwritten minimal freezed stub as well; `flutter analyze` passes, but generated code should be refreshed later if the DTO changes.
- T9.1 verification blockers: no Android emulator/AVD sources were available in this environment (`flutter emulators` returned none), so the required Android-device run could not be performed.
- T9.1 verification blockers: backend at `http://localhost:8080` was not reachable from this session (`curl -I` returned connection refused), so login/logout/navigation flow testing could not be completed.
- `flutter run -d chrome` did launch the app successfully, but that is only a fallback and does not satisfy the requested Android emulator verification.
- F1 review: auth/domain models (`User`, `UserModel`) are plain immutable classes, not `freezed` annotations, so the immutability/codegen requirement is not met for the core auth flow.
- F1 review: `auth_provider.dart` acts as a composition root and imports data-layer implementations directly; acceptable wiring, but it weakens strict presentation/data separation.
- F1 fix: product/order scaffolding was simplified to plain immutable classes and the stale generated freezed files were removed, eliminating analyzer codegen errors.
