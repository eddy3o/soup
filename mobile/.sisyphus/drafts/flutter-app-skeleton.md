# Draft: Flutter App Skeleton with Login/Home/Logout

## User's Original Request
- Flutter app in /home/eddy/proyectos/soup/mobile
- Backend is outside this directory (in soup, but not in mobile)
- Need repo files for fetching data from backend
- App screens: login, home, and logout
- Use shadcn_ui package from pub.dev

## Current State
- Fresh Flutter project (SDK 3.11.3)
- Only has default main.dart
- No dependencies beyond flutter and cupertino_icons
- No existing architecture

## Research Findings

### Backend Analysis (COMPLETE)
- **Framework**: Go with Gin Web Framework
- **Auth Method**: JWT (HS256) - dual delivery (HTTP-only cookies + Bearer token)
- **Token Storage**: Redis-backed stateful validation
- **Token Expiry**: Access token 15min, Refresh token 7 days
- **Base URL**: `http://localhost:8080`
- **Login Credentials**: phone (min 6 chars) + password (min 6 chars)

### Available Endpoints
```
POST /auth/login       → {access_token, refresh_token, user}
POST /auth/register
POST /auth/logout
GET  /users/me
PATCH /users/me
GET  /products         (pagination support)
GET  /products/:id
GET  /categories
POST /orders
```

### Response Schemas
- User: {id, phone, name, address, email, photo_url, is_admin, push_token, created_at}
- Product: {id, name, description, price, photo_url, available, category, timestamps}
- Products endpoint returns: {data: [], pagination: {page, limit, total_items, total_pages}}

## Requirements to Clarify

## Technical Decisions
(to be recorded)

## Scope Boundaries
- INCLUDE: TBD
- EXCLUDE: TBD
