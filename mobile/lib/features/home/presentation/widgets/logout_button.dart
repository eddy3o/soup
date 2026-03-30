import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

class LogoutButton extends ConsumerWidget {
  const LogoutButton({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authStateProvider);
    final isLoading = authState is Loading;

    return ShadButton.destructive(
      onPressed: isLoading ? null : () => _handleLogout(context, ref),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          if (isLoading)
            SizedBox(
              width: 16,
              height: 16,
              child: CircularProgressIndicator(
                strokeWidth: 2,
                color: ShadTheme.of(context).colorScheme.destructiveForeground,
              ),
            )
          else
            const Icon(LucideIcons.logOut, size: 16),
          const SizedBox(width: 8),
          const Text('Logout'),
        ],
      ),
    );
  }

  Future<void> _handleLogout(BuildContext context, WidgetRef ref) async {
    final confirm = await showShadDialog<bool>(
      context: context,
      builder: (context) => ShadDialog.alert(
        title: const Text('Confirm Logout'),
        description: const Padding(
          padding: EdgeInsets.only(bottom: 8),
          child: Text('Are you sure you want to log out of your account?'),
        ),
        actions: [
          ShadButton.outline(
            child: const Text('Cancel'),
            onPressed: () => Navigator.pop(context, false),
          ),
          ShadButton.destructive(
            child: const Text('Logout'),
            onPressed: () => Navigator.pop(context, true),
          ),
        ],
      ),
    );

    if (confirm == true) {
      ref.read(authStateProvider.notifier).logout();
    }
  }
}
