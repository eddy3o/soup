import 'exceptions.dart';

sealed class Either<L, R> {
  const Either();

  T fold<T>(T Function(L left) onLeft, T Function(R right) onRight);
}

final class Left<L, R> extends Either<L, R> {
  final L value;

  const Left(this.value);

  @override
  T fold<T>(T Function(L left) onLeft, T Function(R right) onRight) =>
      onLeft(value);
}

final class Right<L, R> extends Either<L, R> {
  final R value;

  const Right(this.value);

  @override
  T fold<T>(T Function(L left) onLeft, T Function(R right) onRight) =>
      onRight(value);
}

sealed class Failure {
  final String message;

  const Failure(this.message);
}

final class ServerFailure extends Failure {
  const ServerFailure(super.message);
}

final class CacheFailure extends Failure {
  const CacheFailure(super.message);
}

final class NetworkFailure extends Failure {
  const NetworkFailure(super.message);
}

final class UnauthorizedFailure extends Failure {
  const UnauthorizedFailure(super.message);
}

final class UnimplementedFailure extends Failure {
  const UnimplementedFailure(super.message);
}

Failure exceptionToFailure(Exception e) {
  return switch (e) {
    ServerException(:final message) => ServerFailure(message),
    CacheException(:final message) => CacheFailure(message),
    NetworkException(:final message) => NetworkFailure(message),
    UnauthorizedException(:final message) => UnauthorizedFailure(message),
    _ => ServerFailure(e.toString()),
  };
}
