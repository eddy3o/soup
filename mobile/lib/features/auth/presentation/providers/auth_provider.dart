import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:mobile/features/auth/data/datasources/auth_local_datasource.dart';
import 'package:mobile/features/auth/data/datasources/auth_remote_datasource.dart';
import 'package:mobile/features/auth/data/repositories/auth_repository_impl.dart';
import 'package:mobile/features/auth/domain/entities/user.dart';
import 'package:mobile/features/auth/domain/repositories/auth_repository.dart';

sealed class AuthState {
  const AuthState();

  const factory AuthState.initial() = Initial;
  const factory AuthState.loading() = Loading;
  const factory AuthState.authenticated(User user) = Authenticated;
  const factory AuthState.unauthenticated() = Unauthenticated;
  const factory AuthState.error(String message) = Error;
}

final class Initial extends AuthState {
  const Initial();
}

final class Loading extends AuthState {
  const Loading();
}

final class Authenticated extends AuthState {
  final User user;

  const Authenticated(this.user);
}

final class Unauthenticated extends AuthState {
  const Unauthenticated();
}

final class Error extends AuthState {
  final String message;

  const Error(this.message);
}

final authRepositoryProvider = Provider<AuthRepository>((ref) {
  return AuthRepositoryImpl(
    remoteDataSource: AuthRemoteDataSourceImpl(),
    localDataSource: const AuthLocalDataSource(),
  );
});

final authStateProvider = StateNotifierProvider<AuthNotifier, AuthState>((ref) {
  final repository = ref.watch(authRepositoryProvider);
  return AuthNotifier(repository);
});

final currentUserProvider = Provider<User?>((ref) {
  final authState = ref.watch(authStateProvider);
  return switch (authState) {
    Authenticated(:final user) => user,
    _ => null,
  };
});

final isAuthenticatedProvider = Provider<bool>((ref) {
  return ref.watch(authStateProvider) is Authenticated;
});

class AuthNotifier extends StateNotifier<AuthState> {
  final AuthRepository repository;

  AuthNotifier(this.repository) : super(const AuthState.initial());

  Future<void> login(String phone, String password) async {
    state = const AuthState.loading();
    final result = await repository.login(phone: phone, password: password);
    state = result.fold(
      (failure) => AuthState.error(failure.message),
      (user) => AuthState.authenticated(user),
    );
  }

  Future<void> logout() async {
    state = const AuthState.loading();
    final result = await repository.logout();
    state = result.fold(
      (failure) => AuthState.error(failure.message),
      (_) => const AuthState.unauthenticated(),
    );
  }

  Future<void> checkAuthStatus() async {
    state = const AuthState.loading();

    final authResult = await repository.isAuthenticated();
    final isAuthenticated = authResult.fold((failure) {
      state = AuthState.error(failure.message);
      return false;
    }, (value) => value);

    if (!isAuthenticated) {
      state = const AuthState.unauthenticated();
      return;
    }

    final userResult = await repository.getCurrentUser();
    state = userResult.fold(
      (failure) => AuthState.error(failure.message),
      (user) => user == null
          ? const AuthState.unauthenticated()
          : AuthState.authenticated(user),
    );
  }
}
