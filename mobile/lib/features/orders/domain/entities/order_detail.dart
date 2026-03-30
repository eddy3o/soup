class OrderDetail {
  final String id;
  final String productId;
  final int quantity;
  final double price;
  final double subtotal;

  const OrderDetail({
    required this.id,
    required this.productId,
    required this.quantity,
    required this.price,
    required this.subtotal,
  });
}
