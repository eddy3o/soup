import 'package:dio/dio.dart';
import '../config/environment.dart';
import 'interceptors/auth_interceptor.dart';
import 'interceptors/error_interceptor.dart';

class ApiClient {
  static Dio? _dio;

  static Dio get dio {
    if (_dio != null) return _dio!;

    final baseUrl = getCurrentEnvironment().getBaseUrl();

    _dio = Dio(BaseOptions(
      baseUrl: baseUrl,
      connectTimeout: const Duration(seconds: 30),
      receiveTimeout: const Duration(seconds: 30),
    ));

    _dio!.interceptors.addAll([
      AuthInterceptor(),
      ErrorInterceptor(),
      LogInterceptor(
        requestBody: true,
        responseBody: true,
        error: true,
      ),
    ]);

    return _dio!;
  }
}
