import 'package:mobile/core/errors/failures.dart';

import '../repositories/auth_repository.dart';

class CheckAuthStatusUseCase {
  final AuthRepository repository;

  const CheckAuthStatusUseCase(this.repository);

  Future<Either<Failure, bool>> call() {
    return repository.isAuthenticated();
  }
}
