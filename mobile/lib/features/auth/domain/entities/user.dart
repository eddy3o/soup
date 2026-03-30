class User {
  final String id;
  final String phone;
  final String name;
  final String address;
  final String email;
  final String photoUrl;
  final String isAdmin;
  final String pushToken;
  final String createdAt;

  const User({
    required this.id,
    required this.phone,
    required this.name,
    required this.address,
    required this.email,
    required this.photoUrl,
    required this.isAdmin,
    required this.pushToken,
    required this.createdAt,
  });
}
