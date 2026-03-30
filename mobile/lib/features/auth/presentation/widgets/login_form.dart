import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';

class LoginForm extends ConsumerStatefulWidget {
  const LoginForm({super.key});

  @override
  ConsumerState<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends ConsumerState<LoginForm> {
  final _phoneController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
    _phoneController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  void _submit() {
    final phone = _phoneController.text.trim();
    final password = _passwordController.text;

    if (phone.isEmpty || password.isEmpty) return;

    ref.read(authStateProvider.notifier).login(phone, password);
  }

  @override
  Widget build(BuildContext context) {
    final authState = ref.watch(authStateProvider);
    final isLoading = authState is Loading;
    final error = authState is Error ? authState.message : null;

    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      mainAxisSize: MainAxisSize.min,
      children: [
        if (error != null) ...[
          ShadAlert.destructive(
            title: const Text('Error'),
            description: Text(error),
          ),
          const SizedBox(height: 16),
        ],
        ShadInput(
          controller: _phoneController,
          placeholder: const Text('Phone Number'),
          keyboardType: TextInputType.phone,
          enabled: !isLoading,
        ),
        const SizedBox(height: 16),
        ShadInput(
          controller: _passwordController,
          placeholder: const Text('Password'),
          obscureText: true,
          enabled: !isLoading,
        ),
        const SizedBox(height: 24),
        ShadButton(
          onPressed: isLoading ? null : _submit,
          child: Row(
            mainAxisSize: MainAxisSize.min,
            children: [
              if (isLoading) ...[
                SizedBox(
                  width: 16,
                  height: 16,
                  child: CircularProgressIndicator(strokeWidth: 2),
                ),
                const SizedBox(width: 8),
              ],
              const Text('Sign In'),
            ],
          ),
        ),
      ],
    );
  }
}
