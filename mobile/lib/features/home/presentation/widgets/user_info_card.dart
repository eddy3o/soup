import 'package:flutter/material.dart';
import 'package:mobile/features/auth/domain/entities/user.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

class UserInfoCard extends StatelessWidget {
  const UserInfoCard({super.key, required this.user});

  final User user;

  @override
  Widget build(BuildContext context) {
    return ShadCard(
      title: Text(
        'Profile Information',
        style: ShadTheme.of(context).textTheme.h4,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 16),
          if (user.photoUrl.isNotEmpty) ...[
            Center(
              child: CircleAvatar(
                radius: 40,
                backgroundImage: NetworkImage(user.photoUrl),
                backgroundColor: ShadTheme.of(context).colorScheme.muted,
              ),
            ),
            const SizedBox(height: 16),
          ] else ...[
            Center(
              child: CircleAvatar(
                radius: 40,
                backgroundColor: ShadTheme.of(context).colorScheme.muted,
                child: Icon(
                  LucideIcons.user,
                  size: 40,
                  color: ShadTheme.of(context).colorScheme.mutedForeground,
                ),
              ),
            ),
            const SizedBox(height: 16),
          ],
          _buildInfoRow(
            context,
            icon: LucideIcons.user,
            label: 'Name',
            value: user.name,
          ),
          const SizedBox(height: 12),
          _buildInfoRow(
            context,
            icon: LucideIcons.phone,
            label: 'Phone',
            value: user.phone,
          ),
          const SizedBox(height: 12),
          _buildInfoRow(
            context,
            icon: LucideIcons.mail,
            label: 'Email',
            value: user.email.isEmpty ? 'No email provided' : user.email,
            isMuted: user.email.isEmpty,
          ),
        ],
      ),
    );
  }

  Widget _buildInfoRow(
    BuildContext context, {
    required IconData icon,
    required String label,
    required String value,
    bool isMuted = false,
  }) {
    final theme = ShadTheme.of(context);
    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Icon(icon, size: 16, color: theme.colorScheme.mutedForeground),
        const SizedBox(width: 12),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                label,
                style: theme.textTheme.small.copyWith(
                  color: theme.colorScheme.mutedForeground,
                ),
              ),
              const SizedBox(height: 2),
              Text(
                value,
                style: theme.textTheme.p.copyWith(
                  color: isMuted
                      ? theme.colorScheme.mutedForeground
                      : theme.colorScheme.foreground,
                  fontStyle: isMuted ? FontStyle.italic : FontStyle.normal,
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }
}
