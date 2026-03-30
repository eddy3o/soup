import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';
import 'package:mobile/features/home/presentation/widgets/logout_button.dart';
import 'package:mobile/features/home/presentation/widgets/user_info_card.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

class HomeScreen extends ConsumerWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authStateProvider);

    return Scaffold(
      backgroundColor: ShadTheme.of(context).colorScheme.background,
      appBar: AppBar(
        backgroundColor: ShadTheme.of(context).colorScheme.background,
        elevation: 0,
        title: Text('Dashboard', style: ShadTheme.of(context).textTheme.h4),
      ),
      body: switch (authState) {
        Initial() || Loading() => const Center(child: ShadProgress()),
        Error(:final message) => Center(
          child: Padding(
            padding: const EdgeInsets.all(24.0),
            child: ShadAlert.destructive(
              title: const Text('Error'),
              description: Text(message),
            ),
          ),
        ),
        Unauthenticated() => const Center(
          child: Padding(
            padding: EdgeInsets.all(24.0),
            child: ShadAlert.destructive(
              title: Text('No autenticado'),
              description: Text('Por favor, inicie sesión para continuar.'),
            ),
          ),
        ),
        Authenticated(:final user) => SafeArea(
          child: SingleChildScrollView(
            child: Center(
              child: Container(
                width: double.infinity,
                constraints: const BoxConstraints(maxWidth: 600),
                padding: const EdgeInsets.all(24.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: [
                    Text(
                      'Welcome, ${user.name}',
                      style: ShadTheme.of(context).textTheme.h2,
                      textAlign: TextAlign.center,
                    ),
                    const SizedBox(height: 32),
                    UserInfoCard(user: user),
                    const SizedBox(height: 32),
                    const LogoutButton(),
                  ],
                ),
              ),
            ),
          ),
        ),
      },
    );
  }
}
