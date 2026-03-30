import 'package:mobile/core/errors/failures.dart';

import '../../../auth/domain/entities/user.dart';
import '../repositories/user_repository.dart';

class GetUserProfileUseCase {
  final UserRepository repository;

  const GetUserProfileUseCase(this.repository);

  Future<Either<Failure, User>> call() {
    // TODO: delegate to repository.
    return repository.getUserProfile();
  }
}
