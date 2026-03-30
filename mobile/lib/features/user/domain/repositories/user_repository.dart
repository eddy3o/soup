import 'package:mobile/core/errors/failures.dart';

import '../../data/models/update_user_request.dart';
import '../../../auth/domain/entities/user.dart';

abstract class UserRepository {
  Future<Either<Failure, User>> getUserProfile();

  Future<Either<Failure, User>> updateUserProfile(UpdateUserRequest request);
}
