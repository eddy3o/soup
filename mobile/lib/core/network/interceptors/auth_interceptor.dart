import 'package:dio/dio.dart';
import '../../../features/auth/data/datasources/auth_local_datasource.dart';
import '../../constants/api_constants.dart';

class AuthInterceptor extends Interceptor {
  final AuthLocalDataSource _localDataSource = const AuthLocalDataSource();

  @override
  void onRequest(RequestOptions options, RequestInterceptorHandler handler) async {
    // Skip auth for login/register endpoints
    if (options.path.contains(ApiConstants.AUTH_LOGIN) ||
        options.path.contains(ApiConstants.AUTH_REGISTER)) {
      return handler.next(options);
    }

    // Add Bearer token if exists
    final token = await _localDataSource.getAccessToken();
    if (token != null) {
      options.headers['Authorization'] = 'Bearer $token';
    }

    handler.next(options);
  }
}
