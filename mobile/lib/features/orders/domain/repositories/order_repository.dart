import 'package:mobile/core/errors/failures.dart';

import '../entities/order.dart';

abstract class OrderRepository {
  Future<Either<Failure, Order>> createOrder({required Order order});

  Future<Either<Failure, List<Order>>> getOrders();

  Future<Either<Failure, Order>> getOrderById({required String id});
}
