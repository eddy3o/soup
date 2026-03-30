import 'package:mobile/core/errors/failures.dart';

import '../entities/category.dart';

abstract class CategoryRepository {
  Future<Either<Failure, List<Category>>> getCategories();
}
