# Work Plan: Flutter App Skeleton with Login/Home/Logout

## Overview
Build a Flutter mobile app with login, home (welcome screen with user info), and logout functionality. Use shadcn_ui for all UI components, go_router for navigation, Provider/Riverpod for state management, and flutter_secure_storage for token persistence. Create foundational repository structure ready for future features (products, orders, categories).

## Project Scope

### INCLUDE
- Login screen (phone + password)
- Home screen (welcome with user name, phone, logout button)
- Authentication flow with JWT token management (access tokens, NO refresh endpoint available in backend)
- Environment configuration (dev/staging/prod)
- Repository pattern with services for: Auth, User, Products, Orders, Categories
- HTTP client with interceptors (token injection, 401 re-authentication flow)
- Secure token storage (flutter_secure_storage)
- shadcn_ui components for all UI
- go_router navigation with auth guards
- Error handling and loading states
- Clean architecture folder structure

### EXCLUDE
- Product browsing UI (structure only, no screens)
- Order creation UI (structure only, no screens)
- User profile editing UI (structure only, no screens)
- Unit/widget/integration tests (structure files can be created but no test code)
- Push notifications
- Image upload functionality
- Form validation beyond basic required fields

## Technical Decisions

### Architecture
- **Pattern**: Clean Architecture (data layer → domain layer → presentation layer)
- **State Management**: Riverpod (simpler than Provider, better for async)
- **Navigation**: go_router with redirect guards for authentication
- **Dependency Injection**: Riverpod providers

### Folder Structure
```
lib/
├── main.dart
├── app.dart                          # Root app widget with providers
├── core/
│   ├── config/
│   │   └── environment.dart          # Environment configuration (dev/staging/prod)
│   ├── constants/
│   │   ├── api_constants.dart        # API endpoints
│   │   └── storage_keys.dart         # Secure storage keys
│   ├── errors/
│   │   ├── exceptions.dart           # Custom exceptions
│   │   └── failures.dart             # Failure classes for error handling
│   ├── network/
│   │   ├── api_client.dart           # Dio HTTP client setup
│   │   ├── interceptors/
│   │   │   ├── auth_interceptor.dart # Token injection
│   │   │   └── error_interceptor.dart # Global error handling
│   │   └── network_info.dart         # Connectivity checker (optional)
│   ├── router/
│   │   └── app_router.dart           # go_router configuration with guards
│   └── utils/
│       └── logger.dart                # Logging utility
├── features/
│   ├── auth/
│   │   ├── data/
│   │   │   ├── models/
│   │   │   │   ├── user_model.dart           # User JSON model
│   │   │   │   ├── login_request.dart        # Login request DTO
│   │   │   │   └── login_response.dart       # Login response DTO
│   │   │   ├── repositories/
│   │   │   │   └── auth_repository_impl.dart # Auth repo implementation
│   │   │   └── datasources/
│   │   │       ├── auth_remote_datasource.dart # API calls
│   │   │       └── auth_local_datasource.dart  # Token storage
│   │   ├── domain/
│   │   │   ├── entities/
│   │   │   │   └── user.dart                 # User entity
│   │   │   ├── repositories/
│   │   │   │   └── auth_repository.dart      # Auth repo interface
│   │   │   └── usecases/
│   │   │       ├── login_usecase.dart
│   │   │       ├── logout_usecase.dart
│   │   │       ├── get_current_user_usecase.dart
│   │   │       └── check_auth_status_usecase.dart
│   │   └── presentation/
│   │       ├── providers/
│   │       │   └── auth_provider.dart        # Riverpod auth state
│   │       ├── screens/
│   │       │   └── login_screen.dart         # Login UI
│   │       └── widgets/
│   │           ├── login_form.dart           # Login form with shadcn_ui
│   │           └── auth_error_widget.dart    # Error display
│   ├── home/
│   │   └── presentation/
│   │       ├── screens/
│   │       │   └── home_screen.dart          # Home welcome screen
│   │       └── widgets/
│   │           ├── user_info_card.dart       # User info display
│   │           └── logout_button.dart        # Logout button
│   ├── products/
│   │   ├── data/
│   │   │   ├── models/
│   │   │   │   ├── product_model.dart
│   │   │   │   └── pagination_model.dart
│   │   │   ├── repositories/
│   │   │   │   └── product_repository_impl.dart
│   │   │   └── datasources/
│   │   │       └── product_remote_datasource.dart
│   │   └── domain/
│   │       ├── entities/
│   │       │   ├── product.dart
│   │       │   └── paginated_result.dart
│   │       └── repositories/
│   │           └── product_repository.dart
│   ├── orders/
│   │   ├── data/
│   │   │   ├── models/
│   │   │   │   ├── order_model.dart
│   │   │   │   └── order_detail_model.dart
│   │   │   ├── repositories/
│   │   │   │   └── order_repository_impl.dart
│   │   │   └── datasources/
│   │   │       └── order_remote_datasource.dart
│   │   └── domain/
│   │       ├── entities/
│   │       │   ├── order.dart
│   │       │   └── order_detail.dart
│   │       └── repositories/
│   │           └── order_repository.dart
│   ├── categories/
│   │   ├── data/
│   │   │   ├── models/
│   │   │   │   └── category_model.dart
│   │   │   ├── repositories/
│   │   │   │   └── category_repository_impl.dart
│   │   │   └── datasources/
│   │   │       └── category_remote_datasource.dart
│   │   └── domain/
│   │       ├── entities/
│   │       │   └── category.dart
│   │       └── repositories/
│   │           └── category_repository.dart
│   └── user/
│       ├── data/
│       │   ├── models/
│       │   │   └── update_user_request.dart
│       │   ├── repositories/
│       │   │   └── user_repository_impl.dart
│       │   └── datasources/
│       │       └── user_remote_datasource.dart
│       └── domain/
│           ├── repositories/
│           │   └── user_repository.dart
│           └── usecases/
│               ├── get_user_profile_usecase.dart
│               └── update_user_profile_usecase.dart
└── shared/
    └── widgets/
        ├── loading_indicator.dart     # shadcn_ui loading spinner
        └── error_message.dart         # shadcn_ui error display
```

### Dependencies
```yaml
dependencies:
  flutter:
    sdk: flutter
  
  # UI
  shadcn_ui: ^0.10.2                    # shadcn_ui components
  
  # State Management
  flutter_riverpod: ^2.6.1              # State management
  
  # Routing
  go_router: ^14.6.2                    # Navigation
  
  # HTTP
  dio: ^5.7.0                            # HTTP client
  
  # Storage
  flutter_secure_storage: ^9.2.2       # Secure token storage
  
  # Code Generation
  freezed_annotation: ^2.4.4            # Immutable models
  json_annotation: ^4.9.0               # JSON serialization
  
  # Utilities
  logger: ^2.5.0                         # Logging
  intl: ^0.20.1                          # Date formatting

dev_dependencies:
  flutter_test:
    sdk: flutter
  flutter_lints: ^6.0.0
  
  # Code Generation
  build_runner: ^2.4.13
  freezed: ^2.5.7
  json_serializable: ^6.8.0
```

## Backend Integration Specs

### Base URLs
- **Development**: `http://localhost:8080`
- **Staging**: `https://staging-api.soup.com` (placeholder)
- **Production**: `https://api.soup.com` (placeholder)

### Authentication Flow (JWT - Mobile Approach)
1. User enters phone + password
2. POST to `/auth/login` with credentials
3. Backend returns: `{access_token, refresh_token, user}`
4. Store both tokens in flutter_secure_storage
5. Set up Dio to include Authorization header: `Bearer {access_token}` on all authenticated requests
6. **IMPORTANT**: Backend has NO `/auth/refresh` endpoint. Access token expires in 15 minutes.
7. On 401 Unauthorized: Clear tokens and redirect to login (force re-authentication)
8. Logout: POST to `/auth/logout` + clear local storage

**Note**: Mobile apps use Authorization headers (not cookies). Backend middleware checks Authorization header as fallback when cookie is missing, which is perfect for mobile.

### Error Handling
- 401 Unauthorized → Clear tokens, redirect to login (no refresh endpoint available)
- 400 Bad Request → Display error message from backend
- 500 Server Error → Show generic error message
- Network errors → Show "No connection" message

## TODOs

### Phase 1: Project Setup & Dependencies
- [x] **T1.1**: Update pubspec.yaml with all dependencies
  - Add shadcn_ui, flutter_riverpod, go_router, dio, flutter_secure_storage
  - Add dev dependencies: build_runner, freezed, json_serializable
  - Run `flutter pub get`
  
  **Acceptance Criteria:**
  - pubspec.yaml contains all listed dependencies with correct versions
  - `flutter pub get` completes without errors
  - No version conflicts
  
  **Parallelizable:** No (foundation task)

- [x] **T1.2**: Create complete folder structure
  - Create all directories as per the structure above
  - Create empty `.gitkeep` files in empty directories
  - Verify structure matches the design
  
  **Acceptance Criteria:**
  - All folders exist in `lib/`
  - Structure matches the architecture diagram exactly
  - Empty directories have `.gitkeep` files
  
  **Parallelizable:** No (depends on T1.1)

### Phase 2: Core Infrastructure
- [x] **T2.1**: Implement environment configuration
  - Create `core/config/environment.dart` with Environment enum (dev, staging, prod)
  - Add `getBaseUrl()` method returning correct URL per environment
  - Add `getCurrentEnvironment()` method (reads from `--dart-define` or defaults to dev)
  - Create `core/constants/api_constants.dart` with all endpoint paths
  
  **Acceptance Criteria:**
  - Environment.dart has dev/staging/prod enum values
  - Base URLs defined for each environment
  - API constants file has all backend endpoints as constants
  - Can switch environment via `--dart-define=ENV=prod`
  
  **Parallelizable:** No (foundation for network layer)

- [x] **T2.2**: Implement secure storage service
  - Create `core/constants/storage_keys.dart` with key constants (accessToken, refreshToken, userId)
  - Create `features/auth/data/datasources/auth_local_datasource.dart`
  - Implement: `saveTokens()`, `getAccessToken()`, `getRefreshToken()`, `clearTokens()`
  - Use flutter_secure_storage
  
  **Acceptance Criteria:**
  - Storage keys defined as constants (no magic strings)
  - All storage methods implemented with proper error handling
  - Tokens stored with flutter_secure_storage
  - clearTokens() removes all auth-related data
  
  **Parallelizable:** Yes (independent from T2.1)

- [x] **T2.3**: Set up Dio HTTP client with interceptors
  - Create `core/network/api_client.dart` with Dio singleton
  - Configure base URL from environment
  - Create `core/network/interceptors/auth_interceptor.dart`:
    - Read access_token from secure storage
    - Add `Authorization: Bearer {token}` header to all authenticated requests
  - Create `core/network/interceptors/error_interceptor.dart`:
    - Handle 401 → clear tokens and trigger logout (no refresh endpoint available)
    - Handle 400, 500, network errors with user-friendly messages
  - Add logging interceptor for debugging
  
  **Acceptance Criteria:**
  - Dio instance configured with base URL from environment
  - Auth interceptor adds Bearer token to all requests requiring auth
  - Error interceptor handles 401 by clearing storage and redirecting to login
  - Error interceptor handles 400, 500, network errors with appropriate messages
  - Logging shows request/response details in debug mode
  
  **QA Scenarios:**
  - **Tool**: Dio logging + Manual test
  - **Steps**: 
    1. Configure Dio with logging interceptor
    2. Make authenticated request with valid token → verify logs show `Authorization: Bearer {token}` header
    3. Make request with expired/invalid token → verify 401 triggers token clear + redirect
    4. Simulate network error → verify error message displayed
  - **Expected**: All interceptors work correctly, 401 clears auth state
  
  **Parallelizable:** No (depends on T2.1, T2.2)

- [x] **T2.4**: Create error handling infrastructure
  - Create `core/errors/exceptions.dart` with custom exceptions:
    - ServerException, CacheException, NetworkException, UnauthorizedException
  - Create `core/errors/failures.dart` with failure classes:
    - ServerFailure, CacheFailure, NetworkFailure, UnauthorizedFailure
  - Add utility methods to convert exceptions to failures
  
  **Acceptance Criteria:**
  - All exception types defined with descriptive messages
  - Failure classes are immutable with message property
  - Exception → Failure conversion utility exists
  
  **Parallelizable:** Yes (independent)

### Phase 3: Authentication Feature
- [x] **T3.1**: Create auth data models
  - Create `features/auth/data/models/user_model.dart` with fromJson/toJson:
    - Fields: id, phone, name, address, email, photoUrl, isAdmin (String from backend), pushToken, createdAt
    - Use freezed for immutability
  - Create `features/auth/data/models/login_request.dart`:
    - Fields: phone (required), password (required)
  - Create `features/auth/data/models/login_response.dart`:
    - Fields: accessToken, refreshToken, user (UserModel)
  - Run build_runner to generate code
  
  **Acceptance Criteria:**
  - UserModel matches backend response schema exactly
  - **isAdmin** field is **String** (backend returns "true"/"false" as string, not boolean)
  - All models use freezed with @freezed annotation
  - fromJson/toJson methods generated correctly
  - Field names use camelCase (with @JsonKey for snake_case backend)
  - `build_runner build` completes successfully
  
  **QA Scenarios:**
  - **Tool**: build_runner + Manual verification
  - **Steps**:
    1. Run `flutter pub run build_runner build`
    2. Read generated .freezed.dart and .g.dart files
    3. Test UserModel.fromJson with sample backend response: `{"is_admin": "true"}`
  - **Expected**: No build errors, isAdmin correctly typed as String
  
  **Parallelizable:** Yes (can start while core is being built)

- [x] **T3.2**: Create auth domain layer
  - Create `features/auth/domain/entities/user.dart` (domain entity, not data model)
  - Create `features/auth/domain/repositories/auth_repository.dart` (interface):
    - Methods: `login(phone, password)`, `logout()`, `getCurrentUser()`, `isAuthenticated()`
  - Create `features/auth/domain/usecases/login_usecase.dart`
  - Create `features/auth/domain/usecases/logout_usecase.dart`
  - Create `features/auth/domain/usecases/get_current_user_usecase.dart`
  - Create `features/auth/domain/usecases/check_auth_status_usecase.dart`
  
  **Acceptance Criteria:**
  - User entity is clean domain object (no JSON annotations)
  - Repository interface defines contract (no implementation)
  - Each usecase has single responsibility
  - Usecases return Either<Failure, Success> types (use dartz or Result type)
  
  **Parallelizable:** No (depends on T3.1 for user structure)

- [x] **T3.3**: Implement auth data layer
  - Create `features/auth/data/datasources/auth_remote_datasource.dart`:
    - `login(phone, password)` → POST /auth/login
    - `logout()` → POST /auth/logout
  - Create `features/auth/data/repositories/auth_repository_impl.dart`:
    - Implements auth_repository interface
    - Orchestrates remote + local datasources
    - Converts exceptions to failures
    - Saves tokens to local storage on successful login
    - Clears tokens on logout
  
  **Acceptance Criteria:**
  - Remote datasource makes correct API calls with Dio
  - Login stores both access_token and refresh_token
  - Logout clears secure storage
  - Repository converts UserModel to User entity
  - All exceptions properly caught and converted to failures
  
  **Parallelizable:** No (depends on T2.3 for HTTP client, T3.2 for repository interface)

- [x] **T3.4**: Create auth state management with Riverpod
  - Create `features/auth/presentation/providers/auth_provider.dart`:
    - AuthState with sealed classes: Initial, Loading, Authenticated(User), Unauthenticated, Error
    - AuthNotifier with methods: `login(phone, password)`, `logout()`, `checkAuthStatus()`
  - Create provider for current user (derived from auth state)
  - Create provider for auth status (boolean)
  
  **Acceptance Criteria:**
  - AuthState covers all possible auth states
  - AuthNotifier properly calls usecases
  - State transitions handled correctly (loading → success/error)
  - CurrentUser provider returns null when unauthenticated
  - AuthStatus provider returns true only when authenticated
  
  **Parallelizable:** No (depends on T3.3 for repository implementation)

- [x] **T3.5**: Build login screen UI with shadcn_ui
  - Create `features/auth/presentation/screens/login_screen.dart`
  - Create `features/auth/presentation/widgets/login_form.dart`:
    - Phone input field (shadcn_ui TextField)
    - Password input field (shadcn_ui TextField with obscure text)
    - Login button (shadcn_ui Button)
    - Loading indicator (shadcn_ui Spinner when loading)
    - Error message display (shadcn_ui Alert)
  - Wire up form to auth provider
  - Handle loading states and errors
  
  **Acceptance Criteria:**
  - Login screen uses only shadcn_ui components
  - Form validates required fields (phone min 6 chars, password min 6 chars)
  - Login button disabled during loading
  - Error messages displayed using shadcn_ui Alert component
  - Successful login navigates to home screen
  - UI follows shadcn design system
  
  **Parallelizable:** No (depends on T3.4 for auth provider)

### Phase 4: Home Feature
- [x] **T4.1**: Build home screen UI with shadcn_ui
  - Create `features/home/presentation/screens/home_screen.dart`
  - Create `features/home/presentation/widgets/user_info_card.dart`:
    - Display user name (or "User" if null)
    - Display user phone
    - Use shadcn_ui Card component
  - Create `features/home/presentation/widgets/logout_button.dart`:
    - shadcn_ui Button with logout action
    - Shows confirmation dialog (shadcn_ui AlertDialog) before logout
  - Wire up to auth provider for user data and logout
  
  **Acceptance Criteria:**
  - Home screen displays "Welcome, {name}!" header
  - User info card shows name and phone number
  - If name is null, displays phone number only
  - Logout button triggers confirmation dialog
  - Confirming logout calls auth provider logout method
  - Successful logout redirects to login screen
  - All components use shadcn_ui
  
  **Parallelizable:** No (depends on T3.4 for auth provider)

### Phase 5: Navigation & Routing
- [x] **T5.1**: Configure go_router with auth guards
  - Create `core/router/app_router.dart`:
    - Define routes: `/login`, `/home`
    - Implement redirect logic: unauthenticated users → /login
    - Authenticated users accessing /login → /home
  - Add route transitions
  - Set initial route based on auth status
  
  **Acceptance Criteria:**
  - GoRouter configured with all routes
  - Redirect logic prevents unauthorized access to /home
  - Authenticated users automatically redirected away from /login
  - Initial route determined by checking auth status provider
  - Route transitions are smooth
  
  **Parallelizable:** No (depends on T3.4 for auth state, T3.5 for login screen, T4.1 for home screen)

### Phase 6: App Integration
- [x] **T6.1**: Create root app widget
  - Create `lib/app.dart`:
    - Wrap with ProviderScope (Riverpod)
    - Configure MaterialApp.router with go_router
    - Set up shadcn_ui theme
    - Add app-level error handling
  - Update `lib/main.dart`:
    - Initialize flutter_secure_storage
    - Set up logging
    - Run app with environment configuration
  
  **Acceptance Criteria:**
  - App initializes without errors
  - Riverpod providers available throughout app
  - Router navigation works correctly
  - shadcn_ui theme applied globally
  - App starts on correct initial route based on auth status
  
  **Parallelizable:** No (depends on T5.1 for router)

### Phase 7: Future Feature Foundations (Structure Only)
- [x] **T7.1**: Create product feature structure
  - Create all product feature folders (data/models, data/repositories, data/datasources, domain/entities, domain/repositories)
  - Create `features/products/data/models/product_model.dart` with schema from backend
  - Create `features/products/data/models/pagination_model.dart`
  - Create `features/products/domain/entities/product.dart`
  - Create `features/products/domain/repositories/product_repository.dart` interface
  - Create `features/products/data/datasources/product_remote_datasource.dart` with method signatures (no implementation)
  - Add TODO comments for future implementation
  
  **Acceptance Criteria:**
  - All folder structure exists
  - Models match backend schemas exactly
  - Repository interface defines methods: `getProducts(page, limit)`, `getProductById(id)`
  - Datasource has method signatures with TODO comments
  - No actual UI or business logic implemented
  
  **Parallelizable:** Yes (independent of main flow)

- [x] **T7.2**: Create order feature structure
  - Create all order feature folders
  - Create `features/orders/data/models/order_model.dart` with backend schema
  - Create `features/orders/data/models/order_detail_model.dart`
  - Create `features/orders/domain/entities/order.dart`
  - Create `features/orders/domain/entities/order_detail.dart`
  - Create `features/orders/domain/repositories/order_repository.dart` interface
  - Create `features/orders/data/datasources/order_remote_datasource.dart` with method signatures
  - Add TODO comments for future implementation
  
  **Acceptance Criteria:**
  - All folder structure exists
  - Models match backend schemas (pickup_date, general_notes, order_details array)
  - Repository interface defines `createOrder(order)` method
  - Datasource has method signatures with TODO comments
  - No actual UI or business logic implemented
  
  **Parallelizable:** Yes (independent of main flow)

- [x] **T7.3**: Create category feature structure
  - Create all category feature folders
  - Create `features/categories/data/models/category_model.dart`
  - Create `features/categories/domain/entities/category.dart`
  - Create `features/categories/domain/repositories/category_repository.dart` interface
  - Create `features/categories/data/datasources/category_remote_datasource.dart` with method signatures
  - Add TODO comments for future implementation
  
  **Acceptance Criteria:**
  - All folder structure exists
  - Models match backend schema (id, name, description)
  - Repository interface defines `getCategories()` method
  - Datasource has method signatures with TODO comments
  
  **Parallelizable:** Yes (can run with T7.1, T7.2)

- [x] **T7.4**: Create user profile feature structure
  - Create user feature folders (data/models, data/repositories, data/datasources, domain/repositories, domain/usecases)
  - Create `features/user/data/models/update_user_request.dart` (name, address, email, photo_url)
  - Create `features/user/domain/repositories/user_repository.dart` interface
  - Create `features/user/data/datasources/user_remote_datasource.dart` with method signatures
  - Methods: `getUserProfile()` (GET /users/me), `updateUserProfile(data)` (PATCH /users/me)
  - Add TODO comments for future implementation
  
  **Acceptance Criteria:**
  - All folder structure exists
  - UpdateUserRequest model matches PATCH endpoint schema
  - Repository interface defines user profile methods
  - Datasource has method signatures with TODO comments
  
  **Parallelizable:** Yes (can run with T7.1, T7.2, T7.3)

### Phase 8: Shared Components
- [x] **T8.1**: Create shared UI components
  - Create `shared/widgets/loading_indicator.dart` using shadcn_ui Spinner
  - Create `shared/widgets/error_message.dart` using shadcn_ui Alert
  - Make components reusable across features
  
  **Acceptance Criteria:**
  - Loading indicator uses shadcn_ui component
  - Error message accepts custom text and severity
  - Both widgets properly styled with shadcn theme
  
  **Parallelizable:** Yes (can run anytime)

### Phase 9: Testing & Verification
- [x] **T9.1**: Verify app builds and runs
  - Run `flutter clean`
  - Run `flutter pub get`
  - Run `build_runner build`
  - Run `flutter run` on Android emulator
  - Test login flow with real backend
  - Test logout flow
  - Verify token storage and retrieval
  - Test navigation guards
  - Test token refresh logic (wait 15 minutes or manipulate expiry)
  
  **Acceptance Criteria:**
  - App builds without errors
  - Login successfully authenticates with backend
  - Home screen displays user information correctly
  - Logout clears tokens and redirects to login
  - Auth guards prevent unauthorized access
  - Token refresh works automatically
  - No runtime errors during normal flow
  
  **Parallelizable:** No (final verification of everything)

## Final Verification Wave

### F1: Code Quality Review
**Goal:** Ensure code follows Flutter best practices and clean architecture principles

**Checklist:**
- [ ] All files follow Dart style guide (use `flutter analyze`)
- [ ] No linting errors or warnings
- [ ] Clean architecture layers properly separated (data ↔ domain ↔ presentation)
- [ ] No business logic in UI widgets
- [ ] No direct API calls from presentation layer
- [ ] All models use immutable data structures (freezed)
- [ ] Proper error handling in all layers

**Verification Command:**
```bash
flutter analyze
```

**Expected Result:** 0 issues found

---

### F2: Authentication Flow Verification
**Goal:** Verify complete authentication flow works end-to-end

**Checklist:**
- [ ] Login with valid credentials succeeds
- [ ] Login with invalid credentials shows error
- [ ] Access token stored in flutter_secure_storage
- [ ] Refresh token stored in flutter_secure_storage
- [ ] Bearer token automatically added to API requests
- [ ] Cookies properly handled by dio_cookie_manager
- [ ] Home screen receives user data after login
- [ ] Logout clears all tokens
- [ ] Logout redirects to login screen
- [ ] Cannot access home screen without authentication
- [ ] Authenticated user redirected away from login screen

**Manual Test Steps:**
1. Start app → should land on login screen
2. Enter valid phone + password → should navigate to home
3. Verify user info displayed correctly
4. Close and reopen app → should stay logged in (home screen)
5. Click logout → should return to login screen
6. Close and reopen app → should land on login screen (logged out)

---

### F3: UI/UX Verification with shadcn_ui
**Goal:** Ensure all UI uses shadcn_ui components and follows design system

**Checklist:**
- [ ] Login screen uses only shadcn_ui components (TextField, Button, Alert)
- [ ] Home screen uses only shadcn_ui components (Card, Button, AlertDialog)
- [ ] Loading states show shadcn_ui Spinner
- [ ] Error states show shadcn_ui Alert
- [ ] Consistent theme across all screens
- [ ] Proper spacing and layout
- [ ] Responsive design (works on different screen sizes)
- [ ] No custom-built components that duplicate shadcn_ui functionality

**Manual Review:**
- Open each screen and verify component usage
- Check theme consistency
- Test on different screen sizes (if possible)

---

### F4: Repository Structure Verification
**Goal:** Verify foundation structure for future features is correctly set up

**Checklist:**
- [ ] Products feature folder structure complete
- [ ] Orders feature folder structure complete
- [ ] Categories feature folder structure complete
- [ ] User profile feature folder structure complete
- [ ] All models match backend schemas exactly
- [ ] All repository interfaces defined
- [ ] All datasource method signatures present
- [ ] TODO comments added for future implementation
- [ ] No broken imports or missing files
- [ ] Models have proper JSON serialization code generated

**Verification:**
```bash
# Check folder structure
find lib/features/products -type f
find lib/features/orders -type f
find lib/features/categories -type f
find lib/features/user -type f

# Verify no import errors
flutter analyze
```

---

## Definition of Done

**ALL of the following must be true:**

1. ✅ App successfully builds on Android (`flutter build apk --debug`)
2. ✅ App successfully runs on Android emulator/device
3. ✅ Login flow works with real backend at `http://localhost:8080`
4. ✅ Home screen displays user info from backend
5. ✅ Logout flow works and clears authentication
6. ✅ Navigation guards prevent unauthorized access
7. ✅ All tokens stored in flutter_secure_storage
8. ✅ HTTP interceptors working (Bearer token, cookies, error handling)
9. ✅ All UI components use shadcn_ui
10. ✅ Environment configuration works (can switch between dev/staging/prod)
11. ✅ Repository structure for future features (products, orders, categories, user) is complete
12. ✅ `flutter analyze` reports 0 issues
13. ✅ All tasks marked complete in checklist
14. ✅ All Final Verification Wave checks pass (F1-F4)

**Deliverables:**
- Fully functional Flutter app with login/home/logout
- Clean architecture folder structure
- Ready-to-extend repository structure for future features
- Environment configuration system
- Secure token storage implementation
- Complete API integration with Go backend
