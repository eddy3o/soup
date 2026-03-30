import 'package:dio/dio.dart';
import '../../../features/auth/data/datasources/auth_local_datasource.dart';
import '../../errors/exceptions.dart';

class ErrorInterceptor extends Interceptor {
  final AuthLocalDataSource _localDataSource = const AuthLocalDataSource();

  @override
  void onError(DioException err, ErrorInterceptorHandler handler) async {
    if (err.response?.statusCode == 401) {
      // Unauthorized - clear tokens
      await _localDataSource.clearTokens();
      handler.reject(DioException(
        requestOptions: err.requestOptions,
        error: const UnauthorizedException('Unauthorized. Please login again.'),
      ));
    } else if (err.response?.statusCode == 400) {
      // Bad request
      final message = err.response?.data?['error'] ??
          err.response?.data?['message'] ??
          'Bad request';
      handler.reject(DioException(
        requestOptions: err.requestOptions,
        error: ServerException(message.toString()),
      ));
    } else if (err.response?.statusCode != null &&
        err.response!.statusCode! >= 500) {
      // Server error
      handler.reject(DioException(
        requestOptions: err.requestOptions,
        error: const ServerException('Server error. Please try again later.'),
      ));
    } else if (err.type == DioExceptionType.connectionTimeout ||
        err.type == DioExceptionType.sendTimeout ||
        err.type == DioExceptionType.receiveTimeout ||
        err.type == DioExceptionType.connectionError) {
      // Network error
      handler.reject(DioException(
        requestOptions: err.requestOptions,
        error: const NetworkException(
            'Connection failed. Please check your internet.'),
      ));
    } else {
      handler.next(err);
    }
  }
}
