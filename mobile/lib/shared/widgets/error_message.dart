import 'package:flutter/widgets.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

/// Severity level for a reusable error message.
enum ErrorMessageSeverity { info, destructive }

/// A reusable shadcn alert for displaying errors or warnings.
class ErrorMessage extends StatelessWidget {
  const ErrorMessage({
    super.key,
    required this.message,
    this.severity = ErrorMessageSeverity.destructive,
  });

  /// Message shown in the alert.
  final String message;

  /// Controls the alert style.
  final ErrorMessageSeverity severity;

  @override
  Widget build(BuildContext context) {
    final isDestructive = severity == ErrorMessageSeverity.destructive;
    final title = isDestructive ? 'Error' : 'Notice';

    return isDestructive
        ? ShadAlert.destructive(title: Text(title), description: Text(message))
        : ShadAlert(title: Text(title), description: Text(message));
  }
}
