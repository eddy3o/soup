import 'package:mobile/core/errors/failures.dart';

import '../entities/user.dart';

abstract class AuthRepository {
  Future<Either<Failure, User>> login({
    required String phone,
    required String password,
  });

  Future<Either<Failure, void>> logout();

  Future<Either<Failure, User?>> getCurrentUser();

  Future<Either<Failure, bool>> isAuthenticated();
}
