# Push Run App Backend

Backend service for Push Run App - a Strava-like application for the Russian market.

## Features

- User authentication with GitHub OAuth
- Activity tracking and statistics
- Social features (friends, likes, comments)
- Achievements system
- Challenges
- Route planning with Yandex Maps integration
- Weather information

## Tech Stack

- NestJS
- TypeScript
- PostgreSQL
- TypeORM
- JWT Authentication
- GitHub OAuth
- Yandex Maps API

## Prerequisites

- Node.js (v16 or later)
- PostgreSQL
- GitHub OAuth App credentials
- Yandex Maps API key

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/push-run-app.git
cd push-run-app/backend-nest
```

2. Install dependencies:
```bash
npm install
```

3. Create a `.env` file in the root directory with the following variables:
```env
# Server
PORT=3000

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_DATABASE=push_run_app

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRES_IN=7d

# GitHub OAuth
GITHUB_CLIENT_ID=your-github-client-id
GITHUB_CLIENT_SECRET=your-github-client-secret
GITHUB_CALLBACK_URL=http://localhost:3000/auth/github/callback

# Yandex Maps
YANDEX_MAPS_API_KEY=your-yandex-maps-api-key
```

4. Start the development server:
```bash
npm run start:dev
```

The server will start at http://localhost:3000

## API Documentation

Once the server is running, you can access the Swagger API documentation at:
http://localhost:3000/api

## Project Structure

```
backend-nest/
├── src/
│   ├── models/           # Data models and DTOs
│   ├── modules/          # Feature modules
│   │   ├── users/        # User management
│   │   ├── auth/         # Authentication
│   │   ├── activities/   # Activity tracking
│   │   ├── challenges/   # Challenges system
│   │   ├── achievements/ # Achievements system
│   │   └── maps/         # Maps and routes
│   ├── config/           # Configuration
│   ├── guards/           # Authentication guards
│   ├── decorators/       # Custom decorators
│   ├── interfaces/       # TypeScript interfaces
│   ├── app.module.ts     # Root module
│   └── main.ts           # Application entry point
├── test/                 # Test files
├── .env                  # Environment variables
├── .gitignore           # Git ignore file
├── package.json         # Project dependencies
└── tsconfig.json        # TypeScript configuration
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 