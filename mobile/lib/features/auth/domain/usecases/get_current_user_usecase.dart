import 'package:mobile/core/errors/failures.dart';

import '../entities/user.dart';
import '../repositories/auth_repository.dart';

class GetCurrentUserUseCase {
  final AuthRepository repository;

  const GetCurrentUserUseCase(this.repository);

  Future<Either<Failure, User?>> call() {
    return repository.getCurrentUser();
  }
}
