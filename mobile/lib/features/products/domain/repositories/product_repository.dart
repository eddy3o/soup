import 'package:mobile/core/errors/failures.dart';

import '../entities/product.dart';

abstract class ProductRepository {
  Future<Either<Failure, List<Product>>> getProducts({
    required int page,
    required int limit,
  });

  Future<Either<Failure, Product>> getProductById({required String id});
}
