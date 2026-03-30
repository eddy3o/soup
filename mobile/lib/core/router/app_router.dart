import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';
import 'package:mobile/features/auth/presentation/screens/login_screen.dart';
import 'package:mobile/features/home/presentation/screens/home_screen.dart';

class AppRouteNames {
  static const root = 'root';
  static const login = 'login';
  static const home = 'home';
}

final goRouterProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authStateProvider);

  final isLoading = authState is Initial || authState is Loading;

  return GoRouter(
    initialLocation: '/',
    debugLogDiagnostics: false,
    errorBuilder: (context, state) => Scaffold(
      body: Center(child: Text('Route not found: ${state.uri.path}')),
    ),
    redirect: (context, state) {
      final isAuthenticated = ref.watch(isAuthenticatedProvider);
      final location = state.matchedLocation;

      if (isLoading) {
        return location == '/' ? null : '/';
      }

      if (isAuthenticated && location == '/login') {
        return '/home';
      }

      if (!isAuthenticated && location == '/home') {
        return '/login';
      }

      if (location == '/') {
        return isAuthenticated ? '/home' : '/login';
      }

      return null;
    },
    routes: [
      GoRoute(
        path: '/',
        name: AppRouteNames.root,
        builder: (context, state) => const _LoadingScreen(),
      ),
      GoRoute(
        path: '/login',
        name: AppRouteNames.login,
        builder: (context, state) => const LoginScreen(),
      ),
      GoRoute(
        path: '/home',
        name: AppRouteNames.home,
        builder: (context, state) => const HomeScreen(),
      ),
    ],
  );
});

class _LoadingScreen extends StatelessWidget {
  const _LoadingScreen();

  @override
  Widget build(BuildContext context) {
    return const Scaffold(body: Center(child: CircularProgressIndicator()));
  }
}
