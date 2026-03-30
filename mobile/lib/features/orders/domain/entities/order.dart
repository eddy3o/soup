import 'order_detail.dart';

class Order {
  final String id;
  final String userId;
  final String status;
  final String pickupDate;
  final String generalNotes;
  final List<OrderDetail> orderDetails;
  final String createdAt;
  final String updatedAt;

  const Order({
    required this.id,
    required this.userId,
    required this.status,
    required this.pickupDate,
    required this.generalNotes,
    required this.orderDetails,
    required this.createdAt,
    required this.updatedAt,
  });
}
