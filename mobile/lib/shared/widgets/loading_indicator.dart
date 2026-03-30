import 'package:flutter/widgets.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

/// Reusable loading indicator built with shadcn_ui.
class LoadingIndicator extends StatelessWidget {
  const LoadingIndicator({super.key, this.size = 16});

  /// Size of the indicator in logical pixels.
  final double size;

  @override
  Widget build(BuildContext context) {
    return SizedBox.square(
      dimension: size,
      child: const Center(child: ShadProgress()),
    );
  }
}
