import 'package:mobile/core/errors/failures.dart';

import '../../data/models/update_user_request.dart';
import '../repositories/user_repository.dart';
import '../../../auth/domain/entities/user.dart';

class UpdateUserProfileUseCase {
  final UserRepository repository;

  const UpdateUserProfileUseCase(this.repository);

  Future<Either<Failure, User>> call(UpdateUserRequest request) {
    // TODO: delegate to repository.
    return repository.updateUserProfile(request);
  }
}
