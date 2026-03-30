enum Environment { dev, staging, prod }

extension EnvironmentConfig on Environment {
  String getBaseUrl() {
    switch (this) {
      case Environment.dev:
        return 'http://localhost:8080';
      case Environment.staging:
        return 'https://staging-api.soup.com';
      case Environment.prod:
        return 'https://api.soup.com';
    }
  }
}

Environment getCurrentEnvironment() {
  const envName = String.fromEnvironment('ENV', defaultValue: 'dev');

  switch (envName.toLowerCase()) {
    case 'staging':
      return Environment.staging;
    case 'prod':
      return Environment.prod;
    case 'dev':
    default:
      return Environment.dev;
  }
}
