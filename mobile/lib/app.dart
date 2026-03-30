import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'core/router/app_router.dart';
import 'features/auth/presentation/providers/auth_provider.dart';

class MyApp extends ConsumerStatefulWidget {
  const MyApp({super.key});

  @override
  ConsumerState<MyApp> createState() => _MyAppState();
}

class _MyAppState extends ConsumerState<MyApp> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      ref.read(authStateProvider.notifier).checkAuthStatus();
    });
  }

  @override
  Widget build(BuildContext context) {
    final goRouter = ref.watch(goRouterProvider);

    return ShadApp.custom(
      themeMode: ThemeMode.light,
      theme: ShadThemeData(
        brightness: Brightness.light,
        colorScheme: const ShadZincColorScheme.light(),
      ),
      darkTheme: ShadThemeData(
        brightness: Brightness.dark,
        colorScheme: const ShadZincColorScheme.dark(),
      ),
      appBuilder: (context) {
        return MaterialApp.router(
          title: 'Soup App',
          routerConfig: goRouter,
          theme: Theme.of(context),
          darkTheme: Theme.of(context),
          themeMode: ThemeMode.light,
          debugShowCheckedModeBanner: false,
          builder: (context, child) {
            return ShadAppBuilder(child: child ?? const SizedBox.shrink());
          },
        );
      },
    );
  }
}
